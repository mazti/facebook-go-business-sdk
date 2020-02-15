package adset

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/mazti/facebook-go-business-sdk/sdk"
)

const (
	endpoint = "/"
)

type AdSet struct {
	Context *sdk.APIContext

	AccountId                    string           `json:"account_id"`
	Adlabels                     []interface{}    `json:"adlabels"`
	AdsetSchedule                []interface{}    `json:"adset_schedule"`
	AssetFeedId                  string           `json:"asset_feed_id"`
	AttributionSpec              []interface{}    `json:"attribution_spec"`
	BidAdjustments               interface{}      `json:"bid_adjustments"`
	BidAmount                    int64            `json:"bid_amount"`
	BidConstraints               interface{}      `json:"bid_constraints"`
	BidInfo                      map[string]int64 `json:"bid_info"`
	BidStrategy                  interface{}      `json:"bid_strategy"`
	BillingEvent                 interface{}      `json:"billing_event"`
	BudgetRemaining              string           `json:"budget_remaining"`
	Campaign                     interface{}      `json:"campaign"`
	CampaignId                   string           `json:"campaign_id"`
	ConfiguredStatus             interface{}      `json:"configured_status"`
	CreatedTime                  string           `json:"created_time"`
	CreativeSequence             []interface{}    `json:"creative_sequence"`
	DailyBudget                  string           `json:"daily_budget"`
	DailyMinSpendTarget          string           `json:"daily_min_spend_target"`
	DailySpendCap                string           `json:"daily_spend_cap"`
	DestinationType              string           `json:"destination_type"`
	EffectiveStatus              interface{}      `json:"effective_status"`
	EndTime                      string           `json:"end_time"`
	FrequencyControlSpecs        []interface{}    `json:"frequency_control_specs"`
	FullFunnelExplorationMode    string           `json:"full_funnel_exploration_mode"`
	ID                           string           `json:"id"`
	InstagramActorId             string           `json:"instagram_actor_id"`
	IsDynamicCreative            bool             `json:"is_dynamic_creative"`
	IssuesInfo                   []interface{}    `json:"issues_info"`
	LifetimeBudget               string           `json:"lifetime_budget"`
	LifetimeImps                 int64            `json:"lifetime_imps"`
	LifetimeMinSpendTarget       string           `json:"lifetime_min_spend_target"`
	LifetimeSpendCap             string           `json:"lifetime_spend_cap"`
	Name                         string           `json:"name"`
	OptimizationGoal             interface{}      `json:"optimization_goal"`
	OptimizationSubEvent         string           `json:"optimization_sub_event"`
	PacingType                   []interface{}    `json:"pacing_type"`
	PromotedObject               interface{}      `json:"promoted_object"`
	Recommendations              []interface{}    `json:"recommendations"`
	RecurringBudgetSemantics     bool             `json:"recurring_budget_semantics"`
	ReviewFeedback               string           `json:"review_feedback"`
	RfPredictionId               string           `json:"rf_prediction_id"`
	SourceAdset                  interface{}      `json:"source_adset"`
	SourceAdsetId                string           `json:"source_adset_id"`
	StartTime                    string           `json:"start_time"`
	Status                       interface{}      `json:"status"`
	Targeting                    interface{}      `json:"targeting"`
	TimeBasedAdRotationIdBlocks  []interface{}    `json:"time_based_ad_rotation_id_blocks"`
	TimeBasedAdRotationIntervals []interface{}    `json:"time_based_ad_rotation_intervals"`
	UpdatedTime                  string           `json:"updated_time"`
	UseNewAppClick               bool             `json:"use_new_app_click"`
	// TODO: check & correct field for interface
}

func CreateAPIRequestGet(id string, context *sdk.APIContext) *sdk.APIRequest {
	return sdk.NewAPIRequest(
		context,
		id,
		endpoint,
		http.MethodGet,
		sdk.Parser(parserResponse),
		sdk.ParamNames(params),
		sdk.ReturnFields(fields),
	)
}

func FetchByID(id string, context *sdk.APIContext) (*AdSet, error) {
	req := sdk.NewAPIRequest(
		context,
		id,
		sdk.DefaultEndpoint,
		http.MethodGet,
		sdk.Parser(parserResponse),
		sdk.ReturnFields(fields),
	)
	ent, err := req.Execute()
	if err != nil {
		return nil, err
	}

	account := ent.(*AdSet)
	account.ID = strings.Replace(account.ID, "act_", "", 1)

	return account, nil
}

func (ent *AdSet) Fetch() (*AdSet, error) {
	obj, err := FetchByID(ent.getPrefixID(), ent.Context)
	if err != nil {
		return nil, err
	}
	ent.copy(obj)
	return ent, nil
}

//
// For internal usage
//
func parserResponse(data json.RawMessage) (sdk.APIResponse, error) {
	ent := &AdSet{}
	if err := json.Unmarshal(data, ent); err != nil {
		return ent, err
	}
	return ent, nil
}

func (ent *AdSet) getPrefixID() string {
	return ent.ID
}

func (ent *AdSet) copy(other *AdSet) {
	// TODO: copy fields
}
