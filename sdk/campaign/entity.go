package campaign

import (
	"net/http"

	"github.com/mazti/facebook-go-business-sdk/sdk"
	"github.com/mazti/facebook-go-business-sdk/sdk/adcampaignissueinfo"
	"github.com/mazti/facebook-go-business-sdk/sdk/adlabel"
	"github.com/mazti/facebook-go-business-sdk/sdk/adpromotedobject"
	"github.com/mazti/facebook-go-business-sdk/sdk/adrecommendation"
	"github.com/mazti/facebook-go-business-sdk/sdk/adstudy"
)

const (
	endpoint = "/campaigns"
)

type Campaign struct {
	Context *sdk.APIContext

	AccountId               string                                    `json:"account_id"`
	Adlabels                []adlabel.AdLabel                         `json:"adlabels"`
	BidStrategy             BidStrategy                               `json:"bid_strategy"`
	BoostedObjectId         string                                    `json:"boosted_object_id"`
	BrandLiftStudies        []adstudy.AdStudy                         `json:"brand_lift_studies"`
	BudgetRebalanceFlag     bool                                      `json:"budget_rebalance_flag"`
	BudgetRemaining         string                                    `json:"budget_remaining"`
	BuyingType              string                                    `json:"buying_type"`
	CanCreateBrandLiftStudy bool                                      `json:"can_create_brand_lift_study"`
	CanUseSpendCap          bool                                      `json:"can_use_spend_cap"`
	ConfiguredStatus        ConfiguredStatus                          `json:"configured_status"`
	CreatedTime             string                                    `json:"created_time"`
	DailyBudget             string                                    `json:"daily_budget"`
	EffectiveStatus         EffectiveStatus                           `json:"effective_status"`
	ID                      string                                    `json:"id"`
	IssuesInfo              []adcampaignissueinfo.AdCampaignIssueInfo `json:"issues_info"`
	LastBudgetTogglingTime  string                                    `json:"last_budget_toggling_time"`
	LifetimeBudget          string                                    `json:"lifetime_budget"`
	Name                    string                                    `json:"name"`
	Objective               string                                    `json:"objective"`
	PacingType              []string                                  `json:"pacing_type"`
	PromotedObject          adpromotedobject.AdPromotedObject         `json:"promoted_object"`
	Recommendations         []adrecommendation.AdRecommendation       `json:"recommendations"`
	SourceCampaign          *Campaign                                 `json:"source_campaign"`
	SourceCampaignId        string                                    `json:"source_campaign_id"`
	SpecialAdCategory       string                                    `json:"special_ad_category"`
	SpendCap                string                                    `json:"spend_cap"`
	StartTime               string                                    `json:"start_time"`
	Status                  Status                                    `json:"status"`
	StopTime                string                                    `json:"stop_time"`
	ToplineID               string                                    `json:"topline_id"`
	UpdatedTime             string                                    `json:"updated_time"`
}

func ParseResponse(rawResp sdk.APIResponse) (resp []Campaign, err error) {
	context := rawResp.GetContext()
	nodeList, ok := rawResp.(*sdk.APINodeList)
	if !ok {
		return nil, sdk.UnsupportedResponse
	}
	err = nodeList.Unmarshal(&resp)
	if err != nil {
		context.Log(err)
		return
	}
	for i := 0; i < len(resp); i++ {
		resp[i].SetContext(context)
	}
	return
}

func CreateAPIRequestGet(id string, context *sdk.APIContext) *sdk.APIRequest {
	return sdk.NewAPIRequest(
		context,
		id,
		endpoint,
		http.MethodGet,
		sdk.Parser(sdk.ParserResponse),
		sdk.ParamNames(params),
		sdk.ReturnFields(fields),
	)
}
