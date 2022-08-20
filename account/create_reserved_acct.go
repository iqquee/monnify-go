package account

import (
	"encoding/json"
	"fmt"

	"github.com/hisyntax/monnify-go"
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
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      CreateReservedAcctResBody
}

type CreateReservedAcctResBody struct {
	ContractCode      string
	AccountReference  string
	AccountName       string
	CurrencyCode      string
	CustomerEmail     string
	CustomerName      string
	AccountNumber     string
	BankName          string
	BankCode          string
	Status            string
	CreatedOn         string
	IncomeSplitConfig interface{}
}

func CreateReservedAcct(payload CreateReservedAcctReq) (*CreateReservedAcctRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodPost
	isPayload := true
	url := fmt.Sprintf("%s/bank-transfer/reserved-accounts", client.BaseUrl)
	token := fmt.Sprintf("Bearer %s", client.BearerToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		fmt.Println(err)
	}
	var response CreateReservedAcctRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil

}
