package sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRequestURL1(t *testing.T) {
	param := map[string]interface{}{
		"hehe":   1,
		"hihi":   "tien",
		"khakha": true,
	}

	url, err := createRequestURL("/test", param)

	assert.Nil(t, err)
	assert.Equal(t, url, "/test?hehe=1&hihi=tien&khakha=true")
}
