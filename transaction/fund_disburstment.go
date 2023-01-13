package transaction

import (
	"encoding/json"
	"fmt"

	"github.com/iqquee/monnify-go"
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

func GetSingleTransferDetails(paymentRef string) (*getInitiateSingleTransferStatusRes, int, error) {
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

func GetBulkTransferStatus(batchReference string) (*GetInitiateBulkTransferStatusRes, int, error) {
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

type GetBulkTransferDetailsRes struct {
	RequestSuccessful bool                          `json:"requestSuccessful"`
	ResponseMessage   string                        `json:"responseMessage"`
	ResponseCode      string                        `json:"responseCode"`
	ResponseBody      GetBulkTransferDetailsResBody `json:"responseBody"`
}
type GetBulkTransferDetailsResBody struct {
	Title             string  `json:"title"`
	TotalAmount       float64 `json:"totalAmount"`
	TotalFee          float64 `json:"totalFee"`
	BatchReference    string  `json:"batchReference"`
	TotalTransactions int     `json:"totalTransactions"`
	FailedCount       int     `json:"failedCount"`
	SuccessfulCount   int     `json:"successfulCount"`
	PendingCount      string  `json:"pendingCount"`
	BatchStatus       string  `json:"batchStatus"`
	DateCreated       string  `json:"dateCreated"`
}

func GetBulkTransferDetails(batchReference string) (*GetBulkTransferDetailsRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodGet
	payload := ""
	isPayload := false
	url := fmt.Sprintf("%s/disbursements/batch/summary?reference=?reference=%s", client.BaseUrl, batchReference)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}

	var response GetBulkTransferDetailsRes
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

type ListAllSingleTransfersRes struct {
	RequestSuccessful bool                          `json:"requestSuccessful"`
	ResponseMessage   string                        `json:"responseMessage"`
	ResponseCode      string                        `json:"responseCode"`
	ResponseBody      ListAllSingleTransfersResBody `json:"responseBody"`
}

type ListAllSingleTransfersResBody struct {
	Content []ListAllSingleTransfersResBodyContent `json:"content"`
}

type ListAllSingleTransfersResBodyContent struct {
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

func ListAllSingleTransfers(pageSize int) (*ListAllSingleTransfersRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodGet
	payload := ""
	isPayload := false
	url := fmt.Sprintf("%s/disbursements/single/transactions?pageSize=%d", client.BaseUrl, pageSize)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}

	var response ListAllSingleTransfersRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

type ListAllBulkTransfersRes struct {
	RequestSuccessful bool                        `json:"requestSuccessful"`
	ResponseMessage   string                      `json:"responseMessage"`
	ResponseCode      string                      `json:"responseCode"`
	ResponseBody      ListAllBulkTransfersResBody `json:"responseBody"`
}

type ListAllBulkTransfersResBody struct {
	Content []ListAllBulkTransfersResBodyContent `json:"content"`
}

type ListAllBulkTransfersResBodyContent struct {
	TotalAmount       float64 `json:"totalAmount"`
	TotalFee          float64 `json:"totalFee"`
	BatchReference    string  `json:"batchReference"`
	BatchStatus       string  `json:"batchStatus"`
	TotalTransactions int     `json:"totalTransactions"`
	DateCreated       string  `json:"dateCreated"`
}

func ListAllBulkTransfers(pageSize int) (*ListAllBulkTransfersRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodGet
	payload := ""
	isPayload := false
	url := fmt.Sprintf("%s/disbursements/bulk/transactions?pageSize=%d", client.BaseUrl, pageSize)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}

	var response ListAllBulkTransfersRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

type GetDisburstmentWalletBalRes struct {
	RequestSuccessful bool                            `json:"requestSuccessful"`
	ResponseMessage   string                          `json:"responseMessage"`
	ResponseCode      string                          `json:"responseCode"`
	ResponseBody      GetDisburstmentWalletBalResBody `json:"responseBody"`
}

type GetDisburstmentWalletBalResBody struct {
	AvailableBalance int `json:"availableBalance"`
	LedgerBalance    int `json:"ledgerBalance"`
}

func GetDisburstmentWalletBal(walletId int) (*GetDisburstmentWalletBalRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodGet
	payload := ""
	isPayload := false
	url := fmt.Sprintf("%s/disbursements/wallet-balance?walletId=%d", client.BaseUrl, walletId)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}

	var response GetDisburstmentWalletBalRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}

type ValidateBankAcctDetailsRes struct {
	RequestSuccessful bool                           `json:"requestSuccessful"`
	ResponseMessage   string                         `json:"responseMessage"`
	ResponseCode      string                         `json:"responseCode"`
	ResponseBody      ValidateBankAcctDetailsResBody `json:"responseBody"`
}

type ValidateBankAcctDetailsResBody struct {
	AccountNumber string `json:"accountNumber"`
	AccountName   string `json:"accountName"`
	BankCode      string `json:"bankCode"`
}

func ValidateBankAcctDetails(acctNum, bankCode int) (*ValidateBankAcctDetailsRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodGet
	payload := ""
	isPayload := false
	url := fmt.Sprintf("%s/disbursements/account/validate?accountNumber=%d&bankCode=%d", client.BaseUrl, acctNum, bankCode)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		return nil, 0, err
	}

	var response ValidateBankAcctDetailsRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}
