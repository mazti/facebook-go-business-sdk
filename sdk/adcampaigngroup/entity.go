package adsinsights

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

	BudgetLimitNew          interface{} `json:"budget_limit_new"`
	BudgetLimitOld          interface{} `json:"budget_limit_old"`
	BuyingTypeNew           string      `json:"buying_type_new"`
	BuyingTypeOld           string      `json:"buying_type_old"`
	EventTime               string      `json:"event_time"`
	EventType               string      `json:"event_type"`
	ID                      string      `json:"id"`
	IsAutobidNew            bool        `json:"is_autobid_new"`
	IsAutobidOld            bool        `json:"is_autobid_old"`
	IsAveragePricePacingNew bool        `json:"is_average_price_pacing_new"`
	IsAveragePricePacingOld bool        `json:"is_average_price_pacing_old"`
	NameNew                 string      `json:"name_new"`
	NameOld                 string      `json:"name_old"`
	ObjectiveNew            interface{} `json:"objective_new"`
	ObjectiveOld            interface{} `json:"objective_old"`
	PacingType              int64       `json:"pacing_type"`
	RunStatusNew            string      `json:"run_status_new"`
	RunStatusOld            string      `json:"run_status_old"`
	SpendCapNew             int64       `json:"spend_cap_new"`
	SpendCapOld             int64       `json:"spend_cap_old"`
	TimeCreated             string      `json:"time_created"`
	TimeUpdatedNew          string      `json:"time_updated_new"`
	TimeUpdatedOld          string      `json:"time_updated_old"`
	// TODO: check & correct field for interface
}

func CreateAPIRequestGet(id string, context *sdk.APIContext) *sdk.APIRequest {
	return sdk.NewAPIRequest(
		context,
		id,
		endpoint,
		http.MethodGet,
		sdk.Parser(parserResponse),
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
	ent.BudgetLimitNew = other.BudgetLimitNew
	ent.BudgetLimitOld = other.BudgetLimitOld
	ent.BuyingTypeNew = other.BuyingTypeNew
	ent.BuyingTypeOld = other.BuyingTypeOld
	ent.EventTime = other.EventTime
	ent.EventType = other.EventType
	ent.ID = other.ID
	ent.IsAutobidNew = other.IsAutobidNew
	ent.IsAutobidOld = other.IsAutobidOld
	ent.IsAveragePricePacingNew = other.IsAveragePricePacingNew
	ent.IsAveragePricePacingOld = other.IsAveragePricePacingOld
	ent.NameNew = other.NameNew
	ent.NameOld = other.NameOld
	ent.ObjectiveNew = other.ObjectiveNew
	ent.ObjectiveOld = other.ObjectiveOld
	ent.PacingType = other.PacingType
	ent.RunStatusNew = other.RunStatusNew
	ent.RunStatusOld = other.RunStatusOld
	ent.SpendCapNew = other.SpendCapNew
	ent.SpendCapOld = other.SpendCapOld
	ent.TimeCreated = other.TimeCreated
	ent.TimeUpdatedNew = other.TimeUpdatedNew
	ent.TimeUpdatedOld = other.TimeUpdatedOld
}
