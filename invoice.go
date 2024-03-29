package monnify

import (
	"fmt"
	"net/http"
)

type (
	// CreateInvoiceReq request object
	CreateInvoiceReq struct {
		Amount            int    `json:"amount"`
		InvoiceReference  string `json:"invoiceReference"`
		Description       string `json:"description"`
		CurrencyCode      string `json:"currencyCode"`
		ContractCode      string `json:"contractCode"`
		CustomerEmail     string `josn:"customerEmail"`
		CustomerName      string `json:"customerName"`
		ExpiryDate        string `json:"expiryDate"`
		RedirectUrl       string `josn:"redirectUrl"`
		PaymentMethod     string `json:"paymentMethod"`
		IncomeSplitConfig struct {
			SubAccountCode  string `json:"subAccountCode"`
			SplitPercentage int    `json:"splitPercentage"`
			FeePercentage   int    `json:"feePercentage"`
			FeeBearer       bool   `json:"feeBearer"`
		} `json:"incomeSplitConfig"`
	}

	// CreateInvoiceRes response object
	CreateInvoiceRes struct {
		RequestSuccessful bool   `json:"requestSuccessful"`
		ResponseMessage   string `json:"responseMessage"`
		ResponseCode      string `json:"responseCode"`
		ResponseBody      struct {
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
		} `json:"responseBody"`
	}

	// GetInvoiceDetailsRes response object
	GetInvoiceDetailsRes struct {
		RequestSuccessful bool   `json:"requestSuccessful"`
		ResponseMessage   string `json:"responseMessage"`
		ResponseCode      string `json:"responseCode"`
		ResponseBody      []struct {
			Amount               int    `json:"amount"`
			InvoiceReference     string `json:"invoiceReference"`
			InvoiceStatus        string `json:"invoiceStatus"`
			Description          string `json:"description"`
			ContractCode         string `json:"contractCode"`
			CustomerEmail        string `json:"customerEmail"`
			CustomerName         string `json:"customerName"`
			ExpiryDate           string `json:"expiryDate"`
			CreatedBy            string `json:"createdBy"`
			CreatedOn            string `json:"createdOn"`
			CheckoutUrl          string `json:"checkoutUrl"`
			AccountNumber        string `json:"accountNumber"`
			AccountName          string `json:"accountName"`
			BankName             string `json:"bankName"`
			BankCode             string `json:"bankCode"`
			TransactionReference string `json:"transactionReference"`
			IncomeSplitConfig    struct {
				SubAccountCode  string  `json:"subAccountCode"`
				SplitAmount     int     `json:"splitAmount"`
				FeePercentage   float64 `json:"feePercentage"`
				FeeBearer       bool    `json:"feeBearer"`
				SplitPercentage float64 `json:"splitPercentage"`
			} `json:"incomeSplitConfig"`
		} `json:"responseBody"`
	}

	// GetAllInvoiceDetailsRes response object
	GetAllInvoiceDetailsRes struct {
		RequestSuccessful bool   `json:"requestSuccessful"`
		ResponseMessage   string `json:"responseMessage"`
		ResponseCode      string `json:"responseCode"`
		ResponseBody      []struct {
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
		} `json:"responseBody"`
	}

	// CancelInvoice response object
	CancelInvoiceRes struct {
		RequestSuccessful bool   `json:"requestSuccessful"`
		ResponseMessage   string `json:"responseMessage"`
		ResponseCode      string `json:"responseCode"`
		ResponseBody      struct {
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
		} `json:"responseBody"`
	}

	// CreateInvoiceTypeReservedAccountReq request object
	CreateInvoiceTypeReservedAccountReq struct {
		ReservedAccountType string `json:"reservedAccountType"`
		AccountReference    string `json:"accountReference"`
		AccountName         string `json:"accountName"`
		CurrencyCode        string `json:"currencyCode"`
		ContractCode        string `json:"contractCode"`
		CustomerEmail       string `json:"customerEmail"`
		CustomerName        string `json:"customerName"`
		CustomerBVN         string `json:"customerBVN"`
	}

	// CreateInvoiceTypeReservedAccountRes response object
	CreateInvoiceTypeReservedAccountRes struct {
		RequestSuccessful bool   `json:"requestSuccessful"`
		ResponseMessage   string `json:"responseMessage"`
		ResponseCode      string `json:"responseCode"`
		ResponseBody      struct {
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
		} `json:"responseBody"`
	}

	// AttachReservedAcctToInvoiceReq request object
	AttachReservedAcctToInvoiceReq struct {
		Amount           int    `json:"amount"`
		AccountReference string `json:"accountReference"`
		InvoiceReference string `json:"invoiceReference"`
		Description      string `json:"description"`
		CurrencyCode     string `json:"currencyCode"`
		ContractCode     string `json:"contractCode"`
		CustomerEmail    string `json:"customerEmail"`
		CustomerName     string `json:"customerName"`
		ExpiryDate       string `json:"expiryDate"`
	}

	// AttachReservedAcctToInvoiceRes response object
	AttachReservedAcctToInvoiceRes struct {
		RequestSuccessful bool   `json:"requestSuccessful"`
		ResponseMessage   string `json:"responseMessage"`
		ResponseCode      string `json:"responseCode"`
		ResponseBody      struct {
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
		} `json:"responseBody"`
	}
)

