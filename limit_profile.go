package monnify

import (
	"fmt"
	"net/http"
)

type (
	// CreateLimitProfileReq request object
	CreateLimitProfileReq struct {
		LimitProfileName       string `json:"limitProfileName"`
		SingleTransactionValue string `json:"singleTransactionValue"`
		DailyTransactionVolume string `json:"dailyTransactionVolume"`
		DailyTransactionValue  string `json:"dailyTransactionValue"`
		LimitProfileCode       string `json:"limitProfileCode"`
	}

	// CreateLimitProfileRes response object
	CreateLimitProfileRes struct {
		RequestSuccessful bool   `json:"requestSuccessful"`
		ResponseMessage   string `json:"responseMessage"`
		ResponseCode      string `json:"responseCode"`
		ResponseBody      struct {
			LimitProfileCode       string `json:"limitProfileCode"`
			LimitProfileName       string `json:"limitProfileName"`
			SingleTransactionValue string `json:"singleTransactionValue"`
			DailyTransactionVolume string `json:"dailyTransactionVolume"`
			DailyTransactionValue  string `json:"dailyTransactionValue"`
			DateCreated            string `json:"dateCreated"`
			LastModified           string `json:"lastModified"`
		} `json:"responseBody"`
	}

	// GetLimitProfilesRes response object
	GetLimitProfilesRes struct {
		RequestSuccessful bool   `json:"requestSuccessful"`
		ResponseMessage   string `json:"responseMessage"`
		ResponseCode      string `json:"responseCode"`
		ResponseBody      []struct {
			LimitProfileCode       string `json:"limitProfileCode"`
			LimitProfileName       string `json:"limitProfileName"`
			SingleTransactionValue string `json:"singleTransactionValue"`
			DailyTransactionVolume string `json:"dailyTransactionVolume"`
			DailyTransactionValue  string `json:"dailyTransactionValue"`
			DateCreated            string `json:"dateCreated"`
			LastModified           string `json:"lastModified"`
		} `json:"responseBody"`
	}

	// UpdateLimitProfileReq request object
	UpdateLimitProfileReq struct {
		LimitProfileName       string `json:"limitProfileName"`
		SingleTransactionValue string `json:"singleTransactionValue"`
		DailyTransactionVolume string `json:"dailyTransactionVolume"`
		DailyTransactionValue  string `json:"dailyTransactionValue"`
		LimitProfileCode       string
	}

	// ReserveAcctWithLimitReq request object
	ReserveAcctWithLimitReq struct {
		ContractCode     string `json:"contractCode"`
		AccountName      string `json:"accountName"`
		CurrencyCode     string `json:"currencyCode"`
		AccountReference string `json:"accountReference"`
		CustomerEmail    string `json:"customerEmail"`
		CustomerName     string `json:"customerName"`
		LimitProfileCode string `json:"limitProfileCode"`
	}

	// ReserveAcctWithLimitRes response object
	ReserveAcctWithLimitRes struct {
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
			LimitProfileConfig   struct {
				LimitProfileCode       string `json:"limitProfileCode"`
				SingleTransactionValue int    `json:"singleTransactionValue"`
				DailyTransactionVolume int    `json:"dailyTransactionVolume"`
				DailyTransactionValue  int    `json:"dailyTransactionValue"`
			} `json:"limitProfileConfig"`
		} `json:"responseBody"`
	}

	// UpdateReserveAcctWithLimitReq request object
	UpdateReserveAcctWithLimitReq struct {
		AccountReference string `json:"accountReference"`
		LimitProfileCode string `json:"limitProfileCode"`
	}
)

// CreateLimitProfile creates limit profiles on a customer's account
func (c *Client) CreateLimitProfile(payload CreateLimitProfileReq) (*CreateLimitProfileRes, error) {
	url := fmt.Sprintf("%s/v1/limit-profile", c.baseURL)

	c.isBasic = false
	var response CreateLimitProfileRes
	if err := c.newRequest(http.MethodPut, url, nil, response); err != nil {
		if err = c.generateNewBearerToken(); err != nil {
			return nil, err
		}
	}

	return &response, nil
}

// GetLimitProfiles returns the list of all Limit Profiles that have been created for your customers
func (c *Client) GetLimitProfiles() (*GetLimitProfilesRes, error) {
	url := fmt.Sprintf("%s/v1/limit-profile", c.baseURL)

	c.isBasic = false
	var response GetLimitProfilesRes
	if err := c.newRequest(http.MethodPut, url, nil, response); err != nil {
		if err = c.generateNewBearerToken(); err != nil {
			return nil, err
		}
	}

	return &response, nil
}

// UpdateLimitProfile updates the information on an existing Limit Profile
func (c *Client) UpdateLimitProfile(payload UpdateLimitProfileReq) (*CreateLimitProfileRes, error) {
	url := fmt.Sprintf("%s/v1/limit-profile/%s", c.baseURL, payload.LimitProfileCode)

	c.isBasic = false
	var response CreateLimitProfileRes
	if err := c.newRequest(http.MethodPut, url, payload, response); err != nil {
		if err = c.generateNewBearerToken(); err != nil {
			return nil, err
		}
	}

	return &response, nil
}

// ReserveAcctWithLimit reserves an account for your customers with a transaction limit profile on it
func (c *Client) ReserveAcctWithLimit(payload ReserveAcctWithLimitReq) (*ReserveAcctWithLimitRes, error) {
	url := fmt.Sprintf("%s/v1/bank-transfer/reserved-accounts/limit", c.baseURL)

	c.isBasic = false
	var response ReserveAcctWithLimitRes
	if err := c.newRequest(http.MethodPost, url, payload, response); err != nil {
		if err = c.generateNewBearerToken(); err != nil {
			return nil, err
		}
	}

	return &response, nil
}

// UpdateReserveAcctWithLimit updates the information on an existing Limit Profile for a Reserved Account
func (c *Client) UpdateReserveAcctWithLimit(payload UpdateReserveAcctWithLimitReq) (*ReserveAcctWithLimitRes, error) {
	url := fmt.Sprintf("%s/v1/bank-transfer/reserved-accounts/limit", c.baseURL)

	c.isBasic = false
	var response ReserveAcctWithLimitRes
	if err := c.newRequest(http.MethodPut, url, payload, response); err != nil {
		if err = c.generateNewBearerToken(); err != nil {
			return nil, err
		}
	}

	return &response, nil
}
