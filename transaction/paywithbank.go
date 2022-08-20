package transaction

import (
	"encoding/json"
	"fmt"

	"github.com/hisyntax/monnify-go"
)

type PayWithBankRes struct {
	RequestSuccessful bool               `json:"requestSuccessful"`
	ResponseMessage   string             `json:"responseMessage"`
	ResponseCode      string             `json:"responseCode"`
	ResponseBody      PayWithBankResBody `json:"responseBody"`
}

type PayWithBankResBody struct {
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
}

type PayWithBankReq struct {
	TransactionReference string `json:"transactionReference"`
	BankCode             string `json:"bankCode"`
}

func PayWithBank(payload PayWithBankReq) (*PayWithBankRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodPost
	isPayload := true
	url := fmt.Sprintf("%s/merchant/bank-transfer/init-payment", client.BaseUrl)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		fmt.Println(err)
	}
	var response PayWithBankRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}
