package monnify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/go-monnify/helper"
)

type PayWithBank struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      PayWithBankBody
}

type PayWithBankBody struct {
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
	ProductInformation     string
}

func PayWithBankTransfer(bearerToken, base_url, acctRef, bankCode string) (*PayWithBank, string, error) {
	bearer_token := fmt.Sprintf("Bearer %s", bearerToken)
	client := Client

	url := fmt.Sprintf("%s/api/v1/merchant/bank-transfer/init-payment", base_url)
	body := []byte(fmt.Sprintf("{\n  \"transactionReference\": \"%s\",\n  \"bankCode\": \"%s\"\n}", acctRef, bankCode))

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer_token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(helper.ServerErr)
		return nil, "", err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

	var payWithBank PayWithBank
	json.Unmarshal(resp_body, &payWithBank)
	return &payWithBank, resp.Status, nil
}
