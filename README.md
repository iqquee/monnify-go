![cover](./static/cover.gif)
# monnify-go
Monnify-go is a Go library that allows you to integrate the MONNIFY payment system into your Go project.
Monnify is a payment gateway for businesses to accept payments from customers, either on a recurring or one-time basis. Monnify offers an easier, faster and cheaper way for businesses to get paid on their web and mobile applications using convenient payment methods for customers with the highest success rates obtainable in Nigeria

# Please ensure to create issues in this repo if :
- You encounter any error while using this package and that issue would be attended to immediately.

# Get Started
- In other to use this package, you need to first create an account with monnify via https://app.monnify.com/create-account 
- After your account have been successfully created, locate the developer option at the bottom left of your dashboard to get your:
1. Api key
2. Secret Key
3. Contract code

# Installation
To install this monnify package, you need to install [Go](https://golang.org/) and set your Go workspace first.
1. You can use the below Go command to install monnify-go
```sh
$ go get -u github.com/hisyntax/monnify-go
```
2. Import it in your code:
```sh
import "github.com/hisyntax/monnify-go"
```
## Note : All methods in this package returns three (3) things:
- [x] An object of the response
- [x] An int (status code)
- [x] An error (if any)

# Quick start
```sh
# assume the following codes in example.go file
$ touch example.go
# open the just created example.go file in the text editor of your choice
```
## Accept Payment
Use this to accept payments from customers

### Use this object payload to implement the AcceptPayment() method
Note: CurrencyCode should be "NGN" for naira
```go
type AcceptPaymentReq struct {
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
```
```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/transaction"
)

func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)

	payload := transaction.AcceptPaymentReq{}
	
	res, status, err := transaction.AcceptPayment(payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```
## Get Accepted Payment Status
Use this to get accepted payment status
```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/transaction"
)

func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)

	paymentReference := "ref123"
	res, status, err := transaction.GetTransactionStatus(paymentReference)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```
## Initiate Single Transfer
Use this to initiate single transfers

### Use this object payload to implement the InitiateSingleTransfer() method
```go
type InitiateSingleTransferReq struct {
	Amount        int    `json:"amount"`
	Reference     string `json:"reference"`
	Narration     string `json:"narration"`
	BankCode      int    `json:"bankCode"`
	Currency      string `json:"currency"`
	AccountNumber int    `json:"accountNumber"`
	WalletId      string `json:"walletId"`
}
```
```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/transaction"
)

func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)

	payload := transaction.InitiateSingleTransferReq{}
	res, status, err := transaction.InitiateSingleTransfer(payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```

## Initiate Single Transfer Status
Use this to get the initiated single transfer status
```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/transaction"
)

func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)

	paymentReference := "ref123"
	res, status, err := transaction.GetInitiateSingleTransferStatus(paymentReference)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```

## Initiate Bulk Transfer
Use this to initiate bulk transfer 
### Use this object payload to implement the InitiateBulkTransfer() method
Note: Use BREAK to tell Monnify to reject the entire batch and use CONTINUE to tell Monnify to process the valid transactions fo the OnValidationFailure field
```go
type InitiateBulkTransferReq struct {
	Title                string                                   `json:"title"`
	BatchReference       string                                   `json:"batchReference"`
	Narration            string                                   `json:"narration"`
	WalletId             string                                   `json:"walletId"`
	OnValidationFailure  string                                   `json:"onValidationFailure"`
	NotificationInterval int                                   `json:"notificationInterval"`
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
```

```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/transaction"
)



func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)
	
	payload := transaction.InitiateBulkTransferReq{}
	res, status, err := transaction.InitiateBulkTransfer(payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```

## Pay With Bank Transfer
Use this to make payment using bank ussd code

### Use this object payload to implement the PayWithBank() method
```go
type PayWithBankReq struct {
	TransactionReference string `json:"transactionReference"`
	BankCode             string `json:"bankCode"`
}

```
```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/transaction"
)

func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)

	payload := transaction.PayWithBankReq{}
	res, status, err := transaction.PayWithBank(payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```

## Create Reserved Account
Use this to create a reserved account
### Use this object payload to implement the CreateReservedAcct() method

Note: IncomeSplitConfig is optional
```go
type CreateReservedAcctReq struct {
	AccountName       string `json:"accountName"`
	AccountReference  string `json:"accountReference"`
	CurrencyCode      string `json:"currencyCode"`
	ContractCode      string `json:"contractCode"`
	CustomerName      string `json:"customerName"`
	CustomerEmail     string `json:"customerEmail"`
	IncomeSplitConfig IncomeSplitConfigReqBody `json:"incomeSplitConfig"`
}

type IncomeSplitConfigReqBody struct {
	SubAccountCode  string `json:"subAccountCode"`
	SplitPercentage int    `json:"splitPercentage"`
	FeePercentage   int    `json:"feePercentage"`
	FeeBearer       bool   `json:"feeBearer"`
}
```

```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/account"
)



func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)
	
	//assign the necessary fields with values for the request
	payload := account.CreateReservedAcctReq{}
	res, status, err := account.CreateReservedAcct(payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```

## Get Reserved Account Transactions 
Use this to get reserved account transactions
### Use this object payload to implement the GetReservedAcctTransactions() method

```go
type GetReservedAcctTransactionsReq struct {
	AccountReference string
	Page             string
	Size             string
}
```

```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/account"
)



func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)
	
	//assign the necessary fields with values for the request
	payload := account.GetReservedAcctTransactionsReq{}
	res, status, err := account.GetReservedAcctTransactions(payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```

## Get Reserved Account Sample Request
Use this to get reserved account sample request

```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/account"
)



func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)
	
	accountRef := "red123"
	res, status, err := account.GetReservedAccountSampleRequest(accountRef)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```


## Delete Reserved Account Sample Request
Use this to delete a reserved account sample request

```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/account"
)



func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)
	
	accountRef := "red123"
	res, status, err := account.DeleteReservedAccountSampleRequest(accountRef)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```

## Create Invoices
Use this to create invoices
### Use this object payload to implement the CreateInvoice() method
Note: IncomeSplitConfig is optional

```go
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
```

```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/invoice"
)



func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)
	
	payload := invoice.CreateInvoiceReq{}
	res, status, err := invoice.CreateInvoice(payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```

## Get Invoice Details Request
Use this to get invoice details request

```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/invoice"
)



func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)
	
	invoiceRef := "ref123"
	res, status, err := invoice.GetInvoiceDetailsRequest(invoiceRef)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```

## Get All Invoice Details Request
Use this to get all invoice details request

```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/invoice"
)



func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)
	
	res, status, err := invoice.GetAllInvoiceDetailsRequest()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```


## Cancelling Invoice Request
Use this to cancel invoice request

```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/invoice"
)



func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)
	
	invoiceRef := "ref123"
	res, status, err := invoice.CancellingInvoiceRequest(invoiceRef)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```

## Create Invoice Type Reserved Account
Use this to create invoice type reserved account
### Use this object payload to implement the CreateInvoiceTypeReservedAccount() method
Note: reservedAccountType is "INVOICE"

```go
type CreateInvoiceTypeReservedAccountReq struct {
	ReservedAccountType string `json:"reservedAccountType"`
	AccountReference    string `json:"accountReference"`
	AccountName         string `json:"accountName"`
	CurrencyCode        string `json:"currencyCode"`
	ContractCode        string `json:"contractCode"`
	CustomerEmail       string `json:"customerEmail"`
	CustomerName        string `json:"customerName"`
	CustomerBVN         string `json:"customerBVN"`
}
```

```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/invoice"
)



func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)
	
	payload := invoice.CreateInvoiceTypeReservedAccountReq{}
	res, status, err := invoice.CreateInvoiceTypeReservedAccount(payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```

## Attach Reserved Account To Invoice
Use this to attach a reserved account to an invoice
### Use this object payload to implement the AttachReservedAcctToInvoice() method
Note: This is the format for the ExpiryDate field YYYY-MM-DD HH:MM:SS

```go
type AttachReservedAcctToInvoiceReq struct {
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
```

```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/invoice"
)



func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)
	
	payload := invoice.AttachReservedAcctToInvoiceReq{}
	res, status, err := invoice.AttachReservedAcctToInvoice(payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```

## Create a Sub Account
Use this to create a sub account
### Use this object payload to implement the CreateSubAccount() method

```go
type CreateSubAccountReq struct {
	CurrencyCode           string `json:"currencyCode"`
	AccountNumber          string `json:"accountNumber"`
	AccountName            string `json:"accountName"`
	SubAccountCode         string `json:"subAccountCode"`
	BankCode               string `json:"bankCode"`
	Email                  string `json:"email"`
	DefaultSplitPercentage string `json:"defaultSplitPercentage"`
}
```

```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/subacct"
)



func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)
	
	payload := subacct.CreateSubAccountReq{}
	res, status, err := subacct.CreateSubAccount(payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```

## Get All Sub Accounts
Use this to get all sub accounts

```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/subacct"
)



func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)
	
	res, status, err := subacct.GetAllSubAccounts()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```

## Update a Sub Account
Use this to update a sub account
### Use this object payload to implement the UpdateSubAccount() method

```go
type UpdateSubAccountReq struct {
	CurrencyCode           string `json:"currencyCode"`
	AccountNumber          string `json:"accountNumber"`
	SubAccountCode         string `json:"subAccountCode"`
	BankCode               string `json:"bankCode"`
	Email                  string `json:"email"`
	DefaultSplitPercentage string `json:"defaultSplitPercentage"`
}
```

```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/subacct"
)



func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)
	
	payload := subacct.UpdateSubAccountReq{}
	res, status, err := subacct.UpdateSubAccount(payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```

## Delete Sub Account
Use this to delete a sub account

```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/subacct"
)



func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)
	
	subAcctCode := "ref123"
	res, status, err := subacct.DeleteSubAccount(subAcctCode)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```

## Create Limit Profile
Use this to create a limit profile
### Use this object payload to implement the CreateLimitProfile() method

```go
type CreateLimitProfileReq struct {
	LimitProfileName       string `json:"limitProfileName"`
	SingleTransactionValue string `json:"singleTransactionValue"`
	DailyTransactionVolume string `json:"dailyTransactionVolume"`
	DailyTransactionValue  string `json:"dailyTransactionValue"`
	LimitProfileCode       string `json:"limitProfileCode"`
}
```

```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/limitprofile"
)



func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)
	
	payload := limitprofile.CreateLimitProfileReq{}
	res, status, err := limitprofile.CreateLimitProfile(payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```

## Get Limit Profiles
Use this to get limit profiles

```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/limitprofile"
)



func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)
	
	res, status, err := limitprofile.GetLimitProfiles()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```

## Update Limit Profile
Use this to update a limit profile
### Use this object payload to implement the CreateLimitProfile() method

```go
type UpdateLimitProfileReq struct {
	LimitProfileName       string `json:"limitProfileName"`
	SingleTransactionValue string `json:"singleTransactionValue"`
	DailyTransactionVolume string `json:"dailyTransactionVolume"`
	DailyTransactionValue  string `json:"dailyTransactionValue"`
	LimitProfileCode       string
}
```

```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/limitprofile"
)



func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)
	
	payload := limitprofile.UpdateLimitProfileReq{}
	res, status, err := limitprofile.UpdateLimitProfile(payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```

## Reserve Account with Limit
Use this to reserve account with limit
### Use this object payload to implement the ReserveAcctWithLimit() method

```go
type ReserveAcctWithLimitReq struct {
	ContractCode     string `json:"contractCode"`
	AccountName      string `json:"accountName"`
	CurrencyCode     string `json:"currencyCode"`
	AccountReference string `json:"accountReference"`
	CustomerEmail    string `json:"customerEmail"`
	CustomerName     string `json:"customerName"`
	LimitProfileCode string `json:"limitProfileCode"`
}
```

```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/limitprofile"
)



func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)
	
	payload := limitprofile.ReserveAcctWithLimitReq{}
	res, status, err := limitprofile.ReserveAcctWithLimit(payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```

## Update Reserve Account with Limit
Use this to update reserve account with limit
### Use this object payload to implement the UpdateReserveAcctWithLimit() method

```go
type UpdateReserveAcctWithLimitReq struct {
	AccountReference string `json:"accountReference"`
	LimitProfileCode string `json:"limitProfileCode"`
}
```

```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/monnify-go"
	"github.com/hisyntax/monnify-go/limitprofile"
)



func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	monnify.Options(apiKey, secretKey, baseUrl)
	
	payload := limitprofile.UpdateReserveAcctWithLimitReq{}
	res, status, err := limitprofile.UpdateReserveAcctWithLimit(payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(res)
}
```