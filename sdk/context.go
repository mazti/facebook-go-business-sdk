package sdk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

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

func AppID(id string) func(*APIContext) {
	return func(ctx *APIContext) {
		ctx.appID = id
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

func SHA256(secret, message string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}
