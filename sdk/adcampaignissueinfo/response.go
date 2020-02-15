package adcampaignissueinfo

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

func (ent *AdCampaignIssueInfo) SetContext(context *sdk.APIContext) {
	ent.Context = context
}

func (ent *AdCampaignIssueInfo) GetContext() *sdk.APIContext {
	return ent.Context
}
