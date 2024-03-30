package monnify

import (
	"fmt"
	"net/http"
)

type (

	// PayWithBankReq request object
	PayWithBankReq struct {
		TransactionReference string `json:"transactionReference"`
		BankCode             string `json:"bankCode"`
	}

	// PayWithBankRes response body
	PayWithBankRes struct {
		RequestSuccessful bool   `json:"requestSuccessful"`
		ResponseMessage   string `json:"responseMessage"`
		ResponseCode      string `json:"responseCode"`
		ResponseBody      struct {
			AccountNumber          string      `json:"accountNumber"`
			AccountName            string      `json:"accountName"`
			BankName               string      `json:"bankName"`
			BankCode               string      `json:"bankCode"`
			AccountDurationSeconds int         `json:"accountDurationSeconds"`
			UssdPayment            string      `json:"ussdPayment"`
			RequestTime            string      `json:"requestTime"`
			TransactionReference   string      `json:"transactionReference"`
			PaymentReference       string      `json:"paymentReference"`
			Amount                 int         `json:"amount"`
			Fee                    int         `json:"fee"`
			TotalPayable           int         `json:"totalPayable"`
			CollectionChannel      string      `json:"collectionChannel"`
			ProductInformation     interface{} `json:"productInformation"`
		} `json:"responseBody"`
	}
)

// PayWithBank generates a dynamic account number and its associated bank for one time payment
func (c *Client) PayWithBank(payload PayWithBankReq) (*PayWithBankRes, error) {
	url := fmt.Sprintf("%s/v1/merchant/bank-transfer/init-payment", c.baseURL)

	c.isBasic = false
	var response PayWithBankRes
	if err := c.newRequest(http.MethodPost, url, payload, response); err != nil {
		if err = c.generateNewBearerToken(); err != nil {
			return nil, err
		}
	}

	return &response, nil
}
