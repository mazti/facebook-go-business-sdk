package sdk

import (
	"fmt"
	"net/url"
	"path"

	"github.com/google/go-querystring/query"
)

func createRequestURL(rawURL string, rawParams interface{}) (string, error) {
	// Parse request url
	requestURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// Normalize url path
	requestURL.Path = path.Clean(requestURL.Path)

	// Set optional params to url query
	if rawParams != nil {
		switch params := rawParams.(type) {
		case string:
			requestURL.RawQuery = params
		case []byte:
			requestURL.RawQuery = string(params)
		case map[string]interface{}:
			fmt.Println("cai lon gi day")
		default:
			queryValues, err := query.Values(params)
			if err != nil {
				return "", err
			}
			requestURL.RawQuery = queryValues.Encode()
		}
	}

	// Return request url as string
	return requestURL.String(), nil
}
