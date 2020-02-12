package adsinsights

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

func (ent *AdCampaignGroupActivity) SetContext(context *sdk.APIContext) {
	ent.Context = context
}

func (ent *AdCampaignGroupActivity) GetContext() *sdk.APIContext {
	return ent.Context
}
