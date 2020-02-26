package businesscreateuser

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
	"github.com/mazti/facebook-go-business-sdk/sdk/businessuser"
	"net/http"
)

type APIRequestCreateBusinessUser struct {
	request *sdk.APIRequest
}

func CreateAPIRequestCreateBusinessUser(nodeID string, context *sdk.APIContext) *APIRequestCreateBusinessUser {
	return &APIRequestCreateBusinessUser{
		request: sdk.NewAPIRequest(
			context,
			nodeID,
			endpoint,
			http.MethodPost,
			sdk.Parser(ParserResponse),
			sdk.ParamNames(params),
		),
	}
}

func (r *APIRequestCreateBusinessUser) SetEmail(email string) *APIRequestCreateBusinessUser {
	r.request.SetParam(sdk.EmailKey, email)
	return r
}

func (r *APIRequestCreateBusinessUser) SetRole(role businessuser.Role) *APIRequestCreateBusinessUser {
	r.request.SetParam(sdk.RoleKey, string(role))
	return r
}

func (r *APIRequestCreateBusinessUser) Execute() (*Entity, error) {
	resp, err := r.request.Execute()
	if err != nil {
		return nil, err
	}
	return resp.(*Entity), nil
}