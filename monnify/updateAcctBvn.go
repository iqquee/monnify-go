package monnify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AcctResponse struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      AcctResponseBody
}
type AcctResponseBody struct {
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

func UpdateAcctBVN(base_url, bearerToken, acctRef, bvn string) (*AcctResponse, string, error) {
	bearer_token := fmt.Sprintf("Bearer %s", bearerToken)

	client := Client
	url := fmt.Sprintf("%s/api/v1/bank-transfer/reserved-accounts/update-customer-bvn/%s", base_url, acctRef)
	body := []byte(fmt.Sprintf("{\n  \"bvn\": \"%s\"}", bvn))
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(body))

	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", bearer_token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error occured when sending request to the server")
		log.Println(err)
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	var acctRes AcctResponse
	json.Unmarshal(resp_body, &acctRes)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

	return &acctRes, resp.Status, nil
}
