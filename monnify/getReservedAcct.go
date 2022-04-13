package monnify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ReservedAcct struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      ReservedAcctBody
}

type ReservedAcctBody struct {
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
	Contract             ContractBody
	TotalAmount          int
	TransactionCount     int
}
type ContractBody struct {
	Name                                       string
	Code                                       string
	Description                                string
	SupportsAdvancedSettlementAccountSelection bool
}

func GetReservedAccts(bearerToken, base_url, acctRef string) (*ReservedAcct, string, error) {
	bearer_token := fmt.Sprintf("Bearer %s", bearerToken)

	client := Client
	url := fmt.Sprintf("%s/api/v1/bank-transfer/reserved-accounts/%s", base_url, acctRef)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer_token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return nil, "", err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

	var reservedAcct ReservedAcct
	json.Unmarshal(resp_body, &reservedAcct)

	return &reservedAcct, resp.Status, nil
}
