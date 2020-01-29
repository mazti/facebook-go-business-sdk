package sdk

import (
	"encoding/json"
)

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
	RawData        json.RawMessage `json:"data,omitempty"`
	Data           interface{}
	request        *APIRequest
	body           []byte
	header         []byte
	autoPagination bool
	appSecret      string
}

func (n APINodeList) Parse(data json.RawMessage) error {
	if err := json.Unmarshal(data, &n); err != nil {
		return err
	}
	if err := json.Unmarshal(n.RawData, n.Data); err != nil {
		return err
	}
	return nil
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
