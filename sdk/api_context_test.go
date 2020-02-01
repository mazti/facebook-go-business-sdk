package sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	accessToken = "access_token"
	appSecret   = "app_secret"
	appID       = "app_id"
)

func TestNewContextSimple(t *testing.T) {
	ctx := NewContext()

	assert.NotNil(t, ctx)
	assert.Equal(t, ctx.endpointBase, DefaultAPIBase)
	assert.Equal(t, ctx.version, DefaultAPIVersion)
	assert.Equal(t, ctx.videoEndpointBase, DefaultVideoBase)
}

func TestNewContextWithToken(t *testing.T) {
	ctx := NewContext(
		AccessToken(accessToken),
		AppSecret(appSecret),
	)

	assert.NotNil(t, ctx)
	assert.Equal(t, accessToken, ctx.accessToken)
	assert.Equal(t, appSecret, ctx.appSecret)
}

func TestNewContextWithAppID(t *testing.T) {
	ctx := NewContext(
		AppID(appID),
	)

	assert.NotNil(t, ctx)
	assert.Empty(t, ctx.accessToken)
	assert.Empty(t, ctx.appSecret)
	assert.Equal(t, appID, ctx.appID)
}

func TestNewContextWithLog(t *testing.T) {
	mocker := loggerMock{}
	mocker.On("Log", mock.Anything).Once()

	ctx := NewContext(
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
	result := sha256(appSecret, accessToken)
	assert.Equal(t, expected, result)
}

func TestSHA256_EmptyAccessToken(t *testing.T) {
	result := sha256(appSecret, "")
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
