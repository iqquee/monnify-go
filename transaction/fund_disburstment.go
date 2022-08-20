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
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      initiateSingleTransferResBody
}

type initiateSingleTransferResBody struct {
	Amount      int
	Reference   string
	Status      string
	DateCreated string
}

type getInitiateSingleTransferStatusRes struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      getInitiateSingleTransferStatusResBody
}

type getInitiateSingleTransferStatusResBody struct {
	Amount        int
	Reference     string
	Narration     string
	BankCode      string
	AccountNumber string
	Currency      string
	AccountName   string
	BankName      string
	DateCreated   string
	Fee           string
	Status        string
}

func InitiateSingleTransfer(payload InitiateSingleTransferReq) (*initiateSingleTransferRes, int, error) {
	client := monnify.NewClient()
	method := monnify.MethodPost
	isPayload := true
	url := fmt.Sprintf("%s/disbursements/single", client.BaseUrl)
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	res, status, err := monnify.NewRequest(method, url, token, isPayload, payload)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
	}

	var response getInitiateSingleTransferStatusRes
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, 0, err
	}

	return &response, status, nil
}
