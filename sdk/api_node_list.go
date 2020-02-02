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
	node           *APINode
	Paging         Paging          `json:"paging"`
	Data           json.RawMessage `json:"data,omitempty"`
	request        *APIRequest
	autoPagination bool
	appSecret      string
}

func createAPINodeList() *APINodeList {
	return &APINodeList{
		node: &APINode{},
	}
}

func ParserResponse(data json.RawMessage) (APIResponse, error) {
	nodeList := createAPINodeList()
	if err := json.Unmarshal(data, nodeList); err != nil {
		return nil, err
	}
	return nodeList, nil
}

func (ent APINodeList) GetBody() []byte {
	return ent.node.GetBody()
}

func (ent APINodeList) GetHeader() []byte {
	return ent.node.GetHeader()
}

func (ent *APINodeList) SetBody(body []byte) {
	ent.node.SetBody(body)
}

func (ent *APINodeList) SetHeader(header []byte) {
	ent.node.SetHeader(header)
}
