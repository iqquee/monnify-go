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

func Payment(amount, chargeFee int, base_url, cName, cEmail, cPhone, paymentRef, paymentDesc, redirectUrl, contractCode, bearerToken string) (*InitTransacStatus, string, error) {
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

type DisbursementModel struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      DisbursementModelBody
}

type DisbursementModelBody struct {
	Amount      int
	Reference   string
	Status      string
	DateCreated string
}

func SingleDisbursement(basicToken, base_url, reference, narration, bankCode, accountNumber, walletId string, amount int) (*DisbursementModel, string, error) {
	basic_token := fmt.Sprintf("Basic %s", basicToken)

	client := Client

	body := []byte(fmt.Sprintf("{\n  \"amount\": %d,\n  \"reference\": \"%s\",\n  \"narration\": \"%s\",\n  \"bankCode\": \"%s\",\n  \"accountNumber\": \"%s\",\n  \"currency\": \"NGN\",\n  \"walletId\": \"%s\"\n}", amount, reference, narration, bankCode, accountNumber, walletId))
	url := fmt.Sprintf("%s/api/v1/disbursements/single", base_url)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", basic_token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(helper.ServerErr)
		return nil, "", err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

	var disbursementModel DisbursementModel
	json.Unmarshal(resp_body, &disbursementModel)
	return &disbursementModel, resp.Status, nil
}

func SingleAuthDisbursement(basicToken, base_url, reference, authorizationCode string) (*DisbursementModel, string, error) {
	basic_token := fmt.Sprintf("Basic %s", basicToken)

	client := Client

	url := fmt.Sprintf("%s/api/v1/disbursements/single/validate-otp", base_url)

	body := []byte(fmt.Sprintf("{\n  \"reference\": \"%s\",\n  \"authorizationCode\": \"%s\"\n}", reference, authorizationCode))

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", basic_token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(helper.ServerErr)
		return nil, "", err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

	var disbursementModel DisbursementModel
	json.Unmarshal(resp_body, &disbursementModel)
	return &disbursementModel, resp.Status, nil
}

// func BulkDisburseMent(basicToken, base_url, title string) {
// 	basic_token := fmt.Sprintf("Basic %s", basicToken)

// 	client := Client
// 	body := []byte(fmt.Sprintf("{\n  \"title\": \"Game of Batches\",\n  \"batchReference\": \"batchreference12934\",\n  \"narration\": \"911 Transaction\",\n  \"walletId\": \"4794983C91374AD6B3ECD76F2BEA296D\",\n  \"onValidationFailure\": \"CONTINUE\",\n  \"notificationInterval\": 10,\n  \"transactionList\": [\n    {\n      \"amount\": 1300,\n      \"reference\": \"Final-Reference-1a\",\n      \"narration\": \"911 Transaction\",\n      \"bankCode\": \"058\",\n      \"accountNumber\": \"0111946768\",\n      \"currency\": \"NGN\"\n    },\n    {\n      \"amount\": 570,\n      \"reference\": \"Final-Reference-2a\",\n      \"narration\": \"911 Transaction\",\n      \"bankCode\": \"058\",\n      \"accountNumber\": \"0111946768\",\n      \"currency\": \"NGN\"\n    },\n    {\n      \"amount\": 230,\n      \"reference\": \"Final-Reference-3a\",\n      \"narration\": \"911 Transaction\",\n      \"bankCode\": \"058\",\n      \"accountNumber\": \"0111946768\",\n      \"currency\": \"NGN\"\n    }\n  ]\n}"))
// 	url := fmt.Sprintf("%s/api/v1/disbursements/batch", base_url)
// 	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))

// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Authorization", basic_token)

// 	resp, err := client.Do(req)

// 	if err != nil {
// 		fmt.Println("Errored when sending request to the server")
// 		return
// 	}

// 	defer resp.Body.Close()
// 	resp_body, _ := ioutil.ReadAll(resp.Body)

// 	fmt.Println(resp.Status)
// 	fmt.Println(string(resp_body))
// }
