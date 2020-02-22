package adaccount

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/mazti/facebook-go-business-sdk/sdk"
	"github.com/mazti/facebook-go-business-sdk/sdk/adaccountcreationrequest"
	"github.com/mazti/facebook-go-business-sdk/sdk/adaccountpromotableobjects"
	"github.com/mazti/facebook-go-business-sdk/sdk/adsinsights"
	"github.com/mazti/facebook-go-business-sdk/sdk/agencyclientdeclaration"
	"github.com/mazti/facebook-go-business-sdk/sdk/attributionspec"
	"github.com/mazti/facebook-go-business-sdk/sdk/business"
	"github.com/mazti/facebook-go-business-sdk/sdk/campaign"
	"github.com/mazti/facebook-go-business-sdk/sdk/deliverycheck"
	"github.com/mazti/facebook-go-business-sdk/sdk/extendedcreditinvoicegroup"
	"github.com/mazti/facebook-go-business-sdk/sdk/fundingsourcedetails"
	"github.com/mazti/facebook-go-business-sdk/sdk/reachfrequencyspec"
)

const (
	nodeID   = "me"
	endpoint = "adaccounts"
)

type Entity struct {
	request *sdk.APIRequest

	AccountID                         string                            `json:"account_id"`
	AccountStatus                     int64                             `json:"account_status"`
	AdAccountCreationRequest          adaccountcreationrequest.Entity   `json:"ad_account_creation_request"`
	AdAccountPromotableObjects        adaccountpromotableobjects.Entity `json:"ad_account_promotable_objects"`
	Age                               float64                           `json:"age"`
	AgencyClientDeclaration           agencyclientdeclaration.Entity    `json:"agency_client_declaration"`
	AmountSpent                       string                            `json:"amount_spent"`
	AttributionSpec                   []attributionspec.Entity          `json:"attribution_spec"`
	Balance                           string                            `json:"balance"`
	Business                          business.Entity                   `json:"business"`
	BusinessCity                      string                            `json:"business_city"`
	BusinessCountryCode               string                            `json:"business_country_code"`
	BusinessName                      string                            `json:"business_name"`
	BusinessState                     string                            `json:"business_state"`
	BusinessStreet                    string                            `json:"business_street"`
	BusinessStreet2                   string                            `json:"business_street2"`
	BusinessZip                       string                            `json:"business_zip"`
	CanCreateBrandLiftStudy           bool                              `json:"can_create_brand_lift_study"`
	Capabilities                      []string                          `json:"capabilities"`
	CreatedTime                       string                            `json:"created_time"`
	Currency                          string                            `json:"currency"`
	DisableReason                     int64                             `json:"disable_reason"`
	EndAdvertiser                     string                            `json:"end_advertiser"`
	EndAdvertiserName                 string                            `json:"end_advertiser_name"`
	ExtendedCreditInvoiceGroup        extendedcreditinvoicegroup.Entity `json:"extended_credit_invoice_group"`
	FailedDeliveryChecks              []deliverycheck.Entity            `json:"failed_delivery_checks"`
	FbEntity                          int64                             `json:"fb_entity"`
	FundingSource                     string                            `json:"funding_source"`
	FundingSourceDetails              fundingsourcedetails.Entity       `json:"funding_source_details"`
	HasMigratedPermissions            bool                              `json:"has_migrated_permissions"`
	HasPageAuthorizedAdaccount        bool                              `json:"has_page_authorized_adaccount"`
	ID                                string                            `json:"id"`
	IoNumber                          string                            `json:"io_number"`
	IsAttributionSpecSystemDefault    bool                              `json:"is_attribution_spec_system_default"`
	IsDirectDealsEnabled              bool                              `json:"is_direct_deals_enabled"`
	IsIn3dsAuthorizationEnabledMarket bool                              `json:"is_in_3ds_authorization_enabled_market"`
	IsInMiddleOfLocalEntityMigration  bool                              `json:"is_in_middle_of_local_entity_migration"`
	IsNotificationsEnabled            bool                              `json:"is_notifications_enabled"`
	IsPersonal                        int64                             `json:"is_personal"`
	IsPrepayAccount                   bool                              `json:"is_prepay_account"`
	IsTaxIdRequired                   bool                              `json:"is_tax_id_required"`
	LineNumbers                       []int64                           `json:"line_numbers"`
	MediaAgency                       string                            `json:"media_agency"`
	MinCampaignGroupSpendCap          string                            `json:"min_campaign_group_spend_cap"`
	MinDailyBudget                    int64                             `json:"min_daily_budget"`
	Name                              string                            `json:"name"`
	OffsitePixelsTosAccepted          bool                              `json:"offsite_pixels_tos_accepted"`
	Owner                             string                            `json:"owner"`
	Partner                           string                            `json:"partner"`
	RfSpec                            reachfrequencyspec.Entity         `json:"rf_spec"`
	ShowCheckoutExperience            bool                              `json:"show_checkout_experience"`
	SpendCap                          string                            `json:"spend_cap"`
	TaxID                             string                            `json:"tax_id"`
	TaxIDStatus                       int64                             `json:"tax_id_status"`
	TaxIDType                         string                            `json:"tax_id_type"`
	TimezoneID                        int64                             `json:"timezone_id"`
	TimezoneName                      string                            `json:"timezone_name"`
	TimezoneOffsetHoursUtc            float64                           `json:"timezone_offset_hours_utc"`
	TosAccepted                       map[string]int64                  `json:"tos_accepted"`
	UserTasks                         []string                          `json:"user_tasks"`
	UserTosAccepted                   map[string]int64                  `json:"user_tos_accepted"`
}

