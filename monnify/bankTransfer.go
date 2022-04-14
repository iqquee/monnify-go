package monnify

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PayWithBankTransfer(bearerToken, base_url, acctRef string) {
	bearer_token := fmt.Sprintf("Bearer %s", bearerToken)
	client := Client

	url := fmt.Sprintf("%s/api/v1/merchant/bank-transfer/init-payment", base_url)
	body := []byte(fmt.Sprintf("{\n  \"transactionReference\": \"%s\",\n  \"bankCode\": \"058\"\n}", acctRef))

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))

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
