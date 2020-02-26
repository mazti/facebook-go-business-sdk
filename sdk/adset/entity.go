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

type Entity struct {
	request *sdk.APIRequest

	AccountId                    string           `json:"account_id"`
	Adlabels                     json.RawMessage  `json:"adlabels"`
	EntitySchedule               json.RawMessage  `json:"Entity_schedule"`
	AssetFeedId                  string           `json:"asset_feed_id"`
	AttributionSpec              json.RawMessage  `json:"attribution_spec"`
	BidAdjustments               json.RawMessage  `json:"bid_adjustments"`
	BidAmount                    int64            `json:"bid_amount"`
	BidConstraints               json.RawMessage  `json:"bid_constraints"`
	BidInfo                      map[string]int64 `json:"bid_info"`
	BidStrategy                  json.RawMessage  `json:"bid_strategy"`
	BillingEvent                 json.RawMessage  `json:"billing_event"`
	BudgetRemaining              string           `json:"budget_remaining"`
	Campaign                     json.RawMessage  `json:"campaign"`
	CampaignId                   string           `json:"campaign_id"`
	ConfiguredStatus             json.RawMessage  `json:"configured_status"`
	CreatedTime                  string           `json:"created_time"`
	CreativeSequence             json.RawMessage  `json:"creative_sequence"`
	DailyBudget                  string           `json:"daily_budget"`
	DailyMinSpendTarget          string           `json:"daily_min_spend_target"`
	DailySpendCap                string           `json:"daily_spend_cap"`
	DestinationType              string           `json:"destination_type"`
	EffectiveStatus              json.RawMessage  `json:"effective_status"`
	EndTime                      string           `json:"end_time"`
	FrequencyControlSpecs        json.RawMessage  `json:"frequency_control_specs"`
	FullFunnelExplorationMode    string           `json:"full_funnel_exploration_mode"`
	ID                           string           `json:"id"`
	InstagramActorId             string           `json:"instagram_actor_id"`
	IsDynamicCreative            bool             `json:"is_dynamic_creative"`
	IssuesInfo                   json.RawMessage  `json:"issues_info"`
	LifetimeBudget               string           `json:"lifetime_budget"`
	LifetimeImps                 int64            `json:"lifetime_imps"`
	LifetimeMinSpendTarget       string           `json:"lifetime_min_spend_target"`
	LifetimeSpendCap             string           `json:"lifetime_spend_cap"`
	Name                         string           `json:"name"`
	OptimizationGoal             json.RawMessage  `json:"optimization_goal"`
	OptimizationSubEvent         string           `json:"optimization_sub_event"`
	PacingType                   json.RawMessage  `json:"pacing_type"`
	PromotedObject               json.RawMessage  `json:"promoted_object"`
	Recommendations              json.RawMessage  `json:"recommendations"`
	RecurringBudgetSemantics     bool             `json:"recurring_budget_semantics"`
	ReviewFeedback               string           `json:"review_feedback"`
	RfPredictionId               string           `json:"rf_prediction_id"`
	SourceEntity                 json.RawMessage  `json:"source_Entity"`
	SourceEntityId               string           `json:"source_Entity_id"`
	StartTime                    string           `json:"start_time"`
	Status                       json.RawMessage  `json:"status"`
	Targeting                    json.RawMessage  `json:"targeting"`
	TimeBasedAdRotationIdBlocks  json.RawMessage  `json:"time_based_ad_rotation_id_blocks"`
	TimeBasedAdRotationIntervals json.RawMessage  `json:"time_based_ad_rotation_intervals"`
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

func FetchByID(id string, context *sdk.APIContext) (*Entity, error) {
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

	account := ent.(*Entity)
	account.ID = strings.Replace(account.ID, "act_", "", 1)

	return account, nil
}

func (ent *Entity) Fetch() (*Entity, error) {
	obj, err := FetchByID(ent.getPrefixID(), ent.GetRequest().Context)
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
	ent := &Entity{}
	if err := json.Unmarshal(data, ent); err != nil {
		return ent, err
	}
	return ent, nil
}

func (ent *Entity) getPrefixID() string {
	return ent.ID
}

func (ent *Entity) copy(other *Entity) {
	// TODO: copy fields
}
