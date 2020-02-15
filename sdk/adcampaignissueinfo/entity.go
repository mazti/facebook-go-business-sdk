package adcampaignissueinfo

import (
	"encoding/json"

	"github.com/mazti/facebook-go-business-sdk/sdk"
)

type AdCampaignIssueInfo struct {
	Context *sdk.APIContext

	ErrorCode    int64 `json:"error_code"`
	ErrorMessage int64 `json:"error_message"`
	ErrorSummary int64 `json:"error_summary"`
	ErrorType    int64 `json:"error_type"`
	Level        int64 `json:"level"`
}

func parserResponse(data json.RawMessage) (sdk.APIResponse, error) {
	ent := &AdCampaignIssueInfo{}
	if err := json.Unmarshal(data, ent); err != nil {
		return ent, err
	}
	return ent, nil
}
