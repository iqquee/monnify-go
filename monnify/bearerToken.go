package monnify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hisyntax/go-monnify/helper"
)

var Base_url string = "https://sandbox.monnify.com"

func GetBearerToken() {
	api_key := "MK_TEST_GVMPXQZ1FK"
	secret_key := "99QPF9WCMGC488SYS2YWVHWEZ9R2JD7E"
	BearerGenToken, err := BearerToken(Base_url, api_key, secret_key)
	if err != nil {
		log.Println(err)
	}

	log.Println(BearerGenToken)
}

type Token struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      TokenBody
}

type TokenBody struct {
	AccessToken string
	ExpiresIn   int
}

func BearerToken(base_url, apiKey, SecretKey string) (*Token, error) {
	api_secret_key := fmt.Sprintf("%s:%s", apiKey, SecretKey)
	basic_auth := helper.Base64(api_secret_key)

	client := Client
	url := fmt.Sprintf("%s/api/v1/auth/login/", base_url)
	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", basic_auth))

	resp, err := client.Do(req)
	if err != nil {
		log.Println(helper.ServerErr)
		return nil, err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	log.Println(resp.Status)
	log.Println(string(resp_body))

	var token Token
	json.Unmarshal(resp_body, &token)
	return &token, nil
}
