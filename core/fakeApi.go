// This file was generated by counterfeiter
package core

import (
  "sync"

  "github.com/dev-op-spec/engine/core/models"
)

type FakeApi struct {
  AddOperationStub                     func(req models.AddOperationReq) (err error)
  addOperationMutex                    sync.RWMutex
  addOperationArgsForCall              []struct {
    req models.AddOperationReq
  }
  addOperationReturns                  struct {
                                         result1 error
                                       }
  AddSubOperationStub                  func(req models.AddSubOperationReq) (err error)
  addSubOperationMutex                 sync.RWMutex
  addSubOperationArgsForCall           []struct {
    req models.AddSubOperationReq
  }
  addSubOperationReturns               struct {
                                         result1 error
                                       }
  ListOperationsStub                   func(projectUrl *models.Url) (operations []models.OperationDetailedView, err error)
  listOperationsMutex                  sync.RWMutex
  listOperationsArgsForCall            []struct {
    projectUrl *models.Url
  }
  listOperationsReturns                struct {
                                         result1 []models.OperationDetailedView
                                         result2 error
                                       }
  RunOperationStub                     func(req models.RunOperationReq) (operationRun models.OperationRunDetailedView, err error)
  runOperationMutex                    sync.RWMutex
  runOperationArgsForCall              []struct {
    req models.RunOperationReq
  }
  runOperationReturns                  struct {
                                         result1 models.OperationRunDetailedView
                                         result2 error
                                       }
  SetDescriptionOfOperationStub        func(req models.SetDescriptionOfOperationReq) (err error)
  setDescriptionOfOperationMutex       sync.RWMutex
  setDescriptionOfOperationArgsForCall []struct {
    req models.SetDescriptionOfOperationReq
  }
  setDescriptionOfOperationReturns     struct {
                                         result1 error
                                       }
}

func (fake *FakeApi) AddOperation(req models.AddOperationReq) (err error) {
  fake.addOperationMutex.Lock()
  fake.addOperationArgsForCall = append(fake.addOperationArgsForCall, struct {
    req models.AddOperationReq
  }{req})
  fake.addOperationMutex.Unlock()
  if fake.AddOperationStub != nil {
    return fake.AddOperationStub(req)
  } else {
    return fake.addOperationReturns.result1
  }
}

func (fake *FakeApi) AddOperationCallCount() int {
  fake.addOperationMutex.RLock()
  defer fake.addOperationMutex.RUnlock()
  return len(fake.addOperationArgsForCall)
}

func (fake *FakeApi) AddOperationArgsForCall(i int) models.AddOperationReq {
  fake.addOperationMutex.RLock()
  defer fake.addOperationMutex.RUnlock()
  return fake.addOperationArgsForCall[i].req
}

func (fake *FakeApi) AddOperationReturns(result1 error) {
  fake.AddOperationStub = nil
  fake.addOperationReturns = struct {
    result1 error
  }{result1}
}

func (fake *FakeApi) AddSubOperation(req models.AddSubOperationReq) (err error) {
  fake.addSubOperationMutex.Lock()
  fake.addSubOperationArgsForCall = append(fake.addSubOperationArgsForCall, struct {
    req models.AddSubOperationReq
  }{req})
  fake.addSubOperationMutex.Unlock()
  if fake.AddSubOperationStub != nil {
    return fake.AddSubOperationStub(req)
  } else {
    return fake.addSubOperationReturns.result1
  }
}

func (fake *FakeApi) AddSubOperationCallCount() int {
  fake.addSubOperationMutex.RLock()
  defer fake.addSubOperationMutex.RUnlock()
  return len(fake.addSubOperationArgsForCall)
}

func (fake *FakeApi) AddSubOperationArgsForCall(i int) models.AddSubOperationReq {
  fake.addSubOperationMutex.RLock()
  defer fake.addSubOperationMutex.RUnlock()
  return fake.addSubOperationArgsForCall[i].req
}

func (fake *FakeApi) AddSubOperationReturns(result1 error) {
  fake.AddSubOperationStub = nil
  fake.addSubOperationReturns = struct {
    result1 error
  }{result1}
}

