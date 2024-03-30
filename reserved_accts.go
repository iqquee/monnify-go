package monnify

import (
	"fmt"
	"net/http"
)

type (
	// GetReservedAcctTransactionsReq request object
	GetReservedAcctTransactionsReq struct {
		AccountReference string
		Page             string
		Size             string
	}

	// GetReservedAcctTransactionsRes respose object
	GetReservedAcctTransactionsRes struct {
		RequestSuccessful bool   `json:"requestSuccessful"`
		ResponseMessage   string `json:"responseMessage"`
		ResponseCode      string `json:"responseCode"`
		ResponseBody      struct {
			Content []struct {
				CustomerDTO struct {
					Email        string `json:"email"`
					Name         string `json:"name"`
					MerchantCode string `json:"merchantCode"`
				} `json:"customerDTO"`
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
			} `json:"content"`
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
		} `json:"responseBody"`
	}

	// GetReservedAccountDetailsRes response object
	GetReservedAccountDetailsRes struct {
		RequestSuccessful bool   `json:"requestSuccessful"`
		ResponseMessage   string `json:"responseMessage"`
		ResponseCode      string `json:"responseCode"`
		ResponseBody      struct {
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
		} `json:"responseBody"`
	}

	// DeleteReservedAccountRes response object
	DeleteReservedAccountRes struct {
		RequestSuccessful bool   `json:"requestSuccessful"`
		ResponseMessage   string `json:"responseMessage"`
		ResponseCode      string `json:"responseCode"`
		ResponseBody      struct {
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
		} `json:"responseBody"`
	}
)

// GetReservedAcctTransactions returns the list of all transactions done on a reserved account.
func (c *Client) GetReservedAcctTransactions(payload GetReservedAcctTransactionsReq) (*GetReservedAcctTransactionsRes, error) {
	url := fmt.Sprintf("%s/v1/bank-transfer/reserved-accounts/transactions?accountReference=%s&page=%s&size=%s", c.baseURL, payload.AccountReference, payload.Page, payload.Size)

	c.isBasic = false
	var response GetReservedAcctTransactionsRes
	if err := c.newRequest(http.MethodGet, url, payload, response); err != nil {
		if err = c.generateNewBearerToken(); err != nil {
			return nil, err
		}
	}

	return &response, nil
}

// GetReservedAccountDetails returns details of an account reserved for a customer
func (c *Client) GetReservedAccountDetails(accountRef string) (*GetReservedAccountDetailsRes, error) {
	url := fmt.Sprintf("%s/v1/bank-transfer/reserved-accounts/%s", c.baseURL, accountRef)

	c.isBasic = false
	var response GetReservedAccountDetailsRes
	if err := c.newRequest(http.MethodGet, url, nil, response); err != nil {
		if err = c.generateNewBearerToken(); err != nil {
			return nil, err
		}
	}

	return &response, nil
}

// DeleteReservedAccount allows you to deallocate/delete already created a reserved account.
func (c *Client) DeleteReservedAccount(accountRef string) (*DeleteReservedAccountRes, error) {
	url := fmt.Sprintf("%s/v1/bank-transfer/reserved-accounts/%s", c.baseURL, accountRef)

	c.isBasic = false
	var response DeleteReservedAccountRes
	if err := c.newRequest(http.MethodDelete, url, nil, response); err != nil {
		if err = c.generateNewBearerToken(); err != nil {
			return nil, err
		}
	}

	return &response, nil
}
