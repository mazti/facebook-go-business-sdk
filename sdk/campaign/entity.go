package campaign

import (
	"net/http"

	"github.com/mazti/facebook-go-business-sdk/sdk/adsinsights"

	"github.com/mazti/facebook-go-business-sdk/sdk"
	"github.com/mazti/facebook-go-business-sdk/sdk/adcampaignissueinfo"
	"github.com/mazti/facebook-go-business-sdk/sdk/adlabel"
	"github.com/mazti/facebook-go-business-sdk/sdk/adpromotedobject"
	"github.com/mazti/facebook-go-business-sdk/sdk/adrecommendation"
	"github.com/mazti/facebook-go-business-sdk/sdk/adstudy"
)

const (
	endpoint = "campaigns"
)

// https://developers.facebook.com/docs/marketing-api/reference/ad-account/campaigns/
type Entity struct {
	request *sdk.APIRequest

	AccountId               string                       `json:"account_id"`
	Adlabels                []adlabel.Entity             `json:"adlabels"`
	BidStrategy             BidStrategy                  `json:"bid_strategy"`
	BoostedObjectId         string                       `json:"boosted_object_id"`
	BrandLiftStudies        []adstudy.Entity             `json:"brand_lift_studies"`
	BudgetRebalanceFlag     bool                         `json:"budget_rebalance_flag"`
	BudgetRemaining         string                       `json:"budget_remaining"`
	BuyingType              string                       `json:"buying_type"`
	CanCreateBrandLiftStudy bool                         `json:"can_create_brand_lift_study"`
	CanUseSpendCap          bool                         `json:"can_use_spend_cap"`
	ConfiguredStatus        ConfiguredStatus             `json:"configured_status"`
	CreatedTime             string                       `json:"created_time"`
	DailyBudget             string                       `json:"daily_budget"`
	EffectiveStatus         EffectiveStatus              `json:"effective_status"`
	ID                      string                       `json:"id"`
	IssuesInfo              []adcampaignissueinfo.Entity `json:"issues_info"`
	LastBudgetTogglingTime  string                       `json:"last_budget_toggling_time"`
	LifetimeBudget          string                       `json:"lifetime_budget"`
	Name                    string                       `json:"name"`
	Objective               string                       `json:"objective"`
	PacingType              []string                     `json:"pacing_type"`
	PromotedObject          adpromotedobject.Entity      `json:"promoted_object"`
	Recommendations         []adrecommendation.Entity    `json:"recommendations"`
	SourceEntity            *Entity                      `json:"source_campaign"`
	SourceEntityId          string                       `json:"source_campaign_id"`
	SpecialAdCategory       string                       `json:"special_ad_category"`
	SpendCap                string                       `json:"spend_cap"`
	StartTime               string                       `json:"start_time"`
	Status                  Status                       `json:"status"`
	StopTime                string                       `json:"stop_time"`
	ToplineID               string                       `json:"topline_id"`
	UpdatedTime             string                       `json:"updated_time"`
}

func ParseResponse(rawResp sdk.APIResponse) (resp []Entity, err error) {
	request := rawResp.GetRequest()
	context := request.Context
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
		resp[i].SetRequest(request)
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

func (ent *Entity) GetInsights() *adsinsights.APIRequestGetInsights {
	return adsinsights.CreateAPIRequestGetInsights(
		ent.getPrefixID(),
		ent.GetRequest().Context,
	)
}

func (ent *Entity) getPrefixID() string {
	return ent.ID
}

func (ent *Entity) GetAds() (*sdk.APINodeList, error) {
	return nil, nil
}

func (ent *Entity) GetAdSets() (*sdk.APINodeList, error) {
	return nil, nil
}
