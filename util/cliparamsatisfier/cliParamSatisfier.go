package cliparamsatisfier

//go:generate counterfeiter -o ./fake.go --fake-name Fake ./ CLIParamSatisfier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"
	"sort"

	"github.com/ghodss/yaml"
	"github.com/opctl/opctl/util/cliexiter"
	"github.com/opctl/opctl/util/clioutput"
	"github.com/opspec-io/sdk-golang/data/coerce"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/op/interpreter/opcall/params"
)

// CLIParamSatisfier attempts to satisfy the provided inputs via the provided inputSourcer
//
// if all fails an error is logged and we exit with a nonzero code.
type CLIParamSatisfier interface {
	InputSrcFactory

	Satisfy(
		inputSourcer InputSourcer,
		inputs map[string]*model.Param,
	) map[string]*model.Value
}

func New(
	cliExiter cliexiter.CliExiter,
	cliOutput clioutput.CliOutput,
) CLIParamSatisfier {

	return &_CLIParamSatisfier{
		cliExiter:       cliExiter,
		cliOutput:       cliOutput,
		coerce:          coerce.New(),
		paramsValidator: params.NewValidator(),
		InputSrcFactory: newInputSrcFactory(),
		InputInterpolater: newInputInterpolater(),
	}
}

type _CLIParamSatisfier struct {
	cliExiter       cliexiter.CliExiter
	cliOutput       clioutput.CliOutput
	coerce          coerce.Coerce
	paramsValidator params.Validator
	InputSrcFactory
	InputInterpolater
}

func (cps _CLIParamSatisfier) Satisfy(
	inputSourcer InputSourcer,
	inputs map[string]*model.Param,
) map[string]*model.Value {

	argMap := map[string]*model.Value{}
	for _, paramName := range cps.getSortedParamNames(inputs) {
		param := inputs[paramName]

	paramLoop:
		for {
			var arg *model.Value

			rawArg, ok := inputSourcer.Source(paramName)
			if !ok {
				msg := fmt.Sprintf(`
-
  Prompt for "%v" failed; running in non-interactive terminal
-`, paramName)
				cps.cliExiter.Exit(cliexiter.ExitReq{Message: msg, Code: 1})
			}

			if nil != rawArg {
				interpolatedArg, err := cps.Interpolate(*rawArg)
				if nil != err {
					// param not satisfied; notify & re-attempt!
					cps.notifyOfArgErrors([]error{err}, paramName)
					continue
				}
				rawArg = &interpolatedArg
			}

			switch {
			case nil == rawArg:
				// handle nil (returned by inputSourcer.Source for static defaults)
				break paramLoop
			case nil != param.Array:
				argValue := []interface{}{}
				argJsonBytes, err := yaml.YAMLToJSON([]byte(*rawArg))
				if nil != err {
					// param not satisfied; notify & re-attempt!
					cps.notifyOfArgErrors([]error{err}, paramName)
					continue
				}
				err = json.Unmarshal(argJsonBytes, &argValue)
				if nil != err {
					// param not satisfied; notify & re-attempt!
					cps.notifyOfArgErrors([]error{err}, paramName)
					continue
				}
				arg = &model.Value{Array: argValue}
			case nil != param.Boolean:
				var err error
				if arg, err = cps.coerce.ToBoolean(&model.Value{String: rawArg}); nil != err {
					// param not satisfied; notify & re-attempt!
					cps.notifyOfArgErrors([]error{err}, paramName)
					continue
				}
			case nil != param.Dir:
				absPath, err := filepath.Abs(*rawArg)
				if nil != err {
					// param not satisfied; notify & re-attempt!
					cps.notifyOfArgErrors([]error{err}, paramName)
					continue
				}
				arg = &model.Value{Dir: &absPath}
			case nil != param.File:
				absPath, err := filepath.Abs(*rawArg)
				if nil != err {
					// param not satisfied; notify & re-attempt!
					cps.notifyOfArgErrors([]error{err}, paramName)
					continue
				}
				arg = &model.Value{File: &absPath}
			case nil != param.Number:
				var err error
				if arg, err = cps.coerce.ToNumber(&model.Value{String: rawArg}); nil != err {
					// param not satisfied; notify & re-attempt!
					cps.notifyOfArgErrors([]error{err}, paramName)
					continue
				}
			case nil != param.Object:
				argValue := map[string]interface{}{}
				argJsonBytes, err := yaml.YAMLToJSON([]byte(*rawArg))
				if nil != err {
					// param not satisfied; notify & re-attempt!
					cps.notifyOfArgErrors([]error{err}, paramName)
					continue
				}
				err = json.Unmarshal(argJsonBytes, &argValue)
				if nil != err {
					// param not satisfied; notify & re-attempt!
					cps.notifyOfArgErrors([]error{err}, paramName)
					continue
				}
				arg = &model.Value{Object: argValue}
			case nil != param.Socket:
				arg = &model.Value{Socket: rawArg}
			case nil != param.String:
				arg = &model.Value{String: rawArg}
			}

			validateErr := cps.paramsValidator.Validate(
				map[string]*model.Value{paramName: arg},
				map[string]*model.Param{paramName: param},
			)
			if nil != validateErr {
				cps.notifyOfArgErrors([]error{validateErr}, paramName)

				// param not satisfied; re-attempt it!
				continue
			}

			if nil != arg {
				// only store non-nil args
				argMap[paramName] = arg
			}

			// param satisfied; move to next!
			break paramLoop
		}
	}

	return argMap
}

func (this _CLIParamSatisfier) getSortedParamNames(
	params map[string]*model.Param,
) []string {
	paramNames := []string{}
	for paramname := range params {
		paramNames = append(paramNames, paramname)
	}
	sort.Strings(paramNames)
	return paramNames
}

func (this _CLIParamSatisfier) notifyOfArgErrors(
	errors []error,
	paramName string,
) {
	messageBuffer := bytes.NewBufferString(
		fmt.Sprintf(`
-
  %v invalid; provide valid value to proceed.
  Error(s):`, paramName),
	)
	for _, validationError := range errors {
		messageBuffer.WriteString(fmt.Sprintf(`
    - %v`, validationError.Error()))
	}
	this.cliOutput.Error(`
%v
-`, messageBuffer.String())
}
