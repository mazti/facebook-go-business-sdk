package main

import (
	"fmt"
	"github.com/mazti/facebook-go-business-sdk/sdk"
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
	_, err := req.Execute()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
