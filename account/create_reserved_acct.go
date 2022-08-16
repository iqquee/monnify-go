package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/monnify-go"
)

type CreateReservedAcctReq struct {
	AccountName       string `json:"accountName"`
	AccountReference  string `json:"accountReference"`
	CurrencyCode      string `json:"currencyCode"`
	ContractCode      string `json:"contractCode"`
	CustomerName      string `json:"customerName"`
	CustomerEmail     string `json:"customerEmail"`
	IncomeSplitConfig IncomeSplitConfigReqBody
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
	url := fmt.Sprintf("%s/bank-transfer/reserved-accounts", client.BaseUrl)
	method := "POST"
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	jsonReq, jsonReqErr := json.Marshal(&payload)
	if jsonReqErr != nil {
		return nil, 0, jsonReqErr
	}

	req, reqErr := http.NewRequest(method, url, bytes.NewBuffer(jsonReq))
	if reqErr != nil {
		return nil, 0, reqErr
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	resp, respErr := client.Http.Do(req)
	if respErr != nil {
		return nil, 0, respErr
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	var response CreateReservedAcctRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil

}
