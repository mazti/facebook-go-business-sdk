package adsactionstats

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

func (ent *Entity) SetRequest(request *sdk.APIRequest) {
	ent.request = request
}

func (ent *Entity) GetRequest() *sdk.APIRequest {
	return ent.request
}
