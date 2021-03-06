package sdk

import (
	"crypto/hmac"
	csha256 "crypto/sha256"
	"encoding/hex"
)

type APIContext struct {
	endpointBase      string
	videoEndpointBase string
	accessToken       string
	appSecret         string
	appID             string
	version           string
	userAgent         string
	isDebug           bool
	logger            func(...interface{})
}

func (ctx *APIContext) Log(a ...interface{}) {
	if ctx.isDebug && ctx.logger != nil {
		ctx.logger(a...)
	}
}

func (ctx *APIContext) EnableDebug(debug bool) *APIContext {
	ctx.isDebug = debug
	return ctx
}

func (ctx *APIContext) GetAppSecretProof() string {
	return sha256(ctx.appSecret, ctx.accessToken)
}

// NewContext create a context with options for constructor
func NewContext(options ...func(*APIContext)) *APIContext {
	ctx := &APIContext{
		endpointBase:      DefaultAPIBase,
		videoEndpointBase: DefaultVideoBase,
		version:           DefaultAPIVersion,
		userAgent:         DefaultUserAgent,
	}

	for _, option := range options {
		option(ctx)
	}
	return ctx
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

func Logger(logger func(...interface{})) func(*APIContext) {
	return func(ctx *APIContext) {
		ctx.logger = logger
	}
}

func Debug(isDebug bool) func(*APIContext) {
	return func(ctx *APIContext) {
		ctx.isDebug = isDebug
	}
}

func sha256(secret, message string) string {
	mac := hmac.New(csha256.New, []byte(secret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}
