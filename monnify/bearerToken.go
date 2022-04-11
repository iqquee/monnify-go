package monnify

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hisyntax/go-monnify/helper"
)

func BearerToken(base_url, basic_auth string) (string, error) {
	client := Client
	url := fmt.Sprintf("%s/api/v1/auth/login/", base_url)
	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", basic_auth)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(helper.ServerErr)
		return "", err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	log.Println(resp.Status)
	log.Println(string(resp_body))

	return string(resp_body), nil
}
