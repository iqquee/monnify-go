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
