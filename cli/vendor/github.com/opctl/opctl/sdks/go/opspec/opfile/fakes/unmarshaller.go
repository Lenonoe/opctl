// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/opspec/opfile"
)

type FakeUnmarshaller struct {
	UnmarshalStub        func([]byte) (*model.OpFile, error)
	unmarshalMutex       sync.RWMutex
	unmarshalArgsForCall []struct {
		arg1 []byte
	}
	unmarshalReturns struct {
		result1 *model.OpFile
		result2 error
	}
	unmarshalReturnsOnCall map[int]struct {
		result1 *model.OpFile
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUnmarshaller) Unmarshal(arg1 []byte) (*model.OpFile, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.unmarshalMutex.Lock()
	ret, specificReturn := fake.unmarshalReturnsOnCall[len(fake.unmarshalArgsForCall)]
	fake.unmarshalArgsForCall = append(fake.unmarshalArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("Unmarshal", []interface{}{arg1Copy})
	fake.unmarshalMutex.Unlock()
	if fake.UnmarshalStub != nil {
		return fake.UnmarshalStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.unmarshalReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUnmarshaller) UnmarshalCallCount() int {
	fake.unmarshalMutex.RLock()
	defer fake.unmarshalMutex.RUnlock()
	return len(fake.unmarshalArgsForCall)
}

func (fake *FakeUnmarshaller) UnmarshalCalls(stub func([]byte) (*model.OpFile, error)) {
	fake.unmarshalMutex.Lock()
	defer fake.unmarshalMutex.Unlock()
	fake.UnmarshalStub = stub
}

func (fake *FakeUnmarshaller) UnmarshalArgsForCall(i int) []byte {
	fake.unmarshalMutex.RLock()
	defer fake.unmarshalMutex.RUnlock()
	argsForCall := fake.unmarshalArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUnmarshaller) UnmarshalReturns(result1 *model.OpFile, result2 error) {
	fake.unmarshalMutex.Lock()
	defer fake.unmarshalMutex.Unlock()
	fake.UnmarshalStub = nil
	fake.unmarshalReturns = struct {
		result1 *model.OpFile
		result2 error
	}{result1, result2}
}

func (fake *FakeUnmarshaller) UnmarshalReturnsOnCall(i int, result1 *model.OpFile, result2 error) {
	fake.unmarshalMutex.Lock()
	defer fake.unmarshalMutex.Unlock()
	fake.UnmarshalStub = nil
	if fake.unmarshalReturnsOnCall == nil {
		fake.unmarshalReturnsOnCall = make(map[int]struct {
			result1 *model.OpFile
			result2 error
		})
	}
	fake.unmarshalReturnsOnCall[i] = struct {
		result1 *model.OpFile
		result2 error
	}{result1, result2}
}

func (fake *FakeUnmarshaller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.unmarshalMutex.RLock()
	defer fake.unmarshalMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUnmarshaller) recordInvocation(key string, args []interface{}) {
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

var _ opfile.Unmarshaller = new(FakeUnmarshaller)
