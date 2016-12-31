package engineclient

//go:generate counterfeiter -o ./fakeEngineClient.go --fake-name FakeEngineClient ./ EngineClient

import (
  "github.com/opspec-io/sdk-golang/pkg/model"
  "github.com/opspec-io/sdk-golang/pkg/engineprovider"
  "github.com/opspec-io/sdk-golang/util/format"
  "github.com/opspec-io/sdk-golang/util/http"
  "github.com/sethgrid/pester"
)

type EngineClient interface {
  GetEventStream(
  ) (
  stream chan model.Event,
  err error,
  )

  KillOp(
  req model.KillOpReq,
  ) (
  err error,
  )

  StartOp(
  req model.StartOpReq,
  ) (
  opId string,
  err error,
  )
}

func New(
engineProvider engineprovider.EngineProvider,
) EngineClient {

  httpClient := pester.New()
  httpClient.Backoff = pester.ExponentialBackoff

  return &_engineClient{
    engineProvider:engineProvider,
    httpClient:httpClient,
    jsonFormat: format.NewJsonFormat(),
  }
}

type _engineClient struct {
  engineProvider engineprovider.EngineProvider
  httpClient     http.Client
  jsonFormat     format.Format
}
