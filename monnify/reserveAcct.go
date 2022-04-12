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

type ReserveAccoutReq struct {
	AccountName       string              `json:"account_name"`
	AccountReference  string              `json:"account_reference"`
	CurrencyCode      string              `json:"currency_code"`
	ContractCode      string              `json:"contract_code"`
	CustomerName      string              `json:"customer_name"`
	CustomerEmail     string              `json:"customer_email"`
	IncomeSplitConfig []IncomeSplitConfig `json:"income_split_config"`
}

type IncomeSplitConfig struct {
	SubAccountCode  string `json:"Sub_account_code"`
	SplitPercentage int    `json:"split_percentage"`
	FeePercentage   int    `json:"fee_percentage"`
	FeeBearer       bool   `jsos:"fee_bearer"`
}

var Client = &http.Client{}

func ReserveAcct(base_url, bearer_auth string) (*ReserveAcctResponse, string, error) {
	client := Client
	url := fmt.Sprintf("%s/api/v1/bank-transfer/reserved-accounts", base_url)
	var reserveAccoutReq ReserveAccoutReq
	body, err := json.Marshal(&reserveAccoutReq)
	if err != nil {
		return nil, "", err
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bearer_auth))

	resp, err := client.Do(req)
	if err != nil {
		log.Println(helper.ServerErr)
		return nil, "", err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	log.Println(resp.Status)
	log.Println(string(resp_body))

	var reservedAcct ReserveAcctResponse
	json.Unmarshal(resp_body, &reservedAcct)

	return &reservedAcct, resp.Status, nil
}
