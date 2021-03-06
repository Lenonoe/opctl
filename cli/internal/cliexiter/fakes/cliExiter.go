// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/opctl/opctl/cli/internal/cliexiter"
)

type FakeCliExiter struct {
	ExitStub        func(cliexiter.ExitReq)
	exitMutex       sync.RWMutex
	exitArgsForCall []struct {
		arg1 cliexiter.ExitReq
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCliExiter) Exit(arg1 cliexiter.ExitReq) {
	fake.exitMutex.Lock()
	fake.exitArgsForCall = append(fake.exitArgsForCall, struct {
		arg1 cliexiter.ExitReq
	}{arg1})
	fake.recordInvocation("Exit", []interface{}{arg1})
	fake.exitMutex.Unlock()
	if fake.ExitStub != nil {
		fake.ExitStub(arg1)
	}
}

func (fake *FakeCliExiter) ExitCallCount() int {
	fake.exitMutex.RLock()
	defer fake.exitMutex.RUnlock()
	return len(fake.exitArgsForCall)
}

func (fake *FakeCliExiter) ExitCalls(stub func(cliexiter.ExitReq)) {
	fake.exitMutex.Lock()
	defer fake.exitMutex.Unlock()
	fake.ExitStub = stub
}

func (fake *FakeCliExiter) ExitArgsForCall(i int) cliexiter.ExitReq {
	fake.exitMutex.RLock()
	defer fake.exitMutex.RUnlock()
	argsForCall := fake.exitArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCliExiter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.exitMutex.RLock()
	defer fake.exitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCliExiter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ cliexiter.CliExiter = new(FakeCliExiter)
