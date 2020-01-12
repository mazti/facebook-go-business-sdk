package sdk

import (
	"github.com/mazti/facebook-go-business-sdk/sdk/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContext1(t *testing.T) {
	ctx := New()

	assert.NotNil(t, ctx)
	assert.Equal(t, ctx.endpointBase, config.DefaultAPIBase)
	assert.Equal(t, ctx.version, config.DefaultAPIVersion)
	assert.Equal(t, ctx.videoEndpointBase, config.DefaultVideoBase)
}

func TestNewContext2(t *testing.T) {
	ctx := New(
		AccessToken("this_is_secret_token"),
		AppSecret("this_is_app_secret"),
	)

	assert.NotNil(t, ctx)
	assert.Equal(t, "this_is_secret_token", ctx.accessToken)
	assert.Equal(t, "3.0", ctx.version)
}
