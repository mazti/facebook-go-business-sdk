package sdk

import "fmt"

type APIRequest struct {
	userAgent        string
	context          *APIContext
	useVideoEndpoint bool
	nodeID           string
	endpoint         string
	method           string
	paramNames       []string
	params           map[string]interface{}
	returnFields     []string
	overrideUrl      string
	lastResponse     APIResponse
}

func (req APIRequest) executeInternal(extraParams map[string]interface{}) {
}

func (req APIRequest) getApiUrl() string {
	if len(req.overrideUrl) > 0 {
		return req.overrideUrl
	}
	ctx := req.context
	endPointBase := ctx.endpointBase
	if req.useVideoEndpoint {
		endPointBase = ctx.videoEndpointBase
	}
	return fmt.Sprintf("%s/%s/%s/%s", endPointBase, ctx.version, req.nodeID, req.endpoint)
}
