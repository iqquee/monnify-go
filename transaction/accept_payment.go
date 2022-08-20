package transaction

import (
	"encoding/json"
	"fmt"

	"github.com/hisyntax/monnify-go"
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

func AcceptPayment(payload AcceptPaymentReq) (*AcceptPaymentRes, int, error) {
	client := monnify.NewClient()
	isPayload := true
	method := monnify.MethodPost
	url := fmt.Sprintf("%s/merchant/transactions/init-transaction", client.BaseUrl)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
	}
	var response getTransacStatusRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}
