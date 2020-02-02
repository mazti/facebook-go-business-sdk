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

	account := adaccount.NewAdAccount("2745055142177332", context)

	account, err := account.Fetch()

	context.Log(err)
	context.Log(account.AccountID, account.AccountStatus)
}
