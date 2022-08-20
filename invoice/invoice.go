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
	RequestSuccessful bool                 `json:"requestSuccessful"`
	ResponseMessage   string               `json:"responseMessage"`
	ResponseCode      string               `json:"responseCode"`
	ResponseBody      CreateInvoiceResBody `json:"responseBody"`
}

type CreateInvoiceResBody struct {
	Amount           int    `json:"amount"`
	InvoiceReference string `json:"invoiceReference"`
	InvoiceStatus    string `json:"invoiceStatus"`
	Description      string `json:"description"`
	ContractCode     string `json:"contractCode"`
	CustomerEmail    string `json:"customerEmail"`
	CustomerName     string `json:"customerName"`
	ExpiryDate       string `json:"expiryDate"`
	CreatedBy        string `json:"createdBy"`
	CreatedOn        string `json:"createdOn"`
	CheckoutUrl      string `json:"checkoutUrl"`
	AccountNumber    string `json:"accountNumber"`
	AccountName      string `json:"accountName"`
	BankName         string `json:"bankName"`
	BankCode         string `json:"bankCode"`
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
	RequestSuccessful bool                                 `json:"requestSuccessful"`
	ResponseMessage   string                               `json:"responseMessage"`
	ResponseCode      string                               `json:"responseCode"`
	ResponseBody      []GetAllInvoiceDetailsRequestResBody `json:"responseBody"`
}

type GetAllInvoiceDetailsRequestResBody struct {
	Amount           int    `json:"amount"`
	InvoiceReference string `json:"invoiceReference"`
	InvoiceStatus    string `json:"invoiceStatus"`
	Description      string `json:"description"`
	ContractCode     string `json:"contractCode"`
	CustomerEmail    string `json:"customerEmail"`
	CustomerName     string `json:"customerName"`
	ExpiryDate       string `json:"expiryDate"`
	CreatedBy        string `json:"createdBy"`
	CreatedOn        string `json:"createdOn"`
	CheckoutUrl      string `json:"checkoutUrl"`
	AccountNumber    string `json:"accountNumber"`
	AccountName      string `json:"accountName"`
	BankName         string `json:"bankName"`
	BankCode         string `json:"bankCode"`
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

type CreateInvoiceTypeReservedAccountReq struct {
	ReservedAccountType string `json:"reservedAccountType"`
	AccountReference    string `json:"accountReference"`
	AccountName         string `json:"accountName"`
	CurrencyCode        string `json:"currencyCode"`
	ContractCode        string `json:"contractCode"`
	CustomerEmail       string `json:"customerEmail"`
	CustomerName        string `json:"customerName"`
	CustomerBVN         string `json:"customerBVN"`
}

type CreateInvoiceTypeReservedAccountRes struct {
	RequestSuccessful bool                                    `json:"requestSuccessful"`
	ResponseMessage   string                                  `json:"responseMessage"`
	ResponseCode      string                                  `json:"responseCode"`
	ResponseBody      CreateInvoiceTypeReservedAccountResBody `json:"responseBody"`
}

type CreateInvoiceTypeReservedAccountResBody struct {
	ContractCode         string `json:"contractCode"`
	AccountReference     string `json:"accountReference"`
	AccountName          string `json:"accountName"`
	CurrencyCode         string `json:"currencyCode"`
	CustomerEmail        string `json:"customerEmail"`
	CustomerName         string `json:"name"`
	AccountNumber        string `json:"accountNumber"`
	BankName             string `json:"bankName"`
	BankCode             string `json:"bankCode"`
	ReservationReference string `json:"reservationReference"`
	ReservedAccountType  string `json:"reservedAccountType"`
	Status               string `json:"status"`
	CreatedOn            string `json:"createdOn"`
}

func CreateInvoiceTypeReservedAccount(payload CreateInvoiceTypeReservedAccountReq) (*CreateInvoiceTypeReservedAccountRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodPost
	isPayload := true
	url := fmt.Sprintf("%s/bank-transfer/reserved-accounts", client.BaseUrl)
	token := fmt.Sprintf("Bearer %s", client.BearerToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		fmt.Println(err)
	}

	var response CreateInvoiceTypeReservedAccountRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}
