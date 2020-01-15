package sdk

import "encoding/json"

type APIResponse interface {
	GetRawResponse() string
	GetRawResponseAsJsonObject() json.RawMessage
	GetHeader() string
}
