package sdk

import "encoding/json"

type APINode struct {
	context  APIContext
	rawValue string
	header   string
}

func LoadJSON(ctx APIContext, json string, header string) APINode {
	return APINode{
		context:  ctx,
		rawValue: json,
		header:   header,
	}
}

func (n APINode) GetRawResponse() string {
	return ""
}

func (n APINode) GetRawResponseAsJsonObject() json.RawMessage {
	return json.RawMessage{}
}

func (n APINode) GetHeader() string {
	return ""
}
