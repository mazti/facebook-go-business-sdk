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

	nodeList, err := adaccount.GetAdAccounts(context)

	if err != nil {
		context.Log(err)
		return
	}
	for nodeList != nil {
		adaccounts, err := adaccount.ParseResponse(nodeList)
		context.Log(err)
		context.Log("adaccounts:", adaccounts)

		nodeList, err = nodeList.Next(0)
		context.Log("next err:", err)
	}
}
