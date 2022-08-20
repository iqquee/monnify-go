package account

import (
	"encoding/json"
	"fmt"

	"github.com/hisyntax/monnify-go"
)

type GetReservedAcctTransactionsReq struct {
	AccountReference string
	Page             string
	Size             string
}

type GetReservedAcctTransactionsRes struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      GetReservedAcctTransactionsResBody
}

type GetReservedAcctTransactionsResBody struct {
	Content  []contentBody
	Pageable struct {
		Sort struct {
			Sorted   bool
			Unsorted bool
			Empty    bool
		}
		PageSize   int
		PageNumber int
		Offset     int
		Unpaged    bool
		Paged      bool
	}
	TotalPages    int
	Last          bool
	TotalElements int
	Sort          struct {
		Sorted   bool
		Unsorted bool
		Empty    bool
	}
	First            bool
	NumberOfElements int
	Size             int
	Number           int
	Empty            bool
}

type contentBody struct {
	CustomerDTO          customerDTO
	ProviderAmount       int
	PaymentMethod        string
	CreatedOn            string
	Amount               int
	Flagged              bool
	ProviderCode         string
	Fee                  int
	CurrencyCode         string
	CompletedOn          string
	PaymentDescription   string
	PaymentStatus        string
	TransactionReference string
	PaymentReference     string
	MerchantCode         string
	MerchantName         string
	SettleInstantly      bool
	PayableAmount        int
	AmountPaid           int
	Completed            bool
	PaymentMethodList    interface{}
	CollectionChannel    string
	AccountReference     string
	AccountNumber        string
	CustomerEmail        string
	CustomerName         string
}
type customerDTO struct {
	Email        string
	Name         string
	MerchantCode string
}

func GetReservedAcctTransactions(payload GetReservedAcctTransactionsReq) (*GetReservedAcctTransactionsRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodGet
	isPayload := true
	url := fmt.Sprintf("%s/bank-transfer/reserved-accounts/transactions?accountReference=%s&page=%s&size=%s", client.BaseUrl, payload.AccountReference, payload.Page, payload.Size)
	token := fmt.Sprintf("Bearer %s", client.BearerToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		fmt.Println(err)
	}
	var response GetReservedAcctTransactionsRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

type GetReservedAccountSampleRes struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      GetReservedAccountSampleResBody
}

type GetReservedAccountSampleResBody struct {
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
	Status               string
	CreatedOn            string
	Contract             struct {
		Name                                       string
		Code                                       string
		Description                                interface{}
		SupportsAdvancedSettlementAccountSelection string
	}
	TotalAmount      int
	TransactionCount int
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
		fmt.Println(err)
	}
	var response GetReservedAccountSampleRes

	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

type DeleteReservedAccountSampleRes struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      DeleteReservedAccountSampleResBody
}

type DeleteReservedAccountSampleResBody struct {
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
	Status               string
	CreatedOn            string
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
		fmt.Println(err)
	}
	var response DeleteReservedAccountSampleRes

	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}
