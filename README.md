# go-monnify
A monnify library written in golang to enable the interaction with the monnify API 

# Please ensure to create issues in this repo if :
- You encounter any error while using this package and that issue would be attended to immediately.

# Get Started
- In other to use this package, you need to first create an account with monnify via https://app.monnify.com/create-account 
- After your account have been successfully created, locate the developer option on the bottom left of your dashboard to get your:
1. Api key
2. Secret Key
3. Contract code

# Installation
To install this monnify package, you need to install [Go](https://golang.org/) and set your Go workspace first.
1. You can use the below Go command to install go-monnify
```sh
$ go get -u github.com/hisyntax/go-monnify
```
2. Import it in your code:
```sh
import "github.com/hisyntax/go-monnify"
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
Use this to accept payments from users
```go
package main

import (
	"fmt"
	monnify "github.com/hisyntax/go-monnify"
	"github.com/hisyntax/go-monnify/transaction"
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
	contractCode := "4000910988"
	customerName := "john doe"
	customerEmail := "johbdoe@gmail.com"
	customerNumber := "09132600841" 
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
	monnify "github.com/hisyntax/go-monnify"
	"github.com/hisyntax/go-monnify/transaction"
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