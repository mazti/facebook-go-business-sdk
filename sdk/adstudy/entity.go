package adstudy

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
	// TODO: adding more field here
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
