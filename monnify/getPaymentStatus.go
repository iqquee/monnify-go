package monnify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AcctPaymentStatus struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      AcctPaymentStatusBody
}

type AcctPaymentStatusBody struct {
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

func GetPaymentStatus(bearerToken, base_url, paymentRef string) (*AcctPaymentStatus, string, error) {
	bearer_token := fmt.Sprintf("Bearer %s", bearerToken)

	client := Client

	url := fmt.Sprintf("%s/api/v1/merchant/transactions/query?paymentReference=%s", base_url, paymentRef)
	req, _ := http.NewRequest("GET", url, nil)

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

	var acctPaymentStatus AcctPaymentStatus
	json.Unmarshal(resp_body, &acctPaymentStatus)
	return &acctPaymentStatus, resp.Status, nil
}
