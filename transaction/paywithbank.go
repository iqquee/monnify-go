package transaction

import (
	"encoding/json"
	"fmt"

	"github.com/hisyntax/monnify-go"
)

type PayWithBankRes struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      PayWithBankResBody
}

type PayWithBankResBody struct {
	AccountNumber          string
	AccountName            string
	BankName               string
	BankCode               string
	AccountDurationSeconds int
	UssdPayment            string
	RequestTime            string
	TransactionReference   string
	PaymentReference       string
	Amount                 int
	Fee                    int
	TotalPayable           int
	CollectionChannel      string
	ProductInformation     interface{}
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
