package sdkv2

import (
	"strings"
)

type APIRequest struct {
	nodeID      string
	method      string
	endpoint    string
	path        []string
	fields      []string
	params      map[string]interface{}
	fileParams  map[string]interface{}
	fileCounter int
}

func NewAPIRequest(nodeID, method, endpoint string) *APIRequest {
	endpoint = strings.Replace(endpoint, "/", "", 1)
	req := &APIRequest{
		nodeID:      nodeID,
		method:      method,
		endpoint:    endpoint,
		path:        []string{nodeID, endpoint},
		fields:      []string{},
		params:      map[string]interface{}{},
		fileParams:  nil,
		fileCounter: 0,
	}
	return req
}

func (api *APIRequest) AddField(field string) *APIRequest {
	if !contains(api.fields, field) {
		api.fields = append(api.fields, field)
	}
	return api
}

func (api *APIRequest) AddParam(key string, value interface{}) *APIRequest {
	api.params[key] = value
	return api
}

func (api *APIRequest) AddParams(params map[string]interface{}) *APIRequest {
	api.params = params
	return api
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}

	return false
}
