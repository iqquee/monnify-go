package limitprofile

import (
	"encoding/json"
	"fmt"

	"github.com/hisyntax/monnify-go"
)

type CreateLimitProfileReq struct {
	LimitProfileName       string `json:"limitProfileName"`
	SingleTransactionValue string `json:"singleTransactionValue"`
	DailyTransactionVolume string `json:"dailyTransactionVolume"`
	DailyTransactionValue  string `json:"dailyTransactionValue"`
	LimitProfileCode       string `json:"limitProfileCode"`
}

type CreateLimitProfileRes struct {
	RequestSuccessful bool                      `json:"requestSuccessful"`
	ResponseMessage   string                    `json:"responseMessage"`
	ResponseCode      string                    `json:"responseCode"`
	ResponseBody      CreateLimitProfileReqBody `json:"responseBody"`
}

type CreateLimitProfileReqBody struct {
	LimitProfileCode       string `json:"limitProfileCode"`
	LimitProfileName       string `json:"limitProfileName"`
	SingleTransactionValue string `json:"singleTransactionValue"`
	DailyTransactionVolume string `json:"dailyTransactionVolume"`
	DailyTransactionValue  string `json:"dailyTransactionValue"`
	DateCreated            string `json:"dateCreated"`
	LastModified           string `json:"lastModified"`
}

func CreateLimitProfile(payload CreateLimitProfileReq) (*CreateLimitProfileRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodPost
	isPayload := true
	url := fmt.Sprintf("%s/limit-profile", client.BaseUrl)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}

	var response CreateLimitProfileRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}
