package main

import (
	. "github.com/mazti/facebook-go-business-sdk/example"
	"github.com/mazti/facebook-go-business-sdk/sdk"
	"github.com/mazti/facebook-go-business-sdk/sdk/adaccount"
	"github.com/mazti/facebook-go-business-sdk/sdk/campaign"
)

func main() {
	context := sdk.NewContext(
		sdk.AppID(AppID),
		sdk.AppSecret(AppSecret),
		sdk.AccessToken(AccessToken),
		sdk.Logger(Log),
		sdk.Debug(true),
	)

	account := adaccount.NewAdAccount(AccountID, context)

	account, err := account.Fetch()
	context.Log("fetch err:", err)
	context.Log(account.AccountID, account.AccountStatus)

	nodeList, err := account.GetCampaigns()
	context.Log("get campaigns err:", err)
	for nodeList != nil {
		context.Log("---- Campaign ----")

		campaigns, err := campaign.ParseResponse(nodeList)
		context.Log(err)
		context.Log("campaigns:", campaigns)

		nodeList, err = nodeList.Next(0)
		context.Log("next err:", err)
	}
}
