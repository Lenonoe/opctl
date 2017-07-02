// Code generated by counterfeiter. DO NOT EDIT.
package pubsub

import (
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type FakeEventSubscriber struct {
	SubscribeStub        func(filter *model.EventFilter, eventChannel chan *model.Event)
	subscribeMutex       sync.RWMutex
	subscribeArgsForCall []struct {
		filter       *model.EventFilter
		eventChannel chan *model.Event
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEventSubscriber) Subscribe(filter *model.EventFilter, eventChannel chan *model.Event) {
	fake.subscribeMutex.Lock()
	fake.subscribeArgsForCall = append(fake.subscribeArgsForCall, struct {
		filter       *model.EventFilter
		eventChannel chan *model.Event
	}{filter, eventChannel})
	fake.recordInvocation("Subscribe", []interface{}{filter, eventChannel})
	fake.subscribeMutex.Unlock()
	if fake.SubscribeStub != nil {
		fake.SubscribeStub(filter, eventChannel)
	}
}

func (fake *FakeEventSubscriber) SubscribeCallCount() int {
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	return len(fake.subscribeArgsForCall)
}

func (fake *FakeEventSubscriber) SubscribeArgsForCall(i int) (*model.EventFilter, chan *model.Event) {
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	return fake.subscribeArgsForCall[i].filter, fake.subscribeArgsForCall[i].eventChannel
}

func (fake *FakeEventSubscriber) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEventSubscriber) recordInvocation(key string, args []interface{}) {
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

var _ EventSubscriber = new(FakeEventSubscriber)