func (fake *FakeApi) ListOperations(projectUrl *models.Url) (operations []models.OperationDetailedView, err error) {
  fake.listOperationsMutex.Lock()
  fake.listOperationsArgsForCall = append(fake.listOperationsArgsForCall, struct {
    projectUrl *models.Url
  }{projectUrl})
  fake.listOperationsMutex.Unlock()
  if fake.ListOperationsStub != nil {
    return fake.ListOperationsStub(projectUrl)
  } else {
    return fake.listOperationsReturns.result1, fake.listOperationsReturns.result2
  }
}

func (fake *FakeApi) ListOperationsCallCount() int {
  fake.listOperationsMutex.RLock()
  defer fake.listOperationsMutex.RUnlock()
  return len(fake.listOperationsArgsForCall)
}

func (fake *FakeApi) ListOperationsArgsForCall(i int) *models.Url {
  fake.listOperationsMutex.RLock()
  defer fake.listOperationsMutex.RUnlock()
  return fake.listOperationsArgsForCall[i].projectUrl
}

func (fake *FakeApi) ListOperationsReturns(result1 []models.OperationDetailedView, result2 error) {
  fake.ListOperationsStub = nil
  fake.listOperationsReturns = struct {
    result1 []models.OperationDetailedView
    result2 error
  }{result1, result2}
}

func (fake *FakeApi) RunOperation(req models.RunOperationReq) (operationRun models.OperationRunDetailedView, err error) {
  fake.runOperationMutex.Lock()
  fake.runOperationArgsForCall = append(fake.runOperationArgsForCall, struct {
    req models.RunOperationReq
  }{req})
  fake.runOperationMutex.Unlock()
  if fake.RunOperationStub != nil {
    return fake.RunOperationStub(req)
  } else {
    return fake.runOperationReturns.result1, fake.runOperationReturns.result2
  }
}

func (fake *FakeApi) RunOperationCallCount() int {
  fake.runOperationMutex.RLock()
  defer fake.runOperationMutex.RUnlock()
  return len(fake.runOperationArgsForCall)
}

func (fake *FakeApi) RunOperationArgsForCall(i int) models.RunOperationReq {
  fake.runOperationMutex.RLock()
  defer fake.runOperationMutex.RUnlock()
  return fake.runOperationArgsForCall[i].req
}

func (fake *FakeApi) RunOperationReturns(result1 models.OperationRunDetailedView, result2 error) {
  fake.RunOperationStub = nil
  fake.runOperationReturns = struct {
    result1 models.OperationRunDetailedView
    result2 error
  }{result1, result2}
}

func (fake *FakeApi) SetDescriptionOfOperation(req models.SetDescriptionOfOperationReq) (err error) {
  fake.setDescriptionOfOperationMutex.Lock()
  fake.setDescriptionOfOperationArgsForCall = append(fake.setDescriptionOfOperationArgsForCall, struct {
    req models.SetDescriptionOfOperationReq
  }{req})
  fake.setDescriptionOfOperationMutex.Unlock()
  if fake.SetDescriptionOfOperationStub != nil {
    return fake.SetDescriptionOfOperationStub(req)
  } else {
    return fake.setDescriptionOfOperationReturns.result1
  }
}

func (fake *FakeApi) SetDescriptionOfOperationCallCount() int {
  fake.setDescriptionOfOperationMutex.RLock()
  defer fake.setDescriptionOfOperationMutex.RUnlock()
  return len(fake.setDescriptionOfOperationArgsForCall)
}

func (fake *FakeApi) SetDescriptionOfOperationArgsForCall(i int) models.SetDescriptionOfOperationReq {
  fake.setDescriptionOfOperationMutex.RLock()
  defer fake.setDescriptionOfOperationMutex.RUnlock()
  return fake.setDescriptionOfOperationArgsForCall[i].req
}

func (fake *FakeApi) SetDescriptionOfOperationReturns(result1 error) {
  fake.SetDescriptionOfOperationStub = nil
  fake.setDescriptionOfOperationReturns = struct {
    result1 error
  }{result1}
}

var _ Api = new(FakeApi)
