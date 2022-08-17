package account

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
	url := fmt.Sprintf("%s/bank-transfer/reserved-accounts/transactions?accountReference=%s&page=%s&size=%s", client.BaseUrl, payload.AccountReference, payload.Page, payload.Size)
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
	var response GetReservedAcctTransactionsRes

	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
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
	url := fmt.Sprintf("%s/bank-transfer/reserved-accounts/%s", client.BaseUrl, accountRef)
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
	var response GetReservedAccountSampleRes

	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
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
	url := fmt.Sprintf("%s/bank-transfer/reserved-accounts/%s", client.BaseUrl, accountRef)
	method := "DELETE"
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
	var response DeleteReservedAccountSampleRes

	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
