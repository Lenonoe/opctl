// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

type FakeHostConfigFactory struct {
	ConstructStub        func(map[string]string, map[string]string, map[string]string, nat.PortMap) *container.HostConfig
	constructMutex       sync.RWMutex
	constructArgsForCall []struct {
		arg1 map[string]string
		arg2 map[string]string
		arg3 map[string]string
		arg4 nat.PortMap
	}
	constructReturns struct {
		result1 *container.HostConfig
	}
	constructReturnsOnCall map[int]struct {
		result1 *container.HostConfig
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHostConfigFactory) Construct(arg1 map[string]string, arg2 map[string]string, arg3 map[string]string, arg4 nat.PortMap) *container.HostConfig {
	fake.constructMutex.Lock()
	ret, specificReturn := fake.constructReturnsOnCall[len(fake.constructArgsForCall)]
	fake.constructArgsForCall = append(fake.constructArgsForCall, struct {
		arg1 map[string]string
		arg2 map[string]string
		arg3 map[string]string
		arg4 nat.PortMap
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Construct", []interface{}{arg1, arg2, arg3, arg4})
	fake.constructMutex.Unlock()
	if fake.ConstructStub != nil {
		return fake.ConstructStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.constructReturns
	return fakeReturns.result1
}

func (fake *FakeHostConfigFactory) ConstructCallCount() int {
	fake.constructMutex.RLock()
	defer fake.constructMutex.RUnlock()
	return len(fake.constructArgsForCall)
}

func (fake *FakeHostConfigFactory) ConstructCalls(stub func(map[string]string, map[string]string, map[string]string, nat.PortMap) *container.HostConfig) {
	fake.constructMutex.Lock()
	defer fake.constructMutex.Unlock()
	fake.ConstructStub = stub
}

func (fake *FakeHostConfigFactory) ConstructArgsForCall(i int) (map[string]string, map[string]string, map[string]string, nat.PortMap) {
	fake.constructMutex.RLock()
	defer fake.constructMutex.RUnlock()
	argsForCall := fake.constructArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeHostConfigFactory) ConstructReturns(result1 *container.HostConfig) {
	fake.constructMutex.Lock()
	defer fake.constructMutex.Unlock()
	fake.ConstructStub = nil
	fake.constructReturns = struct {
		result1 *container.HostConfig
	}{result1}
}

func (fake *FakeHostConfigFactory) ConstructReturnsOnCall(i int, result1 *container.HostConfig) {
	fake.constructMutex.Lock()
	defer fake.constructMutex.Unlock()
	fake.ConstructStub = nil
	if fake.constructReturnsOnCall == nil {
		fake.constructReturnsOnCall = make(map[int]struct {
			result1 *container.HostConfig
		})
	}
	fake.constructReturnsOnCall[i] = struct {
		result1 *container.HostConfig
	}{result1}
}

func (fake *FakeHostConfigFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.constructMutex.RLock()
	defer fake.constructMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHostConfigFactory) recordInvocation(key string, args []interface{}) {
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
