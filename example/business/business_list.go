package main

import (
	. "github.com/mazti/facebook-go-business-sdk/example"
	"github.com/mazti/facebook-go-business-sdk/sdk"
	"github.com/mazti/facebook-go-business-sdk/sdk/business"
	"github.com/mazti/facebook-go-business-sdk/sdk/businessrolerequest"
	"github.com/mazti/facebook-go-business-sdk/sdk/businessuser"
)

func main() {
	context := sdk.NewContext(
		sdk.AppSecret(AppSecret),
		sdk.AccessToken(AccessToken),
		sdk.Logger(Log),
		sdk.Debug(true),
	)

	nodeList, err := business.GetBusinesses(context)

	if err != nil {
		context.Log(err)
		return
	}
	for nodeList != nil {
		businesses, err := business.ParseResponse(nodeList)
		context.Log(err)

		for _, biz := range businesses {
			nodeListUsers, err := biz.GetUsers()
			if err != nil {
				context.Log(err)
				continue
			}

			bizUsers, err := businessuser.ParseResponse(nodeListUsers)
			if err != nil {
				context.Log(err)
				continue
			}
			context.Log("business users:", len(bizUsers))

			createdUser, err := biz.CreateUser("xyxuxixuxyxu434xyxyuxuxuxyxuxuxux@gmail.com", businessuser.Employee)
			if err != nil {
				context.Log(err)
				continue
			}

			context.Log("created business role request id:", createdUser.Id)

			nodeListPendingUsers, err := biz.GetPendingUsers()
			if err != nil {
				context.Log(err)
				continue
			}

			pendingUsers, err := businessrolerequest.ParseResponse(nodeListPendingUsers)
			if err != nil {
				context.Log(err)
				continue
			}

			context.Log("pending users:", len(pendingUsers))
		}

		nodeList, err = nodeList.Next(0)
		context.Log("next err:", err)
	}
}
