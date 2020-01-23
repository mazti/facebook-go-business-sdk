package main

import (
	"fmt"
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

var (
	accessToken = "EAAI2nqB1r4MBAJoRlfhMCr1zrDN81PUtt64H9LN5r2I35Q9xOEjLyQSIDWZAqrxJ6U5hpcZAIBFZAZCkpZAeL7C7gczhRZBEl8oe9DSxOjt73rQ4GJHBm8VbuaTvhnRdztl1pp8RzdY8lSJcg3WvEXhP4jIGGCDeASvcHGK5OBCAZDZD"
	appSecret   = "277fc82bedf9841cef91a15acef02d00"
)

func Log(a ...interface{}) {
	fmt.Println(a...)
}

func main() {
	context := sdk.NewContext(
		sdk.AppSecret(appSecret),
		sdk.AccessToken(accessToken),
		sdk.Logger(Log),
		sdk.Debug(true),
	)
	req := sdk.NewAPIRequest(context, "me", "adaccounts", "GET")
	resp, err := req.Execute()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(resp.GetRawResponse()))
}
