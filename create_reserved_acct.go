package monnify

import (
	"fmt"
	"net/http"
)

type (
	// CreateReservedAcctReq request object
	CreateReservedAcctReq struct {
		AccountName       string                   `json:"accountName"`
		AccountReference  string                   `json:"accountReference"`
		CurrencyCode      string                   `json:"currencyCode"`
		ContractCode      string                   `json:"contractCode"`
		CustomerName      string                   `json:"customerName"`
		CustomerEmail     string                   `json:"customerEmail"`
		IncomeSplitConfig IncomeSplitConfigReqBody `json:"incomeSplitConfig"`
	}

	// IncomeSplitConfigReqBody request object
	IncomeSplitConfigReqBody struct {
		SubAccountCode  string `json:"subAccountCode"`
		SplitPercentage int    `json:"splitPercentage"`
		FeePercentage   int    `json:"feePercentage"`
		FeeBearer       bool   `json:"feeBearer"`
	}

	// CreateReservedAcctRes response object
	CreateReservedAcctRes struct {
		RequestSuccessful bool   `json:"requestSuccessful"`
		ResponseMessage   string `json:"responseMessage"`
		ResponseCode      string `json:"responseCode"`
		ResponseBody      struct {
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
		} `json:"responseBody"`
	}
)

// CreateReservedAcct allows the creation of dedicated virtual accounts for your customers.
func (c *Client) CreateReservedAcct(payload CreateReservedAcctReq) (*CreateReservedAcctRes, error) {
	url := fmt.Sprintf("%s/bank-transfer/reserved-accounts", c.baseURL)

	c.isBasic = false
	var response CreateReservedAcctRes
	if err := c.newRequest(http.MethodPost, url, payload, response); err != nil {
		if err = c.generateNewBearerToken(); err != nil {
			return nil, err
		}
	}

	return &response, nil
}
