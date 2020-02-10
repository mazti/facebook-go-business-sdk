package adsinsights

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

func (ent *AdsInsights) SetContext(context *sdk.APIContext) {
	ent.Context = context
}

func (ent *AdsInsights) GetContext() *sdk.APIContext {
	return ent.Context
}
