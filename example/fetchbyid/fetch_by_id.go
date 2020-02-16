package main

import (
	. "github.com/mazti/facebook-go-business-sdk/example"
	"github.com/mazti/facebook-go-business-sdk/sdk"
	"github.com/mazti/facebook-go-business-sdk/sdk/adaccount"
	"github.com/mazti/facebook-go-business-sdk/sdk/adsinsights"
)

func main() {
	context := sdk.NewContext(
		sdk.AppID(AppID),
		sdk.AppSecret(AppSecret),
		sdk.AccessToken(AccessToken),
		sdk.Logger(Log),
		sdk.Debug(true),
	)

	account, err := adaccount.FetchByID(AccountID, context)

	context.Log(err)
	context.Log(account)

	nodeList, err := account.GetInsights()
	context.Log("get insights err:", err)
	context.Log(err)

	for nodeList != nil {
		insights, err := adsinsights.ParseResponse(nodeList)
		context.Log(err)
		context.Log("insights:", insights)

		nodeList, err = nodeList.Next(0)
		context.Log("next err:", err)
	}
}
