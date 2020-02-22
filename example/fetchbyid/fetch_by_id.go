package main

import (
	"time"

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

	getInsightsRequest := account.GetInsights()

	since, _ := time.Parse(sdk.TimeFormat, "2020-02-13")
	until, _ := time.Parse(sdk.TimeFormat, "2020-02-13")

	nodeList, err := getInsightsRequest.
		SetTimeRange(since, until).
		Execute()

	for nodeList != nil {
		insights, err := adsinsights.ParseResponse(nodeList)
		context.Log(err)
		context.Log("insights:", insights)

		nodeList, err = nodeList.Next(0)
		context.Log("next err:", err)
	}
}
