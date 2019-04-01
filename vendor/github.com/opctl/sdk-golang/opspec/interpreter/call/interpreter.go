package call

import (
	"fmt"
	"github.com/opctl/sdk-golang/model"
	"github.com/opctl/sdk-golang/opspec/interpreter/call/container"
	"github.com/opctl/sdk-golang/opspec/interpreter/call/if/predicates"
	"github.com/opctl/sdk-golang/opspec/interpreter/call/op"
)

//go:generate counterfeiter -o ./fakeInterpreter.go --fake-name FakeInterpreter ./ Interpreter

type Interpreter interface {
	// Interpret interprets an SCG into a DCG
	Interpret(
		scope map[string]*model.Value,
		scg *model.SCG,
		id string,
		opHandle model.DataHandle,
		rootOpID string,
	) (*model.DCG, error)
}

// NewInterpreter returns an initialized Interpreter instance
func NewInterpreter(
	containerCallInterpreter container.Interpreter,
	dataDirPath string,
) Interpreter {
	return _interpreter{
		containerCallInterpreter: containerCallInterpreter,
		predicatesInterpreter:    predicates.NewInterpreter(),
		opCallInterpreter:        op.NewInterpreter(dataDirPath),
	}
}

type _interpreter struct {
	containerCallInterpreter container.Interpreter
	opCallInterpreter        op.Interpreter
	predicatesInterpreter    predicates.Interpreter
}

func (itp _interpreter) Interpret(
	scope map[string]*model.Value,
	scg *model.SCG,
	id string,
	opHandle model.DataHandle,
	rootOpID string,
) (*model.DCG, error) {
	var ifPredicate *bool
	if len(scg.If) > 0 {
		ifPredicateBool, err := itp.predicatesInterpreter.Interpret(
			opHandle,
			scg.If,
			scope,
		)
		if nil != err {
			return nil, err
		}

		ifPredicate = &ifPredicateBool
	}

	switch {
	case nil != scg.Container:
		scgContainerCall, err := itp.containerCallInterpreter.Interpret(
			scope,
			scg.Container,
			id,
			rootOpID,
			opHandle,
		)
		if nil != err {
			return nil, err
		}

		return &model.DCG{
			Container: scgContainerCall,
			If:        ifPredicate,
		}, nil
	case nil != scg.Op:
		dcgOpCall, err := itp.opCallInterpreter.Interpret(
			scope,
			scg.Op,
			id,
			opHandle,
			rootOpID,
		)
		if nil != err {
			return nil, err
		}

		return &model.DCG{
			Op: dcgOpCall,
			If: ifPredicate,
		}, nil
	case len(scg.Parallel) > 0:
		return &model.DCG{
			Parallel: scg.Parallel,
			If:       ifPredicate,
		}, nil
	case len(scg.Serial) > 0:
		return &model.DCG{
			Serial: scg.Serial,
			If:     ifPredicate,
		}, nil
	default:
		return nil, fmt.Errorf("Invalid call graph %+v\n", scg)
	}
}