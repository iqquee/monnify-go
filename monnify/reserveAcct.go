package monnify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hisyntax/go-monnify/helper"
)

type ReserveAcctResponse struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      ReserveAcctResponseBody
}
type ReserveAcctResponseBody struct {
	ContractCode          string
	AccountReference      string
	AccountName           string
	CurrencyCode          string
	CustomerEmail         string
	CustomerName          string
	AccountNumber         string
	BankName              string
	BankCode              string
	CollectionChannel     string
	ReservationReference  string
	ReservedAccountType   string
	Status                string
	CreatedOn             string
	Bvn                   string
	IncomeSplitConfig     []string
	RestrictPaymentSource bool
}

var Client = &http.Client{}

func ReserveAcct(base_url, BearerToken, acctName, acctRef, contractcode, customerName, customerEmail string) (*ReserveAcctResponse, string, error) {

	bearer_token := fmt.Sprintf("Bearer %s", BearerToken)
	client := Client
	url := fmt.Sprintf("%s/api/v1/bank-transfer/reserved-accounts", base_url)
	body := []byte(fmt.Sprintf("{\n  \"accountName\": \"%s\",\n  \"accountReference\": \"%s\",\n  \"currencyCode\": \"NGN\",\n  \"contractCode\": \"%s\",\n  \"customerName\": \"%s\",\n  \"customerEmail\": \"%s\"\n}", acctName, acctRef, contractcode, customerName, customerEmail))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer_token)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(helper.ServerErr)
		return nil, "", err
	}

	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	log.Println(resp.Status)
	log.Println(string(resp_body))

	var reservedAcct ReserveAcctResponse
	json.Unmarshal(resp_body, &reservedAcct)

	log.Println(reservedAcct.ResponseBody.AccountName)

	return &reservedAcct, resp.Status, nil
}
