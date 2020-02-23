package main

import (
	. "github.com/mazti/facebook-go-business-sdk/example"
	"github.com/mazti/facebook-go-business-sdk/sdk"
	"github.com/mazti/facebook-go-business-sdk/sdk/adaccount"
	"github.com/mazti/facebook-go-business-sdk/sdk/adaccountuser"
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
		if err != nil {
			context.Log(err)
			continue
		}

		for _, acc := range adaccounts {
			nodeListUsers, err := acc.GetUsers()
			if err != nil {
				context.Log(err)
				continue
			}

			accUsers, err := adaccountuser.ParseResponse(nodeListUsers)
			if err != nil {
				context.Log(err)
				continue
			}

			context.Log("adaccount users:", accUsers)

		}

		nodeList, err = nodeList.Next(0)
		context.Log("next err:", err)
	}
}
