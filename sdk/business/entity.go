package business

import (
	"encoding/json"
	"net/http"

	"github.com/mazti/facebook-go-business-sdk/sdk"
	"github.com/mazti/facebook-go-business-sdk/sdk/page"
)

const (
	endpoint = "/"
)

type Entity struct {
	request *sdk.APIRequest

	BlockOfflineAnalytics           bool        `json:"block_offline_analytics"`
	CreatedBy                       string      `json:"created_by"`
	CreatedTime                     string      `json:"created_time"`
	ExtendedUpdatedTime             string      `json:"extended_updated_time"`
	ID                              string      `json:"id"`
	IsHidden                        bool        `json:"is_hidden"`
	IsInstagramEnabledInFbAnalytics bool        `json:"is_instagram_enabled_in_fb_analytics"`
	Link                            string      `json:"link"`
	Name                            string      `json:"name"`
	PaymentAccountID                string      `json:"payment_account_id"`
	PrimaryPage                     page.Entity `json:"primary_page"`
	ProfilePictureUri               string      `json:"profile_picture_uri"`
	TimezoneID                      int64       `json:"timezone_id"`
	TwoFactorType                   string      `json:"two_factor_type"`
	UpdatedBy                       string      `json:"updated_by"`
	UpdatedTime                     string      `json:"updated_time"`
	VerificationStatus              string      `json:"verification_status"`
	Vertical                        string      `json:"vertical"`
	VerticalID                      int64       `json:"vertical_id"`
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

func parserResponse(data json.RawMessage) (sdk.APIResponse, error) {
	ent := &Entity{}
	if err := json.Unmarshal(data, ent); err != nil {
		return ent, err
	}
	return ent, nil
}
