package monnify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/go-monnify/helper"
)

type AcctTransferStatus struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      AcctTransferStatusBody
}

type AcctTransferStatusBody struct {
	CreatedOn            string
	Amount               int
	CurrencyCode         string
	CustomerName         string
	CustomerEmail        string
	PaymentDescription   string
	PaymentStatus        string
	TransactionReference string
	PaymentReference     string
}

func GetAcctTransferStatus(bearerToken, base_url, transactRef string) (*AcctTransferStatus, string, error) {
	bearer_token := fmt.Sprintf("Bearer %s", bearerToken)

	client := Client

	url := fmt.Sprintf("%s/api/v1/merchant/transactions/query?transactionReference=%s", base_url, transactRef)
	req, _ := http.NewRequest("GET", url, nil)

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

	var acctTransferStatus AcctTransferStatus
	json.Unmarshal(resp_body, &acctTransferStatus)
	return &acctTransferStatus, resp.Status, nil
}
