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
- [x] The object of the response
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

	amount := 100 
	paymentReference := "ref123"
	paymentDesc := "test payment"
	currencyCode := "NGN"
	contractCode := ""
	customerName := ""
	customerEmail := ""
	customerNumber := "" 
	redirectUrl := "https://google.com" // test redirect url
	res, status, err := transaction.AcceptPayment(amount,paymentReference , paymentDesc, currencyCode, contractCode, customerName, customerEmail, customerNumber, redirectUrl)
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

	amount := 100
	paymentReference := "ref123"
	narration := "example transaction"
	bankCode := "058" // for GT Bank
	currency := "NGN"
	accountNumber := ""
	walletId := ""
	res, status, err := transaction.InitiateSingleTransfer(amount, paymentReference, narration, currency, bankCode, accountNumber, walletId)
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
