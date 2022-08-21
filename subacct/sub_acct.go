package subacct

import (
	"encoding/json"
	"fmt"

	"github.com/hisyntax/monnify-go"
)

type CreateSubAccountReq struct {
	CurrencyCode           string `json:"currencyCode"`
	AccountNumber          string `json:"accountNumber"`
	AccountName            string `json:"accountName"`
	SubAccountCode         string `json:"subAccountCode"`
	BankCode               string `json:"bankCode"`
	Email                  string `json:"email"`
	DefaultSplitPercentage string `json:"defaultSplitPercentage"`
}

type CreateSubAccountRes struct {
	RequestSuccessful bool                    `json:"requestSuccessful"`
	ResponseMessage   string                  `json:"responseMessage"`
	ResponseCode      string                  `json:"responseCode"`
	ResponseBody      CreateSubAccountResBody `json:"responseBody"`
}

type CreateSubAccountResBody struct {
	SubAccountCode         string `json:"subAccountCode"`
	AccountNumber          string `json:"accountNumber"`
	AccountName            string `json:"accountName"`
	CurrencyCode           string `json:"currencyCode"`
	Email                  string `json:"email"`
	BankCode               string `json:"bankCode"`
	BankName               string `json:"bankName"`
	DefaultSplitPercentage string `json:"defaultSplitPercentage"`
}

func CreateSubAccount(payload CreateSubAccountReq) (*CreateSubAccountRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodPost
	isPayload := true
	url := fmt.Sprintf("%s/sub-accounts", client.BaseUrl)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		fmt.Println(err)
	}

	var response CreateSubAccountRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}
