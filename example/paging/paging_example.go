package main

import (
	. "github.com/mazti/facebook-go-business-sdk/example"
	"github.com/mazti/facebook-go-business-sdk/sdk"
	"github.com/mazti/facebook-go-business-sdk/sdk/adaccount"
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
	context.Log(err)
	context.Log(account.AccountID, account.AccountStatus)

	campaigns, err := account.GetCampaigns()

	for campaigns != nil {
		context.Log("---- Campaign ----")

		context.Log(err)
		context.Log(string(campaigns.Data))

		campaigns, err = campaigns.Next(0)
	}

}
