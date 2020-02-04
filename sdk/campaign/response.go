package campaign

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

func (ent Campaign) GetBody() []byte {
	return ent.node.GetBody()
}

func (ent Campaign) GetHeader() []byte {
	return ent.node.GetBody()
}

func (ent *Campaign) SetRequest(request *sdk.APIRequest) {
	ent.node.SetRequest(request)
}

func (ent *Campaign) SetBody(body []byte) {
	ent.node.SetBody(body)
}

func (ent *Campaign) SetHeader(header []byte) {
	ent.node.SetHeader(header)
}
