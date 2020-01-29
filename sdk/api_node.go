package sdk

import "encoding/json"

type APINode struct {
	context  *APIContext
	rawValue []byte
	header   []byte
}

func LoadJSON(ctx *APIContext, body []byte, header []byte) APINode {
	return APINode{
		context:  ctx,
		rawValue: body,
		header:   header,
	}
}

func (n APINode) GetRawResponse() []byte {
	return n.rawValue
}

func (n APINode) GetRawResponseAsJsonObject() json.RawMessage {
	return n.rawValue
}

func (n APINode) GetHeader() []byte {
	return n.header
}

func (n APINode) Parse(data json.RawMessage) (APIResponse, error) {
	return nil, nil
}
