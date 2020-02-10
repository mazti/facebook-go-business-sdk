package sdk

import (
	"errors"
	"fmt"
)

var (
	UnsupportedResponse = errors.New("unsupported response")
	ReachToTheEnd       = errors.New("reach to the end")
)

var (
	UnsupportedMethod = func(method string) error {
		return fmt.Errorf("unsupported method: %s", method)
	}
)
