package transaction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

func PayWithBank(transactionReference, bankCode string) (*PayWithBankRes, int, error) {
	client := monnify.NewClient()
	url := fmt.Sprintf("%s/merchant/bank-transfer/init-payment", client.BaseUrl)
	method := "POST"
	token := fmt.Sprintf("Basic %s", client.BasicToken)
	payload := PayWithBankReq{
		TransactionReference: transactionReference,
		BankCode:             bankCode,
	}

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
	var response PayWithBankRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
