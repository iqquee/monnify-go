package transaction

import (
	"encoding/json"
	"fmt"

	"github.com/hisyntax/monnify-go"
)

type InitiateSingleTransferReq struct {
	Amount        int    `json:"amount"`
	Reference     string `json:"reference"`
	Narration     string `json:"narration"`
	BankCode      int    `json:"bankCode"`
	Currency      string `json:"currency"`
	AccountNumber int    `json:"accountNumber"`
	WalletId      string `json:"walletId"`
}

type initiateSingleTransferRes struct {
	RequestSuccessful bool                          `json:"requestSuccessful"`
	ResponseMessage   string                        `json:"responseMessage"`
	ResponseCode      string                        `json:"responseCode"`
	ResponseBody      initiateSingleTransferResBody `json:"responseBody"`
}

type initiateSingleTransferResBody struct {
	Amount      int    `json:"amount"`
	Reference   string `json:"reference"`
	Status      string `json:"status"`
	DateCreated string `json:"dateCreated"`
}

type getInitiateSingleTransferStatusRes struct {
	RequestSuccessful bool                                   `json:"requestSuccessful"`
	ResponseMessage   string                                 `json:"responseMessage"`
	ResponseCode      string                                 `json:"responseCode"`
	ResponseBody      getInitiateSingleTransferStatusResBody `json:"responseBody"`
}

type getInitiateSingleTransferStatusResBody struct {
	Amount        int    `json:"amount"`
	Reference     string `json:"reference"`
	Narration     string `json:"narration"`
	BankCode      string `json:"bankCode"`
	AccountNumber string `json:"accountNumber"`
	Currency      string `json:"currency"`
	AccountName   string `json:"accountName"`
	BankName      string `json:"bankName"`
	DateCreated   string `json:"dateCreated"`
	Fee           string `json:"fee"`
	Status        string `json:"status"`
}

func InitiateSingleTransfer(payload InitiateSingleTransferReq) (*initiateSingleTransferRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodPost
	isPayload := true
	url := fmt.Sprintf("%s/disbursements/single", client.BaseUrl)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}
	var response initiateSingleTransferRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

func GetInitiateSingleTransferStatus(paymentRef string) (*getInitiateSingleTransferStatusRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodGet
	payload := ""
	isPayload := false
	url := fmt.Sprintf("%s/disbursements/single/summary?reference=?reference=%s", client.BaseUrl, paymentRef)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}

	var response getInitiateSingleTransferStatusRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

type GetInitiateBulkTransferStatusRes struct {
	RequestSuccessful bool                                   `json:"requestSuccessful"`
	ResponseMessage   string                                 `json:"responseMessage"`
	ResponseCode      string                                 `json:"responseCode"`
	ResponseBody      []GetInitiateBulkTransferStatusResBody `json:"responseBody"`
}

type GetInitiateBulkTransferStatusResBody struct {
	Amount        int    `json:"amount"`
	Reference     string `json:"reference"`
	Narration     string `json:"narration"`
	BankCode      string `json:"bankCode"`
	AccountNumber string `json:"accountNumber"`
	Currency      string `json:"currency"`
	AccountName   string `json:"accountName"`
	BankName      string `json:"bankName"`
	DateCreated   string `json:"dateCreated"`
	Fee           string `json:"fee"`
	Status        string `json:"status"`
}

func GetInitiateBulkTransferStatus(batchReference string) (*GetInitiateBulkTransferStatusRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodGet
	payload := ""
	isPayload := false
	url := fmt.Sprintf("%s/disbursements/bulk/transactions?batchReference=%s", client.BaseUrl, batchReference)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}

	var response GetInitiateBulkTransferStatusRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

type InitiateBulkTransferReq struct {
	Title                string                                   `json:"title"`
	BatchReference       string                                   `json:"batchReference"`
	Narration            string                                   `json:"narration"`
	WalletId             string                                   `json:"walletId"`
	OnValidationFailure  string                                   `json:"onValidationFailure"`
	NotificationInterval int                                      `json:"notificationInterval"`
	TransactionList      []InitiateBulkTransferReqTransactionList `json:"transactionList"`
}

type InitiateBulkTransferReqTransactionList struct {
	Amount        string `json:"amount"`
	Reference     string `json:"reference"`
	Narration     string `json:"narration"`
	BankCode      string `json:"bankCode"`
	AccountNumber string `json:"accountNumber"`
	Currency      string `json:"currency"`
}

type InitiateBulkTransferRes struct {
	RequestSuccessful bool                        `json:"requestSuccessful"`
	ResponseMessage   string                      `json:"responseMessage"`
	ResponseCode      string                      `json:"responseCode"`
	ResponseBody      InitiateBulkTransferResBody `json:"responseBody"`
}

type InitiateBulkTransferResBody struct {
	TotalAmount       float64 `json:"totalAmount"`
	TotalFee          float64 `json:"totalFee"`
	BatchReference    string  `json:"batchReference"`
	BatchStatus       string  `json:"batchStatus"`
	TotalTransactions int     `json:"totalTransactions"`
	DateCreated       string  `json:"dateCreated"`
}

func InitiateBulkTransfer(payload InitiateBulkTransferReq) (*InitiateBulkTransferRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodPost
	isPayload := true
	url := fmt.Sprintf("%s/disbursements/batch", client.BaseUrl)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}
	var response InitiateBulkTransferRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

type AuthorizeTransferReq struct {
	Reference         string `json:"reference"`
	AuthorizationCode string `json:"authorizationCode"`
}

type AuthorizeTransferRes struct {
	RequestSuccessful bool                     `json:"requestSuccessful"`
	ResponseMessage   string                   `json:"responseMessage"`
	ResponseCode      string                   `json:"responseCode"`
	ResponseBody      AuthorizeTransferResBody `json:"responseBody"`
}

type AuthorizeTransferResBody struct {
	Amount      string `json:"amount"`
	Reference   string `json:"reference"`
	Status      string `json:"status"`
	DateCreated string `json:"dateCreated"`
}

func AuthorizeSingleTransfer(payload AuthorizeTransferReq) (*AuthorizeTransferRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodPost
	isPayload := true
	url := fmt.Sprintf("%s/disbursements/single/validate-otp", client.BaseUrl)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}
	var response AuthorizeTransferRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

func AuthorizeBulkTransfer(payload AuthorizeTransferReq) (*AuthorizeTransferRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodPost
	isPayload := true
	url := fmt.Sprintf("%s/disbursements/batch/validate-otp", client.BaseUrl)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}
	var response AuthorizeTransferRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

type ResendTransferOtpReq struct {
	Reference string `json:"reference"`
}

type ResendTransferOtpRes struct {
	RequestSuccessful bool                     `json:"requestSuccessful"`
	ResponseMessage   string                   `json:"responseMessage"`
	ResponseCode      string                   `json:"responseCode"`
	ResponseBody      ResendTransferOtpResBody `json:"responseBody"`
}

type ResendTransferOtpResBody struct {
	Message string `json:"message"`
}

func ResendSingleTransferOtp(payload ResendTransferOtpReq) (*ResendTransferOtpRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodPost
	isPayload := true
	url := fmt.Sprintf("%s/disbursements/single/resend-otp", client.BaseUrl)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}
	var response ResendTransferOtpRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

func ResendBulkTransferOtp(payload ResendTransferOtpReq) (*ResendTransferOtpRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodPost
	isPayload := true
	url := fmt.Sprintf("%s/disbursements/batch/resend-otp", client.BaseUrl)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}
	var response ResendTransferOtpRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}
