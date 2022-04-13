package monnify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/go-monnify/helper"
)

type InitTransacStatus struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      InitTransacStatusBody
}
type InitTransacStatusBody struct {
	TransactionReference string
	PaymentReference     string
	AmountPaid           string
	TotalPayable         string
	SettlementAmount     string
	PaidOn               string
	PaymentStatus        string
	PaymentDescription   string
	Currency             string
	PaymentMethod        string
}

func InitTransaction(amount, chargeFee int, base_url, cName, cEmail, cPhone, paymentRef, paymentDesc, redirectUrl, contractCode, bearerToken string) (*InitTransacStatus, string, error) {
	bearer_token := fmt.Sprintf("Bearer %s", bearerToken)

	client := Client

	url := fmt.Sprintf("%s/api/v1/merchant/transactions/init-transaction", base_url)
	body := []byte(fmt.Sprintf("{\n  \"amount\": %d,\n  \"customerName\": \"%s\",\n  \"customerEmail\": \"%s\",\n  \"customerPhoneNumber\": \"%s\",\n  \"paymentReference\": \"%s\",\n  \"paymentDescription\": \"%s\",\n  \"currencyCode\": \"NGN\",\n  \"contractCode\": \"%s\",\n  \"redirectUrl\": \"%s/\"\n}", amount, cName, cEmail, cPhone, paymentRef, paymentDesc, contractCode, redirectUrl))
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

	var initTransacStatus InitTransacStatus
	json.Unmarshal(resp_body, &initTransacStatus)

	return &initTransacStatus, resp.Status, nil
}
