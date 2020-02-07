package adsinsights

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

func (ent AdsInsights) GetBody() []byte {
	return ent.node.GetBody()
}

func (ent AdsInsights) GetHeader() []byte {
	return ent.node.GetBody()
}

func (ent *AdsInsights) SetRequest(request *sdk.APIRequest) {
	ent.node.SetRequest(request)
}

func (ent *AdsInsights) SetBody(body []byte) {
	ent.node.SetBody(body)
}

func (ent *AdsInsights) SetHeader(header []byte) {
	ent.node.SetHeader(header)
}
