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

	nodeList := resp.(*sdk.APINodeList)

	context.Log("-----")
	var adaccounts []adaccount.AdAccount

	err = nodeList.Unmarshal(&adaccounts)
	if err != nil {
		context.Log(err)
		return
	}
	context.Log(adaccounts)

	context.Log("-----")
	insights, err := adaccounts[0].GetInsights()
	if err != nil {
		context.Log(err)
		return
	}
	context.Log(insights)
}
