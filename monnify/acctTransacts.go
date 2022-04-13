package monnify

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// {
// 	"requestSuccessful": true,
// 	"responseMessage": "success",
// 	"responseCode": "0",
// 	"responseBody": {
// 	  "content": [
// 		{
// 		  "customerDTO": {
// 			"email": "tobi@toio.com",
// 			"name": "Mr Tobi",
// 			"merchantCode": "QE5KDSBPD3RN"
// 		  },
// 		  "providerAmount": 10,
// 		  "paymentMethod": "ACCOUNT_TRANSFER",
// 		  "createdOn": "2019-12-10T14:12:00.000+0000",
// 		  "amount": 5000,
// 		  "flagged": false,
// 		  "providerCode": "81884",
// 		  "fee": 10,
// 		  "currencyCode": "NGN",
// 		  "completedOn": "2019-12-10T14:12:00.000+0000",
// 		  "paymentDescription": "Elmer Martin",
// 		  "paymentStatus": "PAID",
// 		  "transactionReference": "MNFY|20191210141200|000247",
// 		  "paymentReference": "MNFY|20191210141200|000247",
// 		  "merchantCode": "QE5KDSBPD3RN",
// 		  "merchantName": "Tobi Limited",
// 		  "settleInstantly": false,
// 		  "payableAmount": 5000,
// 		  "amountPaid": 5000,
// 		  "completed": true,
// 		  "paymentMethodList": [],
// 		  "collectionChannel": "RESERVED_ACCOUNT",
// 		  "accountReference": "L6KHK65ZSZJ23CKTFJKT",
// 		  "accountNumber": "3225593799",
// 		  "customerEmail": "tobi@toio.com",
// 		  "customerName": "Mr Tobi"
// 		},
// 		{
// 		  "customerDTO": {
// 			"email": "tobi@toio.com",
// 			"name": "Mr Tobi",
// 			"merchantCode": "QE5KDSBPD3RN"
// 		  },
// 		  "providerAmount": 10,
// 		  "paymentMethod": "ACCOUNT_TRANSFER",
// 		  "createdOn": "2019-12-10T14:11:46.000+0000",
// 		  "amount": 2500,
// 		  "flagged": false,
// 		  "providerCode": "81884",
// 		  "fee": 10,
// 		  "currencyCode": "NGN",
// 		  "completedOn": "2019-12-10T14:11:46.000+0000",
// 		  "paymentDescription": "Elmer Martin",
// 		  "paymentStatus": "PAID",
// 		  "transactionReference": "MNFY|20191210141146|000246",
// 		  "paymentReference": "MNFY|20191210141146|000246",
// 		  "merchantCode": "QE5KDSBPD3RN",
// 		  "merchantName": "Tobi Limited",
// 		  "settleInstantly": false,
// 		  "payableAmount": 2500,
// 		  "amountPaid": 2500,
// 		  "completed": true,
// 		  "paymentMethodList": [],
// 		  "collectionChannel": "RESERVED_ACCOUNT",
// 		  "accountReference": "L6KHK65ZSZJ23CKTFJKT",
// 		  "accountNumber": "3225593799",
// 		  "customerEmail": "tobi@toio.com",
// 		  "customerName": "Mr Tobi"
// 		},
// 		{
// 		  "customerDTO": {
// 			"email": "tobi@toio.com",
// 			"name": "Mr Tobi",
// 			"merchantCode": "QE5KDSBPD3RN"
// 		  },
// 		  "providerAmount": 10,
// 		  "paymentMethod": "ACCOUNT_TRANSFER",
// 		  "createdOn": "2019-12-10T14:11:25.000+0000",
// 		  "amount": 3000,
// 		  "flagged": false,
// 		  "providerCode": "81884",
// 		  "fee": 10,
// 		  "currencyCode": "NGN",
// 		  "completedOn": "2019-12-10T14:11:26.000+0000",
// 		  "paymentDescription": "Elmer Martin",
// 		  "paymentStatus": "PAID",
// 		  "transactionReference": "MNFY|20191210141125|000245",
// 		  "paymentReference": "MNFY|20191210141125|000245",
// 		  "merchantCode": "QE5KDSBPD3RN",
// 		  "merchantName": "Tobi Limited",
// 		  "settleInstantly": false,
// 		  "payableAmount": 3000,
// 		  "amountPaid": 3000,
// 		  "completed": true,
// 		  "paymentMethodList": [],
// 		  "collectionChannel": "RESERVED_ACCOUNT",
// 		  "accountReference": "L6KHK65ZSZJ23CKTFJKT",
// 		  "accountNumber": "3225593799",
// 		  "customerEmail": "tobi@toio.com",
// 		  "customerName": "Mr Tobi"
// 		}
// 	  ],
// 	  "pageable": {
// 		"sort": {
// 		  "sorted": true,
// 		  "unsorted": false,
// 		  "empty": false
// 		},
// 		"pageSize": 10,
// 		"pageNumber": 0,
// 		"offset": 0,
// 		"unpaged": false,
// 		"paged": true
// 	  },
// 	  "totalPages": 1,
// 	  "last": true,
// 	  "totalElements": 3,
// 	  "sort": {
// 		"sorted": true,
// 		"unsorted": false,
// 		"empty": false
// 	  },
// 	  "first": true,
// 	  "numberOfElements": 3,
// 	  "size": 10,
// 	  "number": 0,
// 	  "empty": false
// 	}
//   }

func GetAcctsTransact(base_url, bearerToken string) {
	bearer_token := fmt.Sprintf("Bearer %s", bearerToken)

	client := Client
	url := fmt.Sprintf("%s/api/v1/bank-transfer/reserved-accounts/transactions", base_url)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer_token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
}
