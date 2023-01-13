package transaction

import (
	"encoding/json"
	"fmt"

	"github.com/iqquee/monnify-go"
)

type AcceptPaymentReq struct {
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
	RequestSuccessful bool                 `json:"requestSuccessful"`
	ResponseMessage   string               `json:"responseMessage"`
	ResponseCode      string               `json:"responseCode"`
	ResponseBody      AcceptPaymentResBody `json:"responseBody"`
}

type AcceptPaymentResBody struct {
	TransactionReference string   `json:"transactionReference"`
	PaymentReference     string   `json:"paymentReference"`
	MerchantName         string   `json:"merchantName"`
	ApiKey               string   `json:"apiKey"`
	EnabledPaymentMethod []string `json:"enabledPaymentMethod"`
	CheckoutUrl          string   `json:"checkoutUrl"`
}

type getTransacStatusRes struct {
	RequestSuccessful bool                    `json:"requestSuccessful"`
	ResponseMessage   string                  `json:"responseMessage"`
	ResponseCode      string                  `json:"responseCode"`
	ResponseBody      getTransacStatusResBody `json:"responseBody"`
}
type getTransacStatusResBody struct {
	CreatedOn            string  `json:"createdOn"`
	Amount               float64 `json:"amount"`
	CurrencyCode         string  `json:"currencyCode"`
	CustomerName         string  `json:"customerName"`
	CustomerEmail        string  `json:"customerEmail"`
	PaymentDescription   string  `json:"paymentDescription"`
	PaymentStatus        string  `json:"paymentStatus"`
	TransactionReference string  `json:"transactionReference"`
	PaymentReference     string  `json:"paymentReference"`
}

func AcceptPayment(payload AcceptPaymentReq) (*AcceptPaymentRes, int, error) {
	client := monnify.NewClient()
	isPayload := true
	method := monnify.MethodPost
	url := fmt.Sprintf("%s/merchant/transactions/init-transaction", client.BaseUrl)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}

	var response AcceptPaymentRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

func GetTransactionStatus(paymentRef string) (*getTransacStatusRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodGet
	isPayload := false
	payload := ""
	url := fmt.Sprintf("%s/merchant/transactions/query?paymentReference=%s", client.BaseUrl, paymentRef)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}
	var response getTransacStatusRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}
