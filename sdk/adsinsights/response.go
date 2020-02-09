package adsinsights

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

func (ent *AdsInsights) Load(context *sdk.APIContext, req *sdk.APIRequest, header []byte, body []byte) {
	ent.node.Load(context, req, header, body)
}

func (ent *AdsInsights) SetContext(context *sdk.APIContext) {
	ent.node.SetContext(context)
}
