package monnify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/go-monnify/helper"
)

type LimitProfile struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      LimitProfileBody
}

type LimitProfileBody struct {
	LimitProfileCode       string
	LimitProfileName       string
	SingleTransactionValue int
	DailyTransactionVolume int
	DailyTransactionValue  int
	DateCreated            string
	LastModified           string
}

func CreateLimitProfile(basicToken, base_url, limitProfileName string, singleTransactionValue, dailyTransactionVolume, dailyTransactionValue int) (*LimitProfile, string, error) {
	basic_token := fmt.Sprintf("Basic %s", basicToken)

	client := Client

	url := fmt.Sprintf("%s/api/v1/limit-profile", base_url)
	body := []byte(fmt.Sprintf("{\n  \"limitProfileName\": \"%s\",\n  \"singleTransactionValue\": %d,\n  \"dailyTransactionVolume\": %d,\n  \"dailyTransactionValue\": %d\n}", limitProfileName, singleTransactionValue, dailyTransactionVolume, dailyTransactionValue))

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

	var limitProfile LimitProfile
	json.Unmarshal(resp_body, &limitProfile)
	return &limitProfile, resp.Status, nil
}

type LimitProfiles struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      LimitProfilesBody
}

type LimitProfilesBody struct {
	Content []LimitProfilesBodyContent
}

type LimitProfilesBodyContent struct {
	LimitProfileCode       string
	LimitProfileName       string
	SingleTransactionValue int
	DailyTransactionVolume int
	DailyTransactionValue  int
	DateCreated            string
	LastModified           string
}

func GetLimitProfile(bearerToken, base_url string) (*LimitProfiles, string, error) {

	bearer_token := fmt.Sprintf("Bearer %s", bearerToken)

	client := Client

	url := fmt.Sprintf("%s/api/v1/limit-profile", base_url)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer_token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(helper.ServerErr)
		return nil, "", err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

	var limitProfiles LimitProfiles
	json.Unmarshal(resp_body, &limitProfiles)
	return &limitProfiles, resp.Status, nil
}

func UpdateLimitProfile(bearerToken, base_url, limitProfileCode, limitProfileName string, singleTransactionValue, dailyTransactionVolume, dailyTransactionValue int) (*LimitProfile, string, error) {
	bearer_token := fmt.Sprintf("Bearer %s", bearerToken)

	client := Client

	url := fmt.Sprintf("%s/api/v1/limit-profile/%s", base_url, limitProfileCode)
	body := []byte(fmt.Sprintf("{\n  \"limitProfileName\": \"%s\",\n  \"singleTransactionValue\": %d,\n  \"dailyTransactionVolume\": %d,\n  \"dailyTransactionValue\": %d\n}", limitProfileName, singleTransactionValue, dailyTransactionVolume, dailyTransactionValue))
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(body))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer_token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(helper.ServerErr)
		return nil, "", err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

	var limitProfile LimitProfile
	json.Unmarshal(resp_body, &limitProfile)
	return &limitProfile, resp.Status, nil
}
