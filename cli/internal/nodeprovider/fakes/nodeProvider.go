// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/opctl/opctl/cli/internal/model"
	"github.com/opctl/opctl/cli/internal/nodeprovider"
)

type FakeNodeProvider struct {
	CreateNodeIfNotExistsStub        func() (model.NodeHandle, error)
	createNodeIfNotExistsMutex       sync.RWMutex
	createNodeIfNotExistsArgsForCall []struct {
	}
	createNodeIfNotExistsReturns struct {
		result1 model.NodeHandle
		result2 error
	}
	createNodeIfNotExistsReturnsOnCall map[int]struct {
		result1 model.NodeHandle
		result2 error
	}
	KillNodeIfExistsStub        func(string) error
	killNodeIfExistsMutex       sync.RWMutex
	killNodeIfExistsArgsForCall []struct {
		arg1 string
	}
	killNodeIfExistsReturns struct {
		result1 error
	}
	killNodeIfExistsReturnsOnCall map[int]struct {
		result1 error
	}
	ListNodesStub        func() ([]model.NodeHandle, error)
	listNodesMutex       sync.RWMutex
	listNodesArgsForCall []struct {
	}
	listNodesReturns struct {
		result1 []model.NodeHandle
		result2 error
	}
	listNodesReturnsOnCall map[int]struct {
		result1 []model.NodeHandle
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNodeProvider) CreateNodeIfNotExists() (model.NodeHandle, error) {
	fake.createNodeIfNotExistsMutex.Lock()
	ret, specificReturn := fake.createNodeIfNotExistsReturnsOnCall[len(fake.createNodeIfNotExistsArgsForCall)]
	fake.createNodeIfNotExistsArgsForCall = append(fake.createNodeIfNotExistsArgsForCall, struct {
	}{})
	fake.recordInvocation("CreateNodeIfNotExists", []interface{}{})
	fake.createNodeIfNotExistsMutex.Unlock()
	if fake.CreateNodeIfNotExistsStub != nil {
		return fake.CreateNodeIfNotExistsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createNodeIfNotExistsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNodeProvider) CreateNodeIfNotExistsCallCount() int {
	fake.createNodeIfNotExistsMutex.RLock()
	defer fake.createNodeIfNotExistsMutex.RUnlock()
	return len(fake.createNodeIfNotExistsArgsForCall)
}

func (fake *FakeNodeProvider) CreateNodeIfNotExistsCalls(stub func() (model.NodeHandle, error)) {
	fake.createNodeIfNotExistsMutex.Lock()
	defer fake.createNodeIfNotExistsMutex.Unlock()
	fake.CreateNodeIfNotExistsStub = stub
}

func (fake *FakeNodeProvider) CreateNodeIfNotExistsReturns(result1 model.NodeHandle, result2 error) {
	fake.createNodeIfNotExistsMutex.Lock()
	defer fake.createNodeIfNotExistsMutex.Unlock()
	fake.CreateNodeIfNotExistsStub = nil
	fake.createNodeIfNotExistsReturns = struct {
		result1 model.NodeHandle
		result2 error
	}{result1, result2}
}

func (fake *FakeNodeProvider) CreateNodeIfNotExistsReturnsOnCall(i int, result1 model.NodeHandle, result2 error) {
	fake.createNodeIfNotExistsMutex.Lock()
	defer fake.createNodeIfNotExistsMutex.Unlock()
	fake.CreateNodeIfNotExistsStub = nil
	if fake.createNodeIfNotExistsReturnsOnCall == nil {
		fake.createNodeIfNotExistsReturnsOnCall = make(map[int]struct {
			result1 model.NodeHandle
			result2 error
		})
	}
	fake.createNodeIfNotExistsReturnsOnCall[i] = struct {
		result1 model.NodeHandle
		result2 error
	}{result1, result2}
}

func (fake *FakeNodeProvider) KillNodeIfExists(arg1 string) error {
	fake.killNodeIfExistsMutex.Lock()
	ret, specificReturn := fake.killNodeIfExistsReturnsOnCall[len(fake.killNodeIfExistsArgsForCall)]
	fake.killNodeIfExistsArgsForCall = append(fake.killNodeIfExistsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("KillNodeIfExists", []interface{}{arg1})
	fake.killNodeIfExistsMutex.Unlock()
	if fake.KillNodeIfExistsStub != nil {
		return fake.KillNodeIfExistsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.killNodeIfExistsReturns
	return fakeReturns.result1
}

func (fake *FakeNodeProvider) KillNodeIfExistsCallCount() int {
	fake.killNodeIfExistsMutex.RLock()
	defer fake.killNodeIfExistsMutex.RUnlock()
	return len(fake.killNodeIfExistsArgsForCall)
}

func (fake *FakeNodeProvider) KillNodeIfExistsCalls(stub func(string) error) {
	fake.killNodeIfExistsMutex.Lock()
	defer fake.killNodeIfExistsMutex.Unlock()
	fake.KillNodeIfExistsStub = stub
}

func (fake *FakeNodeProvider) KillNodeIfExistsArgsForCall(i int) string {
	fake.killNodeIfExistsMutex.RLock()
	defer fake.killNodeIfExistsMutex.RUnlock()
	argsForCall := fake.killNodeIfExistsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNodeProvider) KillNodeIfExistsReturns(result1 error) {
	fake.killNodeIfExistsMutex.Lock()
	defer fake.killNodeIfExistsMutex.Unlock()
	fake.KillNodeIfExistsStub = nil
	fake.killNodeIfExistsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNodeProvider) KillNodeIfExistsReturnsOnCall(i int, result1 error) {
	fake.killNodeIfExistsMutex.Lock()
	defer fake.killNodeIfExistsMutex.Unlock()
	fake.KillNodeIfExistsStub = nil
	if fake.killNodeIfExistsReturnsOnCall == nil {
		fake.killNodeIfExistsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.killNodeIfExistsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNodeProvider) ListNodes() ([]model.NodeHandle, error) {
	fake.listNodesMutex.Lock()
	ret, specificReturn := fake.listNodesReturnsOnCall[len(fake.listNodesArgsForCall)]
	fake.listNodesArgsForCall = append(fake.listNodesArgsForCall, struct {
	}{})
	fake.recordInvocation("ListNodes", []interface{}{})
	fake.listNodesMutex.Unlock()
	if fake.ListNodesStub != nil {
		return fake.ListNodesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listNodesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNodeProvider) ListNodesCallCount() int {
	fake.listNodesMutex.RLock()
	defer fake.listNodesMutex.RUnlock()
	return len(fake.listNodesArgsForCall)
}

func (fake *FakeNodeProvider) ListNodesCalls(stub func() ([]model.NodeHandle, error)) {
	fake.listNodesMutex.Lock()
	defer fake.listNodesMutex.Unlock()
	fake.ListNodesStub = stub
}

func (fake *FakeNodeProvider) ListNodesReturns(result1 []model.NodeHandle, result2 error) {
	fake.listNodesMutex.Lock()
	defer fake.listNodesMutex.Unlock()
	fake.ListNodesStub = nil
	fake.listNodesReturns = struct {
		result1 []model.NodeHandle
		result2 error
	}{result1, result2}
}

func (fake *FakeNodeProvider) ListNodesReturnsOnCall(i int, result1 []model.NodeHandle, result2 error) {
	fake.listNodesMutex.Lock()
	defer fake.listNodesMutex.Unlock()
	fake.ListNodesStub = nil
	if fake.listNodesReturnsOnCall == nil {
		fake.listNodesReturnsOnCall = make(map[int]struct {
			result1 []model.NodeHandle
			result2 error
		})
	}
	fake.listNodesReturnsOnCall[i] = struct {
		result1 []model.NodeHandle
		result2 error
	}{result1, result2}
}

func (fake *FakeNodeProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createNodeIfNotExistsMutex.RLock()
	defer fake.createNodeIfNotExistsMutex.RUnlock()
	fake.killNodeIfExistsMutex.RLock()
	defer fake.killNodeIfExistsMutex.RUnlock()
	fake.listNodesMutex.RLock()
	defer fake.listNodesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNodeProvider) recordInvocation(key string, args []interface{}) {
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

var _ nodeprovider.NodeProvider = new(FakeNodeProvider)
