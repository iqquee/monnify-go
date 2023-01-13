package account

import (
	"encoding/json"
	"fmt"

	"github.com/iqquee/monnify-go"
)

type CreateReservedAcctReq struct {
	AccountName       string                   `json:"accountName"`
	AccountReference  string                   `json:"accountReference"`
	CurrencyCode      string                   `json:"currencyCode"`
	ContractCode      string                   `json:"contractCode"`
	CustomerName      string                   `json:"customerName"`
	CustomerEmail     string                   `json:"customerEmail"`
	IncomeSplitConfig IncomeSplitConfigReqBody `json:"incomeSplitConfig"`
}

type IncomeSplitConfigReqBody struct {
	SubAccountCode  string `json:"subAccountCode"`
	SplitPercentage int    `json:"splitPercentage"`
	FeePercentage   int    `json:"feePercentage"`
	FeeBearer       bool   `json:"feeBearer"`
}

type CreateReservedAcctRes struct {
	RequestSuccessful bool                      `json:"requestSuccessful"`
	ResponseMessage   string                    `json:"responseMessage"`
	ResponseCode      string                    `json:"responseCode"`
	ResponseBody      CreateReservedAcctResBody `json:"responseBody"`
}

type CreateReservedAcctResBody struct {
	ContractCode      string      `json:"contractCode"`
	AccountReference  string      `json:"accountReference"`
	AccountName       string      `json:"accountName"`
	CurrencyCode      string      `json:"currencyCode"`
	CustomerEmail     string      `json:"customerEmail"`
	CustomerName      string      `json:"customerName"`
	AccountNumber     string      `json:"accountNumber"`
	BankName          string      `json:"bankName"`
	BankCode          string      `json:"bankCode"`
	Status            string      `json:"status"`
	CreatedOn         string      `json:"createdOn"`
	IncomeSplitConfig interface{} `json:"incomeSplitConfig"`
}

func CreateReservedAcct(payload CreateReservedAcctReq) (*CreateReservedAcctRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodPost
	isPayload := true
	url := fmt.Sprintf("%s/bank-transfer/reserved-accounts", client.BaseUrl)
	token := fmt.Sprintf("Bearer %s", client.BearerToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}
	var response CreateReservedAcctRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil

}
