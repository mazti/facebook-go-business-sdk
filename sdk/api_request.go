package sdk

import (
	"fmt"
	"strings"
)

const (
	accessTokenKey = "access_token"
	appSecretProof = "appsecret_proof"
	fields         = "fields"
)

type APIRequest struct {
	context          *APIContext
	executor         IRequestExecutor
	userAgent        string
	useVideoEndpoint bool
	nodeID           string
	endpoint         string
	method           string
	paramNames       []string
	params           map[string]interface{}
	returnFields     []string
	overrideURL      string
	lastResponse     APIResponse
}

type ResponseWrapper struct {
	body   []byte
	header []byte
}

func NewResponseWrapper(body []byte, header []byte) *ResponseWrapper {
	return &ResponseWrapper{
		body:   body,
		header: header,
	}
}

func (req *APIRequest) Execute() (APIResponse, error) {
	return req.ExecuteWithParams(nil)
}

func (req *APIRequest) ExecuteWithParams(extraParams map[string]interface{}) (APIResponse, error) {
	rw := req.executeInternal(extraParams)
	req.lastResponse = req.parseResponse(rw.body, rw.header)
	return req.lastResponse, nil
}

func (req *APIRequest) executeInternal(extraParams map[string]interface{}) *ResponseWrapper {
	ctx := req.context
	ctx.Log("========Start of API Call========")
	defer ctx.Log("========End of API Call========")
	resp, err := req.executor.Execute(req.method, req.getURL(), req.getAllParams(extraParams), req.context)
	if err != nil {
		return resp
	}
	ctx.Log("Response", resp.body)
	return resp
}

func (req *APIRequest) parseResponse(body []byte, header []byte) APIResponse {
	return nil
}

func (req *APIRequest) getURL() string {
	if len(req.overrideURL) > 0 {
		return req.overrideURL
	}
	ctx := req.context
	endpointBase := ctx.endpointBase
	if req.useVideoEndpoint {
		endpointBase = ctx.videoEndpointBase
	}
	return fmt.Sprintf("%s/%s/%s/%s", endpointBase, ctx.version, req.nodeID, req.endpoint)
}

func (req *APIRequest) getAllParams(extraParams map[string]interface{}) (allParams map[string]interface{}) {
	allParams = map[string]interface{}{}
	ctx := req.context
	if len(req.overrideURL) > 0 {
		return allParams
	}
	for k, v := range req.params {
		allParams[k] = v
	}
	for k, v := range extraParams {
		allParams[k] = v
	}
	allParams[accessTokenKey] = ctx.accessToken
	if len(ctx.appSecret) > 0 {
		allParams[appSecretProof] = ctx.GetAppSecretProof()
	}
	if len(req.returnFields) > 0 {
		allParams[fields] = strings.Join(req.returnFields, ",")
	}
	return allParams
}
