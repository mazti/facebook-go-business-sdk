package adlabel

import (
	"encoding/json"
	"net/http"

	"github.com/mazti/facebook-go-business-sdk/sdk"
)

const (
	endpoint = "/"
)

type AdLabel struct {
	Context *sdk.APIContext

	ID   string `json:"id"`
	Name string `json:"name"`
	//Account     adaccount.AdAccount `json:"account"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
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
	ent := &AdLabel{}
	if err := json.Unmarshal(data, ent); err != nil {
		return ent, err
	}
	return ent, nil
}
