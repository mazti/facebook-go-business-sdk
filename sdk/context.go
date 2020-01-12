package sdk

import (
	"github.com/mazti/facebook-go-business-sdk/sdk/config"
)

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

// accessToken, appSecret, appID, version string

func AccessToken(token string) func(*APIContext) {
	return func(ctx *APIContext) {
		ctx.accessToken = token
	}
}

func AppSecret(secret string) func(*APIContext) {
	return func(ctx *APIContext) {
		ctx.appSecret = secret
	}
}

func New(options ...func(*APIContext)) *APIContext {
	ctx := &APIContext{
		endpointBase:      config.DefaultAPIBase,
		videoEndpointBase: config.DefaultVideoBase,
		version:           config.DefaultAPIVersion,
	}

	for _, option := range options {
		option(ctx)
	}
	return ctx
}
