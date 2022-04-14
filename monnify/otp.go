package monnify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/go-monnify/helper"
)

type Otp struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      OtpBody
}

type OtpBody struct {
	Message string
}

func ResendTransferOTP(basicToken, base_url, batchRef string) (*Otp, string, error) {
	basic_token := fmt.Sprintf("Basic %s", basicToken)

	client := Client

	body := []byte(fmt.Sprintf("{\n  \"reference\": \"%s\"\n}", batchRef))
	url := fmt.Sprintf("%s/v1/disbursements/single/resend-otp", base_url)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", basic_token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(helper.ServerErr)
		return nil, "", err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

	var otp Otp
	json.Unmarshal(resp_body, &otp)
	return &otp, resp.Status, nil
}
