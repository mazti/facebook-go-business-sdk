package sdk

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
	overrideURL      string
	lastResponse     APIResponse
}
