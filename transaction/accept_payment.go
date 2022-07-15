package transaction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/go-monnify"
)

type AcceptPaymentModel struct {
	PaymentReference    string `json:"paymentReference"`
	Amount              int    `json:"amount"`
	CurrencyCode        string `json:"currencyCode"`
	ContractCode        string `json:"contractCode"`
	CustomerEmail       string `json:"customerEmail"`
	CustomerName        string `json:"customerName"`
	CustomerPhoneNumber string `json:"customerPhoneNumber"`
	RedirectUrl         string `json:"redirectUrl"`
	PaymentDescription  string `json:"paymentDescription"`
}

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

type AcceptPaymentRes struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      AcceptPaymentResBody
}

type AcceptPaymentResBody struct {
	TransactionReference string
	PaymentReference     string
	MerchantName         string
	ApiKey               string
	EnabledPaymentMethod []string
	CheckoutUrl          string
}

func AcceptPayment(amount int, paymentReference, paymentDesc, currencyCode, contractCode, customerName, customerEmail, customerPhoneNumber, redirectUrl string) (*AcceptPaymentRes, int, error) {
	client := monnify.NewClient()
	url := fmt.Sprintf("%s/api/v1/merchant/transactions/init-transaction", client.Options.BaseUrl)
	method := "POST"
	payload := AcceptPaymentModel{
		PaymentReference:    paymentReference,
		Amount:              amount,
		CurrencyCode:        currencyCode,
		ContractCode:        contractCode,
		CustomerEmail:       customerEmail,
		CustomerName:        customerName,
		CustomerPhoneNumber: customerPhoneNumber,
		RedirectUrl:         redirectUrl,
		PaymentDescription:  paymentDesc,
	}
	// payload := AcceptPaymentModel{}
	// payload.PaymentReference = paymentReference
	// payload.Amount = amount
	// payload.CurrencyCode = currencyCode
	// payload.ContractCode = contractCode
	// payload.CustomerEmail = customerEmail
	// payload.CustomerName = customerName
	// payload.CustomerPhoneNumber = customerPhoneNumber
	// payload.RedirectUrl = redirectUrl
	// payload.PaymentDescription = paymentDesc

	jsonReq, jsonReqErr := json.Marshal(&payload)
	if jsonReqErr != nil {
		return nil, 0, jsonReqErr
	}

	req, reqErr := http.NewRequest(method, url, bytes.NewBuffer(jsonReq))
	if reqErr != nil {
		return nil, 0, reqErr
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", client.BasicToken)

	resp, respErr := client.Http.Do(req)
	if respErr != nil {
		return nil, 0, respErr
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	var response AcceptPaymentRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
