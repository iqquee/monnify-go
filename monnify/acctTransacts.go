package monnify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AcctTransaction struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      AcctTransactionsBody
}

type AcctTransactionsBody struct {
	Content          []AcctTransactions
	Pageable         Pageable
	TotalPages       int
	Last             bool
	TotalElements    int
	Sort             Sort
	First            bool
	NumberOfElements int
	Size             int
	Number           int
	Empty            bool
}

type AcctTransactions struct {
	CustomerDTO          CustomerDTO
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
	PaymentMethodList    []string
	CollectionChannel    string
	AccountReference     string
	AccountNumber        string
	CustomerEmail        string
	CustomerName         string
}

type CustomerDTO struct {
	Email        string
	Name         string
	MerchantCode string
}

type Pageable struct {
	Sort       Sort
	PageSize   int
	PageNumber int
	Offset     int
	Unpaged    bool
	Paged      bool
}
type Sort struct {
	Sorted   bool
	Unsorted bool
	Empty    bool
}

func GetAcctsTransact(base_url, bearerToken string) (*AcctTransaction, string, error) {
	bearer_token := fmt.Sprintf("Bearer %s", bearerToken)

	client := Client
	url := fmt.Sprintf("%s/api/v1/bank-transfer/reserved-accounts/transactions", base_url)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer_token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return nil, "", err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	var acctTransac AcctTransaction
	json.Unmarshal(resp_body, &acctTransac)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return &acctTransac, resp.Status, nil
}
