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
	BaseUrl     string
}

var (
	ServerErr = errors.New("error occured while sending request to the server")

	httpClient  http.Client
	basicToken  string
	bearerToken string
	baseUrl     string
	apiKey      string
	secretKey   string
)

func (c Client) BasicTokenGen() {
	text := fmt.Sprintf("%s:%s", apiKey, secretKey)
	encodedText := base64.StdEncoding.EncodeToString([]byte(text))
	basicToken = encodedText
}

func (c Client) BearerTokenGen() (*tokenRes, error) {
	client := httpClient
	url := fmt.Sprintf("%s/api/v1/auth/login/", baseUrl)
	method := "POST"

	req, reqErr := http.NewRequest(method, url, nil)
	if reqErr != nil {
		return nil, reqErr
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", basicToken))

	resp, err := client.Do(req)
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
		log.Println(err)
		return err
	}
	bearerToken = genBearerToken.ResponseBody.AccessToken

	return nil
}

func Options(apikey, secretkey, baseurl string) {
	apiKey = apikey
	secretKey = secretkey
	baseUrl = baseurl
}

func NewClient() *Client {
	var client Client
	client.Token()
	return &Client{
		Http:        httpClient,
		BasicToken:  basicToken,
		BearerToken: bearerToken,
		BaseUrl:     baseUrl,
	}
}
