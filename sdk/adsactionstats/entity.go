package adsactionstats

import (
	"encoding/json"
	"net/http"

	"github.com/mazti/facebook-go-business-sdk/sdk"
)

const (
	endpoint = "/"
)

type Entity struct {
	request *sdk.APIRequest

	D1Click                             string `json:"1d_click"`
	D1View                              string `json:"1d_view"`
	D28Click                            string `json:"28d_click"`
	D28View                             string `json:"28d_view"`
	D7Click                             string `json:"7d_click"`
	D7View                              string `json:"7d_view"`
	ActionCanvasComponentId             string `json:"action_canvas_component_id"`
	ActionCanvasComponentName           string `json:"action_canvas_component_name"`
	ActionCarouselCardId                string `json:"action_carousel_card_id"`
	ActionCarouselCardName              string `json:"action_carousel_card_name"`
	ActionConvertedProductId            string `json:"action_converted_product_id"`
	ActionDestination                   string `json:"action_destination"`
	ActionDevice                        string `json:"action_device"`
	ActionEventChannel                  string `json:"action_event_channel"`
	ActionLinkClickDestination          string `json:"action_link_click_destination"`
	ActionLocationCode                  string `json:"action_location_code"`
	ActionReaction                      string `json:"action_reaction"`
	ActionTargetId                      string `json:"action_target_id"`
	ActionType                          string `json:"action_type"`
	ActionVideoAssetId                  string `json:"action_video_asset_id"`
	ActionVideoSound                    string `json:"action_video_sound"`
	ActionVideoType                     string `json:"action_video_type"`
	Dda                                 string `json:"dda"`
	Inline                              string `json:"inline"`
	InteractiveComponentStickerId       string `json:"interactive_component_sticker_id"`
	InteractiveComponentStickerResponse string `json:"interactive_component_sticker_response"`
	Value                               string `json:"value"`
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
