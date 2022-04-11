package monnify

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hisyntax/go-monnify/helper"
)

func BearerToken() (string, error) {
	client := Client
	envErr := helper.GetEnv()
	if envErr != nil {
		log.Println(helper.EnvErr)
	}
	url := helper.Base_url + "/api/v1/auth/login/"
	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", helper.Basic_Auth)

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
