package invoice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/monnify-go"
)

type CreateInvoiceReq struct {
	Amount            int                      `json:"amount"`
	InvoiceReference  string                   `json:"invoiceReference"`
	Description       string                   `json:"description"`
	CurrencyCode      string                   `json:"currencyCode"`
	ContractCode      string                   `json:"contractCode"`
	CustomerEmail     string                   `josn:"customerEmail"`
	CustomerName      string                   `json:"customerName"`
	ExpiryDate        string                   `json:"expiryDate"`
	RedirectUrl       string                   `josn:"redirectUrl"`
	PaymentMethod     string                   `json:"paymentMethod"`
	IncomeSplitConfig IncomeSplitConfigReqBody `json:"incomeSplitConfig"`
}

type IncomeSplitConfigReqBody struct {
	SubAccountCode  string `json:"subAccountCode"`
	SplitPercentage int    `json:"splitPercentage"`
	FeePercentage   int    `json:"feePercentage"`
	FeeBearer       bool   `json:"feeBearer"`
}

type CreateInvoiceRes struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      CreateInvoiceResBody
}

type CreateInvoiceResBody struct {
	Amount           int
	InvoiceReference string
	InvoiceStatus    string
	Description      string
	ContractCode     string
	CustomerEmail    string
	CustomerName     string
	ExpiryDate       string
	CreatedBy        string
	CreatedOn        string
	CheckoutUrl      string
	AccountNumber    string
	AccountName      string
	BankName         string
	BankCode         string
}

func CreateInvoice(payload CreateInvoiceReq) (*CreateInvoiceRes, int, error) {
	client := monnify.NewClient()
	url := fmt.Sprintf("%s/merchant/transactions/init-transaction", client.BaseUrl)
	method := "POST"
	token := fmt.Sprintf("Basic %s", client.BasicToken)

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
	var response CreateInvoiceRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
