package sdk

import "encoding/json"

type APIResponse interface {
	GetRawResponse() []byte
	GetRawResponseAsJsonObject() json.RawMessage
	GetHeader() []byte
	Parse(data json.RawMessage) (APIResponse, error)
}
