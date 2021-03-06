package main

import (
	. "github.com/mazti/facebook-go-business-sdk/example"
	"github.com/mazti/facebook-go-business-sdk/sdk"
	"github.com/mazti/facebook-go-business-sdk/sdk/adaccount"
)

func main() {
	context := sdk.NewContext(
		sdk.AppSecret(AppSecret),
		sdk.AccessToken(AccessToken),
		sdk.Logger(Log),
		sdk.Debug(true),
	)

	req := sdk.NewAPIRequest(
		context,
		"me",
		"adaccounts",
		"GET",
	)
	resp, err := req.Execute()
	if err != nil {
		context.Log(err)
		return
	}

	context.Log("-----")
	adaccounts, err := adaccount.ParseResponse(resp)
	if err != nil {
		context.Log(err)
		return
	}
	context.Log(adaccounts)
}
