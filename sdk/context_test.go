package sdk

import (
	"github.com/mazti/facebook-go-business-sdk/sdk/config"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	accessToken = "access_token"
	appSecret   = "app_secret"
	appID       = "app_id"
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
		AccessToken(accessToken),
		AppSecret(appSecret),
	)

	assert.NotNil(t, ctx)
	assert.Equal(t, accessToken, ctx.accessToken)
	assert.Equal(t, appSecret, ctx.appSecret)
}

func TestNewContext3(t *testing.T) {
	ctx := New(
		AppID(appID),
	)

	assert.NotNil(t, ctx)
	assert.Empty(t, ctx.accessToken)
	assert.Empty(t, ctx.appSecret)
	assert.Equal(t, appID, ctx.appID)
}

func TestNewContext4(t *testing.T) {
	mocker := loggerMock{}
	mocker.On("Log", mock.Anything).Once()

	ctx := New(
		Logger(mocker.Log),
		Debug(true),
	)

	assert.NotNil(t, ctx)
	assert.Empty(t, ctx.accessToken)
	assert.Empty(t, ctx.appSecret)

	ctx.Log("hello", "tien", "!")

	mocker.AssertCalled(t, "Log", mock.Anything)
}

func TestSHA256_OK(t *testing.T) {
	expected := "d52ddf968d622d8af8677906b7fbae09ac1f89f7cd5c1584b27544624cc23e5a"
	result := SHA256(appSecret, accessToken)
	assert.Equal(t, expected, result)
}

func TestSHA256_EmptyAccessToken(t *testing.T) {
	result := SHA256(appSecret, "")
	assert.NotNil(t, result)
}

type loggerMock struct {
	mock.Mock
}

func (_m *loggerMock) Log(_a0 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0...)
	_m.Called(_ca...)
}