func NewAdAccount(id string, context *sdk.APIContext) *Entity {
	return &Entity{
		request: &sdk.APIRequest{Context: context},
		ID:      id,
	}
}

func FetchByID(id string, context *sdk.APIContext) (*Entity, error) {
	req := sdk.NewAPIRequest(
		context,
		prefixID(id),
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

func GetAdAccounts(context *sdk.APIContext) (*sdk.APINodeList, error) {
	req := sdk.NewAPIRequest(
		context,
		nodeID,
		endpoint,
		http.MethodGet,
		sdk.Parser(sdk.ParserResponse),
		sdk.ReturnFields(fields),
	)

	resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return resp.(*sdk.APINodeList), nil
}

func (ent *Entity) Fetch() (*Entity, error) {
	obj, err := FetchByID(ent.getPrefixID(), ent.GetRequest().Context)
	if err != nil {
		return nil, err
	}
	ent.copy(obj)
	return ent, nil
}

func (ent *Entity) GetCampaigns() (*sdk.APINodeList, error) {
	req := campaign.CreateAPIRequestGet(ent.getPrefixID(), ent.GetRequest().Context)
	resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return resp.(*sdk.APINodeList), nil
}

func (ent *Entity) GetInsights() *adsinsights.APIRequestGetInsights {
	return adsinsights.CreateAPIRequestGetInsights(
		ent.getPrefixID(),
		ent.GetRequest().Context,
	)
}

func ParseResponse(rawResp sdk.APIResponse) (resp []Entity, err error) {
	request := rawResp.GetRequest()
	context := rawResp.GetRequest().Context
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
		resp[i].ID = strings.Replace(resp[i].ID, "act_", "", 1)
		resp[i].SetRequest(request)
	}
	return
}

//
// For internal usage
//
func (ent *Entity) copy(other *Entity) {
	ent.ID = other.ID
	ent.AccountStatus = other.AccountStatus
	ent.AccountID = other.AccountID
}

func (ent *Entity) getPrefixID() string {
	return prefixID(ent.ID)
}

func prefixID(id string) string {
	return fmt.Sprintf(prefix, id)
}

func parserResponse(data json.RawMessage) (sdk.APIResponse, error) {
	ent := &Entity{}
	if err := json.Unmarshal(data, ent); err != nil {
		return ent, err
	}
	return ent, nil
}
