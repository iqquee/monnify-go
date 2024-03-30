![cover](./static/cover.gif)
# monnify-go
Monnify-go is a Go library that allows you to integrate the MONNIFY payment gateway into your Go project.

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
$ go get -u github.com/iqquee/monnify-go
```
2. Import it in your code:
```sh
import "github.com/iqquee/monnify-go"
```
## Note : All methods in this package returns three (3) things:
- [x] An object of the response
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
	"net/http"

	monnify "github.com/iqquee/monnify-go"
)

func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	client := monnify.New(*http.DefaultClient, baseUrl,apiKey, secretKey)

	payload := monnify.AcceptPaymentReq{}
	
	res, err := client.AcceptPayment(payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
```
## Get Accepted Payment Status
Use this to get accepted payment status
```go
package main

import (
	"fmt"
	"net/http"

	monnify "github.com/iqquee/monnify-go"
)

func main() {
	apiKey := ""
	secretKey := ""
	baseUrl := "https://sandbox.monnify.com" // for test
	client := monnify.New(*http.DefaultClient, baseUrl,apiKey, secretKey)

	paymentReference := "ref123"
	res, err := client.GetTransactionStatus(paymentReference)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
```
