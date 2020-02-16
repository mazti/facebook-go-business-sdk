package adsinsights

import (
	"time"

	"github.com/mazti/facebook-go-business-sdk/sdk"
)

type APIRequestGetInsights struct {
	request *sdk.APIRequest
}

func CreateAPIRequestGetInsights(nodeID string, context *sdk.APIContext) *APIRequestGetInsights {
	request := CreateAPIRequestGet(nodeID, context)

	return &APIRequestGetInsights{
		request: request,
	}
}

func (r *APIRequestGetInsights) SetTimeRange(since time.Time, until time.Time) *APIRequestGetInsights {
	r.request.SetParam(sdk.TimeRangeKey, sdk.CreateTimeRange(since, until))
	return r
}

func (r *APIRequestGetInsights) Execute() (*sdk.APINodeList, error) {
	resp, err := r.request.Execute()
	if err != nil {
		return nil, err
	}
	return resp.(*sdk.APINodeList), nil
}
