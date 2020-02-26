package businesscreateuser

import (
	"encoding/json"
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

const (
	endpoint = "/business_users"
)

type Entity struct {
	request *sdk.APIRequest

	Id string `json:"id"`
}

func (ent *Entity) CreateBusinessUser(id string, context *sdk.APIContext) *APIRequestCreateBusinessUser {
	return CreateAPIRequestCreateBusinessUser(
		id,
		context,
	)
}

func ParserResponse(data json.RawMessage) (sdk.APIResponse, error) {
	entity := &Entity{}
	if err := json.Unmarshal(data, entity); err != nil {
		return nil, err
	}
	return entity, nil
}