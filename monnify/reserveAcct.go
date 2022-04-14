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

func ReserveAcct(base_url, bearerToken, acctName, acctRef, contractcode, customerName, customerEmail string) (*ReserveAcctResponse, string, error) {

	bearer_token := fmt.Sprintf("Bearer %s", bearerToken)
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

type ReserveAcctLimitModel struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      ReserveAcctLimitBody
}

type ReserveAcctLimitBody struct {
	ContractCode         string
	AccountReference     string
	AccountName          string
	CurrencyCode         string
	CustomerEmail        string
	CustomerName         string
	AccountNumber        string
	BankName             string
	BankCode             string
	ReservationReference string
	Status               string
	CreatedOn            string
	LimitProfileConfig   ReserveAcctLimitBodyConfig
}

type ReserveAcctLimitBodyConfig struct {
	LimitProfileCode       string
	SingleTransactionValue int
	DailyTransactionVolume int
	DailyTransactionValue  int
}

func ReserveAcctLimit(bearerToken, base_url, contractCode, accountName, accountReference, customerEmail, customerName, limitProfileCode string) (*ReserveAcctLimitModel, string, error) {
	bearer_token := fmt.Sprintf("Bearer %s", bearerToken)
	client := Client

	body := []byte(fmt.Sprintf("{\n  \"contractCode\": \"%s\",\n  \"accountName\": \"%s \",\n  \"currencyCode\": \"NGN\",\n  \"accountReference\": \"%s\",\n  \"customerEmail\": \"%s\",\n  \"customerName\": \"%s\",\n  \"limitProfileCode\": \"%s\"\n}", contractCode, accountName, accountReference, customerEmail, customerName, limitProfileCode))
	url := fmt.Sprintf("%s/api/v1/bank-transfer/reserved-accounts/limit", base_url)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer_token)

	resp, err := client.Do(req)

	if err != nil {
		log.Println(helper.ServerErr)
		return nil, "", err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

	var reserveAcctLimitModel ReserveAcctLimitModel
	json.Unmarshal(resp_body, &reserveAcctLimitModel)
	return &reserveAcctLimitModel, resp.Status, nil
}

func UpdateReserveAcctLimit(bearerToken, base_url, accountReference, limitProfileCode string) (*ReserveAcctLimitModel, string, error) {
	bearer_token := fmt.Sprintf("Bearer %s", bearerToken)
	client := Client

	body := []byte(fmt.Sprintf("{\n  \"accountReference\": \"%s\",\n  \"limitProfileCode\": \"%s\"\n}", accountReference, limitProfileCode))

	url := fmt.Sprintf("%s/api/v1/bank-transfer/reserved-accounts/limit", base_url)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(body))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer_token)

	resp, err := client.Do(req)

	if err != nil {
		log.Println(helper.ServerErr)
		return nil, "", err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

	var reserveAcctLimitModel ReserveAcctLimitModel
	json.Unmarshal(resp_body, &reserveAcctLimitModel)
	return &reserveAcctLimitModel, resp.Status, nil
}
