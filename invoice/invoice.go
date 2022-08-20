package invoice

import (
	"encoding/json"
	"fmt"

	"github.com/hisyntax/monnify-go"
)

type CreateInvoiceReq struct {
	Amount            int                      `json:"amount"`
	InvoiceReference  string                   `json:"invoiceReference"`
	Description       string                   `json:"description"`
	CurrencyCode      string                   `json:"currencyCode"`
	ContractCode      string                   `json:"contractCode"`
	CustomerEmail     string                   `josn:"customerEmail"`
	CustomerName      string                   `json:"customerName"`
	ExpiryDate        string                   `json:"expiryDate"`
	RedirectUrl       string                   `josn:"redirectUrl"`
	PaymentMethod     string                   `json:"paymentMethod"`
	IncomeSplitConfig IncomeSplitConfigReqBody `json:"incomeSplitConfig"`
}

type IncomeSplitConfigReqBody struct {
	SubAccountCode  string `json:"subAccountCode"`
	SplitPercentage int    `json:"splitPercentage"`
	FeePercentage   int    `json:"feePercentage"`
	FeeBearer       bool   `json:"feeBearer"`
}

type CreateInvoiceRes struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      CreateInvoiceResBody
}

type CreateInvoiceResBody struct {
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

func CreateInvoice(payload CreateInvoiceReq) (*CreateInvoiceRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodPost
	isPayload := false
	url := fmt.Sprintf("%s/invoice/create", client.BaseUrl)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		fmt.Println(err)
	}
	var response CreateInvoiceRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

func GetInvoiceDetailsRequest(invoiceRef string) (*CreateInvoiceRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodGet
	isPayload := false
	payload := ""
	url := fmt.Sprintf("%s/invoice//details?invoiceReference=%s", client.BaseUrl, invoiceRef)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		fmt.Println(err)
	}

	var response CreateInvoiceRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

type GetAllInvoiceDetailsRequestRes struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      []GetAllInvoiceDetailsRequestResBody
}

type GetAllInvoiceDetailsRequestResBody struct {
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

func GetAllInvoiceDetailsRequest() (*GetAllInvoiceDetailsRequestRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodGet
	isPayload := false
	payload := ""
	url := fmt.Sprintf("%s/invoice/all", client.BaseUrl)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		fmt.Println(err)
	}

	var response GetAllInvoiceDetailsRequestRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

func CancellingInvoiceRequest(invoiceRef string) (*CreateInvoiceRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodDelete
	isPayload := false
	payload := ""
	url := fmt.Sprintf("%s/invoice/cancel?invoiceReference=%s", client.BaseUrl, invoiceRef)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		fmt.Println(err)
	}

	var response CreateInvoiceRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}
