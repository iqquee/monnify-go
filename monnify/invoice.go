package monnify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/go-monnify/helper"
)

type Invoice struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      InvoiceBody
}

type InvoiceBody struct {
	Amount           int
	InvoiceReference string
	InvoiceStatus    string
	Description      string
	ContractCode     string
	CustomerEmail    string
	CustomerName     string
	ExpiryDate       string
	CreatedBy        string
	CreatedOn        string
	CheckoutUrl      string
	AccountNumber    string
	AccountName      string
	BankName         string
	BankCode         string
}

func CreateInvoice(bearerToken, base_url, amount, invoiceReference, description, contractCode, customerEmail, customerName, expiryDate, redirectUrl string) (*Invoice, string, error) {
	bearer_token := fmt.Sprintf("Bearer %s", bearerToken)

	client := Client
	url := fmt.Sprintf("%s/api/v1/invoice/create", base_url)
	body := []byte(fmt.Sprintf("{\n  \"amount\": \"%s\",\n  \"invoiceReference\": \"%s\",\n  \"description\": \"%s\",\n  \"currencyCode\": \"NGN\",\n  \"contractCode\": \"%s\",\n  \"customerEmail\": \"%s\",\n  \"customerName\": \"%s\",\n  \"expiryDate\": \"%s\",\n  \"paymentMethods\": [\n    \"ACCOUNT_TRANSFER\", \"CARD\"  ],\n  \"redirectUrl\": \"%s\",\n }", amount, invoiceReference, description, contractCode, customerEmail, customerName, expiryDate, redirectUrl))

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

	var invoice Invoice
	json.Unmarshal(resp_body, &invoice)
	return &invoice, resp.Status, nil
}

func GetInvoice(bearerToken, base_url, invoiceRef string) (*Invoice, string, error) {
	bearer_token := fmt.Sprintf("Bearer %s", bearerToken)

	client := Client
	url := fmt.Sprintf("%s/api/v1/invoice/details?invoiceReference=%s", base_url, invoiceRef)
	req, _ := http.NewRequest("GET", url, nil)

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

	var invoice Invoice
	json.Unmarshal(resp_body, &invoice)
	return &invoice, resp.Status, nil
}

type Invoices struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      []InvoicesBody
}

type InvoicesBody struct {
	Amount           int
	InvoiceReference string
	InvoiceStatus    string
	Description      string
	ContractCode     string
	CustomerEmail    string
	CustomerName     string
	ExpiryDate       string
	CreatedBy        string
	CreatedOn        string
	CheckoutUrl      string
	AccountNumber    string
	AccountName      string
	BankName         string
	BankCode         string
}

func GetAllInvoices(bearerToken, base_url string) (*Invoices, string, error) {
	bearer_token := fmt.Sprintf("Bearer %s", bearerToken)

	client := Client
	url := fmt.Sprintf("%s/api/v1/invoice/all", base_url)
	req, _ := http.NewRequest("GET", url, nil)

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

	var invoices Invoices
	json.Unmarshal(resp_body, &invoices)
	return &invoices, resp.Status, nil
}

func CancelInvoice(bearerToken, base_url, invoiceRef string) (*Invoice, string, error) {
	bearer_token := fmt.Sprintf("Bearer %s", bearerToken)

	client := Client
	url := fmt.Sprintf("%s/api/v1/invoice/cancel?invoiceReference=%s", base_url, invoiceRef)
	req, _ := http.NewRequest("DELETE", url, nil)

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

	var invoice Invoice
	json.Unmarshal(resp_body, &invoice)
	return &invoice, resp.Status, nil
}

type InvoiceAcct struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      InvoiceAcctBody
}
type InvoiceAcctBody struct {
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
	ReservedAccountType  string
	Status               string
	CreatedOn            string
}

func CreateInvoiceReservedAcct(bearerToken, base_url, contractCode, accountName, accountReference, customerEmail, customerName string) (*InvoiceAcct, string, error) {
	bearer_token := fmt.Sprintf("Bearer %s", bearerToken)

	client := Client
	url := fmt.Sprintf("%s/api/v1/bank-transfer/reserved-accounts", base_url)

	body := []byte(fmt.Sprintf("{\n  \"contractCode\": \"%s\",\n  \"accountName\": \"%s\",\n  \"currencyCode\": \"NGN\",\n  \"accountReference\": \"%s\",\n  \"customerEmail\": \"%s\",\n  \"customerName\": \"%s\",\n  \"reservedAccountType\": \"INVOICE\"\n}", contractCode, accountName, accountReference, customerEmail, customerName))

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

	var invoice InvoiceAcct
	json.Unmarshal(resp_body, &invoice)
	return &invoice, resp.Status, nil
}

// type ReserveInvoiceAccountModel struct {
// 	RequestSuccessful bool
// 	ResponseMessage   string
// 	ResponseCode      string
// 	ResponseBody      ReserveInvoiceAccountBody
// }

// type ReserveInvoiceAccountBody struct {
// 	Amount           int
// 	InvoiceReference string
// 	InvoiceStatus    string
// 	Description      string
// 	ContractCode     string
// 	CustomerEmail    string
// 	CustomerName     string
// 	ExpiryDate       string
// 	CreatedBy        string
// 	CreatedOn        string
// 	CheckoutUrl      string
// 	AccountNumber    string
// 	AccountName      string
// 	BankName         string
// 	BankCode         string
// }

func ReserveInvoiceAccount(bearerToken, base_url, amount, invoiceReference, accountReference, description, contractCode, customerEmail, customerName, expiryDate string) (*Invoice, string, error) {
	bearer_token := fmt.Sprintf("Bearer %s", bearerToken)

	client := Client
	url := fmt.Sprintf("%s/api/v1/invoice/create", base_url)
	body := []byte(fmt.Sprintf("{\n  \"amount\": \"%s\",\n  \"invoiceReference\": \"%s\",\n  \"accountReference\": \"%s\",\n  \"description\": \"%s\",\n  \"currencyCode\": \"NGN\",\n  \"contractCode\": \"%s\",\n  \"customerEmail\": \"%s\",\n  \"customerName\": \"%s\",\n  \"expiryDate\": \"%s\"\n}", amount, invoiceReference, accountReference, description, contractCode, customerEmail, customerName, expiryDate))
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

	var invoices Invoice
	json.Unmarshal(resp_body, &invoices)
	return &invoices, resp.Status, nil
}
