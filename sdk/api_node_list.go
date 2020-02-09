package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
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

// TODO: improve ParserResponse
func ParserResponse(data json.RawMessage) (APIResponse, error) {
	nodeList := createAPINodeList()
	if err := json.Unmarshal(data, nodeList); err != nil {
		return nil, err
	}
	return nodeList, nil
}

func (ent *APINodeList) Next(limit int) (*APINodeList, error) {
	// First check if 'next' url is returned. If so, always use the it.
	if len(ent.Paging.Next) > 0 {
		// App secret won't return with the 'next' URL. Need to append it  for paging.
		nextURL := ent.getNextURL()
		ent.request.SetOverrideURL(nextURL)
		resp, err := ent.request.Execute()
		if err != nil {
			return nil, err
		}
		return resp.(*APINodeList), nil
	}
	after := ent.Paging.Cursors.After
	if len(after) < 1 {
		return nil, errors.New("after empty")
	}
	ent.request.SetOverrideURL("")
	extraParams := map[string]interface{}{
		afterKey: after,
	}
	if limit > 0 {
		extraParams[limitKey] = limit
	}
	resp, err := ent.request.ExecuteWithParams(extraParams)
	if err != nil {
		return nil, err
	}
	return resp.(*APINodeList), nil
}

func (ent APINodeList) Unmarshal(v interface{}) error {
	// TODO: Take care about node context for each value
	if err := json.Unmarshal(ent.Data, v); err != nil {
		return err
	}
	return nil
}

//
// Internal functions
//
func createAPINodeList() *APINodeList {
	return &APINodeList{
		node: &APINode{},
	}
}

func (ent APINodeList) getNextURL() string {
	if len(ent.appSecret) < 1 {
		return ent.Paging.Next
	}
	nextURL, err := url.Parse(ent.Paging.Next)
	if err != nil {
		return ""
	}
	connector := "?"
	if len(nextURL.Query()) > 0 {
		connector = "&"
	}
	return fmt.Sprintf("%s%s%s=%s",
		ent.Paging.Next, connector, appSecretProofKey, ent.appSecret)
}

//
// Functions for APINode
//

func (ent *APINodeList) Load(context *APIContext, req *APIRequest, header []byte, body []byte) {
	ent.node.Load(context, req, header, body)
}

func (ent *APINodeList) SetContext(context *APIContext) {
	ent.node.SetContext(context)
}
