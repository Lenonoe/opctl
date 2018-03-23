// Code generated by counterfeiter. DO NOT EDIT.
package dotyml

import (
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type FakeUnmarshaller struct {
	UnmarshalStub        func(manifestBytes []byte) (*model.OpDotYml, error)
	unmarshalMutex       sync.RWMutex
	unmarshalArgsForCall []struct {
		manifestBytes []byte
	}
	unmarshalReturns struct {
		result1 *model.OpDotYml
		result2 error
	}
	unmarshalReturnsOnCall map[int]struct {
		result1 *model.OpDotYml
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUnmarshaller) Unmarshal(manifestBytes []byte) (*model.OpDotYml, error) {
	var manifestBytesCopy []byte
	if manifestBytes != nil {
		manifestBytesCopy = make([]byte, len(manifestBytes))
		copy(manifestBytesCopy, manifestBytes)
	}
	fake.unmarshalMutex.Lock()
	ret, specificReturn := fake.unmarshalReturnsOnCall[len(fake.unmarshalArgsForCall)]
	fake.unmarshalArgsForCall = append(fake.unmarshalArgsForCall, struct {
		manifestBytes []byte
	}{manifestBytesCopy})
	fake.recordInvocation("Unmarshal", []interface{}{manifestBytesCopy})
	fake.unmarshalMutex.Unlock()
	if fake.UnmarshalStub != nil {
		return fake.UnmarshalStub(manifestBytes)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.unmarshalReturns.result1, fake.unmarshalReturns.result2
}

func (fake *FakeUnmarshaller) UnmarshalCallCount() int {
	fake.unmarshalMutex.RLock()
	defer fake.unmarshalMutex.RUnlock()
	return len(fake.unmarshalArgsForCall)
}

func (fake *FakeUnmarshaller) UnmarshalArgsForCall(i int) []byte {
	fake.unmarshalMutex.RLock()
	defer fake.unmarshalMutex.RUnlock()
	return fake.unmarshalArgsForCall[i].manifestBytes
}

func (fake *FakeUnmarshaller) UnmarshalReturns(result1 *model.OpDotYml, result2 error) {
	fake.UnmarshalStub = nil
	fake.unmarshalReturns = struct {
		result1 *model.OpDotYml
		result2 error
	}{result1, result2}
}

func (fake *FakeUnmarshaller) UnmarshalReturnsOnCall(i int, result1 *model.OpDotYml, result2 error) {
	fake.UnmarshalStub = nil
	if fake.unmarshalReturnsOnCall == nil {
		fake.unmarshalReturnsOnCall = make(map[int]struct {
			result1 *model.OpDotYml
			result2 error
		})
	}
	fake.unmarshalReturnsOnCall[i] = struct {
		result1 *model.OpDotYml
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

var _ Unmarshaller = new(FakeUnmarshaller)
