// This file was generated by counterfeiter
package core

import (
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type fakeOpCaller struct {
	CallStub        func(inboundScope map[string]*model.Value, opId string, pkgRef string, rootOpId string, scgOpCall *model.SCGOpCall) (err error)
	callMutex       sync.RWMutex
	callArgsForCall []struct {
		inboundScope map[string]*model.Value
		opId         string
		pkgRef       string
		rootOpId     string
		scgOpCall    *model.SCGOpCall
	}
	callReturns struct {
		result1 error
	}
	callReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeOpCaller) Call(inboundScope map[string]*model.Value, opId string, pkgRef string, rootOpId string, scgOpCall *model.SCGOpCall) (err error) {
	fake.callMutex.Lock()
	ret, specificReturn := fake.callReturnsOnCall[len(fake.callArgsForCall)]
	fake.callArgsForCall = append(fake.callArgsForCall, struct {
		inboundScope map[string]*model.Value
		opId         string
		pkgRef       string
		rootOpId     string
		scgOpCall    *model.SCGOpCall
	}{inboundScope, opId, pkgRef, rootOpId, scgOpCall})
	fake.recordInvocation("Call", []interface{}{inboundScope, opId, pkgRef, rootOpId, scgOpCall})
	fake.callMutex.Unlock()
	if fake.CallStub != nil {
		return fake.CallStub(inboundScope, opId, pkgRef, rootOpId, scgOpCall)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.callReturns.result1
}

func (fake *fakeOpCaller) CallCallCount() int {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return len(fake.callArgsForCall)
}

func (fake *fakeOpCaller) CallArgsForCall(i int) (map[string]*model.Value, string, string, string, *model.SCGOpCall) {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return fake.callArgsForCall[i].inboundScope, fake.callArgsForCall[i].opId, fake.callArgsForCall[i].pkgRef, fake.callArgsForCall[i].rootOpId, fake.callArgsForCall[i].scgOpCall
}

func (fake *fakeOpCaller) CallReturns(result1 error) {
	fake.CallStub = nil
	fake.callReturns = struct {
		result1 error
	}{result1}
}

func (fake *fakeOpCaller) CallReturnsOnCall(i int, result1 error) {
	fake.CallStub = nil
	if fake.callReturnsOnCall == nil {
		fake.callReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.callReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *fakeOpCaller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return fake.invocations
}

func (fake *fakeOpCaller) recordInvocation(key string, args []interface{}) {
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
