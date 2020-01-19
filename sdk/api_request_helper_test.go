package sdk

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRequestURL1(t *testing.T) {
	kiet := []interface{}{
		"huhu",
		"hehe",
	}

	param := map[string]interface{}{
		"hehe":   1,
		"hihi":   []byte("tien"),
		"khakha": true,
		"kiet":   kiet,
	}

	expected := "/test?hihi=tien&hehe=1&khakha=true&kiet=%5B%22huhu%22%2C%22hehe%22%5D"
	result, err := createRequestURL("/test", param)

	expectedURL, _ := url.Parse(expected)
	resultURL, _ := url.Parse(result)
	expectedQuery, _ := url.ParseQuery(expectedURL.RawQuery)
	resultQuery, _ := url.ParseQuery(resultURL.RawQuery)

	fmt.Println(result)

	assert.Nil(t, err)
	assert.Equal(t, expectedQuery, resultQuery)
}
