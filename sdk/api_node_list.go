package sdk

import "encoding/json"

type Cursors struct {
	Before string `json:"before,omitempty"`
	After  string `json:"after,omitempty"`
}

type Paging struct {
	Cursors Cursors `json:"cursors"`
	Next    string  `json:"next,omitempty"`
}

type APINodeList struct {
	Paging         Paging          `json:"paging"`
	Data           json.RawMessage `json:"data,omitempty"`
	request        *APIRequest
	body           []byte
	header         []byte
	autoPagination bool
	appSecret      string
}

func (n APINodeList) GetRawResponse() []byte {
	return n.body
}

func (n APINodeList) GetRawResponseAsJsonObject() json.RawMessage {
	return n.body
}

func (n APINodeList) GetHeader() []byte {
	return n.header
}
