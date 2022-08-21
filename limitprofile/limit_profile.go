package limitprofile

import (
	"encoding/json"
	"fmt"

	"github.com/hisyntax/monnify-go"
)

type CreateLimitProfileReq struct {
	LimitProfileName       string `json:"limitProfileName"`
	SingleTransactionValue string `json:"singleTransactionValue"`
	DailyTransactionVolume string `json:"dailyTransactionVolume"`
	DailyTransactionValue  string `json:"dailyTransactionValue"`
	LimitProfileCode       string `json:"limitProfileCode"`
}

type CreateLimitProfileRes struct {
	RequestSuccessful bool                      `json:"requestSuccessful"`
	ResponseMessage   string                    `json:"responseMessage"`
	ResponseCode      string                    `json:"responseCode"`
	ResponseBody      CreateLimitProfileReqBody `json:"responseBody"`
}

type CreateLimitProfileReqBody struct {
	LimitProfileCode       string `json:"limitProfileCode"`
	LimitProfileName       string `json:"limitProfileName"`
	SingleTransactionValue string `json:"singleTransactionValue"`
	DailyTransactionVolume string `json:"dailyTransactionVolume"`
	DailyTransactionValue  string `json:"dailyTransactionValue"`
	DateCreated            string `json:"dateCreated"`
	LastModified           string `json:"lastModified"`
}

func CreateLimitProfile(payload CreateLimitProfileReq) (*CreateLimitProfileRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodPost
	isPayload := true
	url := fmt.Sprintf("%s/limit-profile", client.BaseUrl)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}

	var response CreateLimitProfileRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

type GetLimitProfilesRes struct {
	RequestSuccessful bool                      `json:"requestSuccessful"`
	ResponseMessage   string                    `json:"responseMessage"`
	ResponseCode      string                    `json:"responseCode"`
	ResponseBody      []GetLimitProfilesReqBody `json:"responseBody"`
}

type GetLimitProfilesReqBody struct {
	LimitProfileCode       string `json:"limitProfileCode"`
	LimitProfileName       string `json:"limitProfileName"`
	SingleTransactionValue string `json:"singleTransactionValue"`
	DailyTransactionVolume string `json:"dailyTransactionVolume"`
	DailyTransactionValue  string `json:"dailyTransactionValue"`
	DateCreated            string `json:"dateCreated"`
	LastModified           string `json:"lastModified"`
}

func GetLimitProfiles() (*GetLimitProfilesRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodGet
	isPayload := false
	payload := ""
	url := fmt.Sprintf("%s/limit-profile", client.BaseUrl)
	token := fmt.Sprintf("Bearer %s", client.BearerToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}

	var response GetLimitProfilesRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

type UpdateLimitProfileReq struct {
	LimitProfileName       string `json:"limitProfileName"`
	SingleTransactionValue string `json:"singleTransactionValue"`
	DailyTransactionVolume string `json:"dailyTransactionVolume"`
	DailyTransactionValue  string `json:"dailyTransactionValue"`
	LimitProfileCode       string
}

func UpdateLimitProfile(payload UpdateLimitProfileReq) (*CreateLimitProfileRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodUpdate
	isPayload := true
	url := fmt.Sprintf("%s/limit-profile/%s", client.BaseUrl, payload.LimitProfileCode)
	token := fmt.Sprintf("Bearer %s", client.BearerToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}

	var response CreateLimitProfileRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

type ReserveAcctWithLimitReq struct {
	ContractCode     string `json:"contractCode"`
	AccountName      string `json:"accountName"`
	CurrencyCode     string `json:"currencyCode"`
	AccountReference string `json:"accountReference"`
	CustomerEmail    string `json:"customerEmail"`
	CustomerName     string `json:"customerName"`
	LimitProfileCode string `json:"limitProfileCode"`
}

type ReserveAcctWithLimitRes struct {
	RequestSuccessful bool                        `json:"requestSuccessful"`
	ResponseMessage   string                      `json:"responseMessage"`
	ResponseCode      string                      `json:"responseCode"`
	ResponseBody      ReserveAcctWithLimitResBody `json:"responseBody"`
}

type ReserveAcctWithLimitResBody struct {
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
}

func ReserveAcctWithLimit(payload ReserveAcctWithLimitReq) (*ReserveAcctWithLimitRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodPost
	isPayload := true
	url := fmt.Sprintf("%s/bank-transfer/reserved-accounts/limit", client.BaseUrl)
	token := fmt.Sprintf("Bearer %s", client.BearerToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}

	var response ReserveAcctWithLimitRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

type UpdateReserveAcctWithLimitReq struct {
	AccountReference string `json:"accountReference"`
	LimitProfileCode string `json:"limitProfileCode"`
}

func UpdateReserveAcctWithLimit(payload UpdateReserveAcctWithLimitReq) (*ReserveAcctWithLimitRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodUpdate
	isPayload := true
	url := fmt.Sprintf("%s/bank-transfer/reserved-accounts/limit", client.BaseUrl)
	token := fmt.Sprintf("Bearer %s", client.BearerToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}

	var response ReserveAcctWithLimitRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}
