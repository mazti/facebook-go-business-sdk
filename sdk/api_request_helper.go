package sdk

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"net/url"
	"path"
	"reflect"
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
			queryValues := make(url.Values)
			mapParams := rawParams.(map[string]interface{})
			for k, v := range mapParams {
				queryValues.Add(k, convertToString(v))
			}
			requestURL.RawQuery = queryValues.Encode()
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

func convertToString(obj interface{}) string {
	if obj == nil {
		return ""
	}
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.String {
		return obj.(string)
	}
	if bytes, ok := obj.([]byte); ok {
		return string(bytes)
	}

	bytes, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(bytes)
}