// CreateInvoice creates invoice for payments on your integration.
func (c *Client) CreateInvoice(payload CreateInvoiceReq) (*CreateInvoiceRes, error) {
	url := fmt.Sprintf("%s/invoice/create", c.baseURL)

	c.isBasic = false
	var response CreateInvoiceRes
	if err := c.newRequest(http.MethodPost, url, payload, response); err != nil {
		if err = c.generateNewBearerToken(); err != nil {
			return nil, err
		}
	}

	return &response, nil
}

// GetInvoiceDetails returns details of an invoice on your integration.
func (c *Client) GetInvoiceDetails(invoiceReference string) (*GetInvoiceDetailsRes, error) {
	url := fmt.Sprintf("%s/invoice/%s/details", c.baseURL, invoiceReference)

	c.isBasic = false
	var response GetInvoiceDetailsRes
	if err := c.newRequest(http.MethodGet, url, nil, response); err != nil {
		if err = c.generateNewBearerToken(); err != nil {
			return nil, err
		}
	}

	return &response, nil
}

// GetAllInvoiceDetails returns the list of all the invoice available on your integration.
func (c *Client) GetAllInvoiceDetails() (*GetAllInvoiceDetailsRes, error) {
	url := fmt.Sprintf("%s/invoice/all", c.baseURL)

	c.isBasic = false
	var response GetAllInvoiceDetailsRes
	if err := c.newRequest(http.MethodGet, url, nil, response); err != nil {
		if err = c.generateNewBearerToken(); err != nil {
			return nil, err
		}
	}

	return &response, nil
}

// CancelInvoice cancels an Invoice on your integration
func (c *Client) CancelInvoice(invoiceReference string) (*CancelInvoiceRes, error) {
	url := fmt.Sprintf("%s/invoice/%s/cancel", c.baseURL, invoiceReference)

	c.isBasic = false
	var response CancelInvoiceRes
	if err := c.newRequest(http.MethodDelete, url, nil, response); err != nil {
		if err = c.generateNewBearerToken(); err != nil {
			return nil, err
		}
	}

	return &response, nil
}

// CreateInvoiceTypeReservedAccount allows the creation of an invoiced reserved account.
func (c *Client) CreateInvoiceTypeReservedAccount(payload CreateInvoiceTypeReservedAccountReq) (*CreateInvoiceTypeReservedAccountRes, error) {
	url := fmt.Sprintf("%s/bank-transfer/reserved-accounts", c.baseURL)

	c.isBasic = false
	var response CreateInvoiceTypeReservedAccountRes
	if err := c.newRequest(http.MethodPost, url, payload, response); err != nil {
		if err = c.generateNewBearerToken(); err != nil {
			return nil, err
		}
	}

	return &response, nil
}

// AttachReservedAcctToInvoice attaches a Reserved Account to an Invoice
func (c *Client) AttachReservedAcctToInvoice(payload AttachReservedAcctToInvoiceReq) (*AttachReservedAcctToInvoiceRes, error) {
	url := fmt.Sprintf("%s/invoice/create", c.baseURL)

	c.isBasic = false
	var response AttachReservedAcctToInvoiceRes
	if err := c.newRequest(http.MethodPost, url, payload, response); err != nil {
		if err = c.generateNewBearerToken(); err != nil {
			return nil, err
		}
	}

	return &response, nil
}
