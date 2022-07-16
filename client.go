package monnify

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type tokenRes struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      tokenResBody
}

type tokenResBody struct {
	AccessToken string
	ExpiresIn   int
}

type options struct {
	ApiKey    string
	SecretKey string
	BaseUrl   string
}

type Client struct {
	Http        http.Client
	BasicToken  string
	BearerToken string
	Options     options
}

var (
	ServerErr = errors.New("error occured while sending request to the server")

	basicToken  string
	bearerToken string
)

func (c Client) BasicTokenGen() {
	text := fmt.Sprintf("%s:%s", c.Options.ApiKey, c.Options.SecretKey)
	encodedText := base64.StdEncoding.EncodeToString([]byte(text))
	basicToken = encodedText
}

func (c Client) BearerTokenGen() (*tokenRes, error) {
	client := NewClient()
	url := fmt.Sprintf("%s/api/v1/auth/login/", client.Options.BaseUrl)
	method := "POST"

	req, reqErr := http.NewRequest(method, url, nil)
	if reqErr != nil {
		return nil, reqErr
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", basicToken))

	resp, err := client.Http.Do(req)
	if err != nil {
		log.Println(ServerErr)
		return nil, err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	log.Println(resp.Status)
	log.Println(string(resp_body))

	var token tokenRes
	json.Unmarshal(resp_body, &token)
	return &token, nil
}

func (c Client) Token() error {
	c.BasicTokenGen()
	genBearerToken, err := c.BearerTokenGen()
	if err != nil {
		return err
	}
	bearerToken = genBearerToken.ResponseBody.AccessToken

	return nil
}

func Options(apiKey, secretKey, baseUrl string) *options {
	return &options{
		ApiKey:    apiKey,
		SecretKey: secretKey,
		BaseUrl:   baseUrl,
	}
}
func NewClient() *Client {
	var client Client
	client.Token()
	return &Client{
		Http:        http.Client{},
		BasicToken:  basicToken,
		BearerToken: bearerToken,
		Options:     options{},
	}
}
