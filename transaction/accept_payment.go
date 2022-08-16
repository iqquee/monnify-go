package transaction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/monnify-go"
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

type getTransacStatusRes struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      getTransacStatusResBody
}
type getTransacStatusResBody struct {
	CreatedOn            string
	Amount               float64
	CurrencyCode         string
	CustomerName         string
	CustomerEmail        string
	PaymentDescription   string
	PaymentStatus        string
	TransactionReference string
	PaymentReference     string
}

func AcceptPayment(amount int, paymentReference, paymentDesc, currencyCode, contractCode, customerName, customerEmail, customerPhoneNumber, redirectUrl string) (*AcceptPaymentRes, int, error) {
	client := monnify.NewClient()
	url := fmt.Sprintf("%s/merchant/transactions/init-transaction", client.BaseUrl)
	method := "POST"
	token := fmt.Sprintf("Basic %s", client.BasicToken)
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

	jsonReq, jsonReqErr := json.Marshal(&payload)
	if jsonReqErr != nil {
		return nil, 0, jsonReqErr
	}

	req, reqErr := http.NewRequest(method, url, bytes.NewBuffer(jsonReq))
	if reqErr != nil {
		return nil, 0, reqErr
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

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

func GetTransactionStatus(paymentRef string) (*getTransacStatusRes, int, error) {
	client := monnify.NewClient()
	url := fmt.Sprintf("%s/merchant/transactions/query?paymentReference=%s", client.BaseUrl, paymentRef)
	method := "GET"
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	req, reqErr := http.NewRequest(method, url, nil)
	if reqErr != nil {
		return nil, 0, reqErr
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	resp, err := client.Http.Do(req)
	if err != nil {
		return nil, 0, err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	var response getTransacStatusRes

	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
