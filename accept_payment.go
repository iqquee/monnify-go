package monnify

import (
	"fmt"
	"net/http"
)

type (
	// AcceptPaymentReq request object
	AcceptPaymentReq struct {
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

	// AcceptPaymentRes response object
	AcceptPaymentRes struct {
		RequestSuccessful bool   `json:"requestSuccessful"`
		ResponseMessage   string `json:"responseMessage"`
		ResponseCode      string `json:"responseCode"`
		ResponseBody      struct {
			TransactionReference string   `json:"transactionReference"`
			PaymentReference     string   `json:"paymentReference"`
			MerchantName         string   `json:"merchantName"`
			ApiKey               string   `json:"apiKey"`
			EnabledPaymentMethod []string `json:"enabledPaymentMethod"`
			CheckoutUrl          string   `json:"checkoutUrl"`
		} `json:"responseBody"`
	}

	// GetTransacStatusRes response object
	GetTransacStatusRes struct {
		RequestSuccessful bool   `json:"requestSuccessful"`
		ResponseMessage   string `json:"responseMessage"`
		ResponseCode      string `json:"responseCode"`
		ResponseBody      struct {
			CreatedOn            string  `json:"createdOn"`
			Amount               float64 `json:"amount"`
			CurrencyCode         string  `json:"currencyCode"`
			CustomerName         string  `json:"customerName"`
			CustomerEmail        string  `json:"customerEmail"`
			PaymentDescription   string  `json:"paymentDescription"`
			PaymentStatus        string  `json:"paymentStatus"`
			TransactionReference string  `json:"transactionReference"`
			PaymentReference     string  `json:"paymentReference"`
		} `json:"responseBody"`
	}
)

// AcceptPayment initialises the transaction that would be used for card payments and dynamic transfers
func (c *Client) AcceptPayment(payload AcceptPaymentReq) (*AcceptPaymentRes, error) {
	url := fmt.Sprintf("%s/v1/merchant/transactions/init-transaction", c.baseURL)

	c.isBasic = false
	var response AcceptPaymentRes
	if err := c.newRequest(http.MethodPost, url, payload, response); err != nil {
		if err = c.generateNewBearerToken(); err != nil {
			return nil, err
		}
	}

	return &response, nil
}

// GetTransactionStatus returns the status of a transaction
func (c *Client) GetTransactionStatus(transactionReference string) (*GetTransacStatusRes, error) {
	url := fmt.Sprintf("%s/v2/transactions/%s", c.baseURL, transactionReference)

	c.isBasic = false
	var response GetTransacStatusRes
	if err := c.newRequest(http.MethodGet, url, nil, response); err != nil {
		if err = c.generateNewBearerToken(); err != nil {
			return nil, err
		}
	}

	return &response, nil
}
