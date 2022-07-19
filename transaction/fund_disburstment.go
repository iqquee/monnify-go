package transaction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/go-monnify"
)

type initiateSingleTransferModel struct {
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

func InitiateSingleTransfer(amount, bankCode, accountNumber int, reference, narration, currency, walletId string) (*initiateSingleTransferRes, int, error) {
	client := monnify.NewClient()
	url := fmt.Sprintf("%s/api/v1/disbursements/single", client.BaseUrl)
	method := "POST"
	token := fmt.Sprintf("Basic %s", client.BasicToken)
	payload := initiateSingleTransferModel{
		Amount:        amount,
		Reference:     reference,
		Narration:     narration,
		BankCode:      bankCode,
		Currency:      currency,
		AccountNumber: accountNumber,
		WalletId:      walletId,
	}

	jsonReq, jsonReqErr := json.Marshal(&payload)
	if jsonReqErr != nil {
		return nil, 0, jsonReqErr
	}

	req, reqErr := http.NewRequest(method, url, bytes.NewBuffer(jsonReq))
	if reqErr != nil {
		return nil, 0, reqErr
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	resp, respErr := client.Http.Do(req)
	if respErr != nil {
		return nil, 0, respErr
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	var response initiateSingleTransferRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}

func GetInitiateSingleTransferStatus(paymentRef string) (*getInitiateSingleTransferStatusRes, int, error) {
	client := monnify.NewClient()
	url := fmt.Sprintf("%s/api/v1/disbursements/single/summary?reference=?reference=%s", client.BaseUrl, paymentRef)
	method := "GET"
	token := fmt.Sprintf("Basic %s", client.BasicToken)

	req, reqErr := http.NewRequest(method, url, nil)
	if reqErr != nil {
		return nil, 0, reqErr
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	resp, err := client.Http.Do(req)
	if err != nil {
		return nil, 0, err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	var response getInitiateSingleTransferStatusRes

	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
