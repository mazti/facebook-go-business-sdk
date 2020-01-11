package sdk

type APIContext struct {
	endpointBase      string
	videoEndpointBase string
	accessToken       string
	appSecret         string
	appID             string
	version           string
	isDebug           bool
	// log used for logging on debug mode.
	log func(...interface{})
}

func New(endpointBase, videoEndpointBase, accessToken,
	appSecret, appID, version string) *APIContext {
	return &APIContext{
		endpointBase:      endpointBase,
		videoEndpointBase: videoEndpointBase,
		accessToken:       accessToken,
		appSecret:         appSecret,
		appID:             appID,
		version:           version,
		isDebug:           false,
	}
}
