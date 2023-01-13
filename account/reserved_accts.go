package account

import (
	"encoding/json"
	"fmt"

	"github.com/iqquee/monnify-go"
)

type GetReservedAcctTransactionsReq struct {
	AccountReference string
	Page             string
	Size             string
}

type GetReservedAcctTransactionsRes struct {
	RequestSuccessful bool                               `json:"requestSuccessful"`
	ResponseMessage   string                             `json:"responseMessage"`
	ResponseCode      string                             `json:"responseCode"`
	ResponseBody      GetReservedAcctTransactionsResBody `json:"responseBody"`
}

type GetReservedAcctTransactionsResBody struct {
	Content  []contentBody `json:"content"`
	Pageable struct {
		Sort struct {
			Sorted   bool `json:"sorted"`
			Unsorted bool `json:"unsorted"`
			Empty    bool `json:"empty"`
		} `json:"sort"`
		PageSize   int  `json:"pagesize"`
		PageNumber int  `json:"pagenumber"`
		Offset     int  `json:"offset"`
		Unpaged    bool `json:"unpaged"`
		Paged      bool `json:"paged"`
	} `json:"pageable"`
	TotalPages    int  `json:"totalPages"`
	Last          bool `json:"last"`
	TotalElements int  `json:"totalElements"`
	Sort          struct {
		Sorted   bool `json:"sorted"`
		Unsorted bool `json:"unsorted"`
		Empty    bool `json:"empty"`
	} `json:"sort"`
	First            bool `json:"first"`
	NumberOfElements int  `json:"numberOfElements"`
	Size             int  `json:"size"`
	Number           int  `json:"number"`
	Empty            bool `json:"empty"`
}

type contentBody struct {
	CustomerDTO          customerDTO `json:"customerDTO"`
	ProviderAmount       int         `json:"providerAmount"`
	PaymentMethod        string      `json:"paymentMethod"`
	CreatedOn            string      `json:"createdOn"`
	Amount               int         `json:"amount"`
	Flagged              bool        `json:"flagged"`
	ProviderCode         string      `json:"providerCode"`
	Fee                  int         `json:"fee"`
	CurrencyCode         string      `json:"currencyCode"`
	CompletedOn          string      `json:"completedOn"`
	PaymentDescription   string      `json:"paymentDescription"`
	PaymentStatus        string      `json:"paymentStatus"`
	TransactionReference string      `json:"transactionReference"`
	PaymentReference     string      `json:"paymentReference"`
	MerchantCode         string      `json:"merchantCode"`
	MerchantName         string      `json:"merchantName"`
	SettleInstantly      bool        `json:"settleInstantly"`
	PayableAmount        int         `json:"payableAmount"`
	AmountPaid           int         `json:"amountPaid"`
	Completed            bool        `json:"completed"`
	PaymentMethodList    interface{} `json:"paymentMethodList"`
	CollectionChannel    string      `json:"collectionChannel"`
	AccountReference     string      `json:"accountReference"`
	AccountNumber        string      `json:"accountNumber"`
	CustomerEmail        string      `json:"customerEmail"`
	CustomerName         string      `json:"customerName"`
}
type customerDTO struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	MerchantCode string `json:"merchantCode"`
}

func GetReservedAcctTransactions(payload GetReservedAcctTransactionsReq) (*GetReservedAcctTransactionsRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodGet
	isPayload := true
	url := fmt.Sprintf("%s/bank-transfer/reserved-accounts/transactions?accountReference=%s&page=%s&size=%s", client.BaseUrl, payload.AccountReference, payload.Page, payload.Size)
	token := fmt.Sprintf("Bearer %s", client.BearerToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}
	var response GetReservedAcctTransactionsRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

type GetReservedAccountSampleRes struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      GetReservedAccountSampleResBody
}

type GetReservedAccountSampleResBody struct {
	ContractCode         string `json:"contractCode"`
	AccountReference     string `json:"accountReference"`
	AccountName          string `json:"accountName"`
	CurrencyCode         string `json:"currencyCode"`
	CustomerEmail        string `json:"customerEmail"`
	CustomerName         string `json:"customerName"`
	AccountNumber        string `json:"accountNumber"`
	BankName             string `json:"bankName"`
	BankCode             string `json:"bankCode"`
	ReservationReference string `json:"reservationReference"`
	Status               string `json:"status"`
	CreatedOn            string `json:"createdOn"`
	Contract             struct {
		Name                                       string      `json:"name"`
		Code                                       string      `json:"code"`
		Description                                interface{} `json:"description"`
		SupportsAdvancedSettlementAccountSelection string      `json:"supportsAdvancedSettlementAccountSelection"`
	} `json:"contract"`
	TotalAmount      int `json:"totalAmount"`
	TransactionCount int `json:"transactionCount"`
}

func GetReservedAccountSampleRequest(accountRef string) (*GetReservedAccountSampleRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodGet
	isPayload := false
	payload := ""
	url := fmt.Sprintf("%s/bank-transfer/reserved-accounts/%s", client.BaseUrl, accountRef)
	token := fmt.Sprintf("Bearer %s", client.BearerToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}
	var response GetReservedAccountSampleRes

	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

type DeleteReservedAccountSampleRes struct {
	RequestSuccessful bool                               `json:"requestSuccessful"`
	ResponseMessage   string                             `json:"responseMessage"`
	ResponseCode      string                             `json:"responseCode"`
	ResponseBody      DeleteReservedAccountSampleResBody `json:"responseBody"`
}

type DeleteReservedAccountSampleResBody struct {
	ContractCode         string `json:"contractCode"`
	AccountReference     string `json:"accountReference"`
	AccountName          string `json:"accountName"`
	CurrencyCode         string `json:"currencyCode"`
	CustomerEmail        string `json:"customerEmail"`
	CustomerName         string `json:"customerName"`
	AccountNumber        string `json:"accountNumber"`
	BankName             string `json:"bankName"`
	BankCode             string `json:"bankCode"`
	ReservationReference string `json:"reservationReference"`
	Status               string `json:"status"`
	CreatedOn            string `json:"createdOn"`
}

func DeleteReservedAccountSampleRequest(accountRef string) (*DeleteReservedAccountSampleRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodDelete
	isPayload := false
	payload := ""
	url := fmt.Sprintf("%s/bank-transfer/reserved-accounts/%s", client.BaseUrl, accountRef)
	token := fmt.Sprintf("Bearer %s", client.BearerToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}
	var response DeleteReservedAccountSampleRes

	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}
