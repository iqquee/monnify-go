package monnify

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

type (
	// authorizationRes response object
	authorizationRes struct {
		RequestSuccessful bool   `json:"requestSuccessful"`
		ResponseMessage   string `json:"responseMessage"`
		ResponseCode      string `json:"responseCode"`
		ResponseBody      struct {
			AccessToken string `json:"accessToken"`
			ExpiresIn   int    `json:"expiresIn"`
		} `json:"responseBody"`
	}
)

// basicTokenGen generates a basic token by converting into base64
func (c *Client) basicTokenGen() {
	text := fmt.Sprintf("%s:%s", c.apiKey, c.secretKey)
	encodedText := base64.StdEncoding.EncodeToString([]byte(text))
	c.basicToken = encodedText
}

// bearerTokenGen generates a bearer token using the base64 form of the API Key and secret Key
func (c *Client) bearerTokenGen() error {
	url := fmt.Sprintf("%s/v1/auth/login/", c.baseURL)

	// generate basic token
	c.basicTokenGen()

	// set isBasic
	c.isBasic = true

	var response authorizationRes
	if err := c.newRequest(http.MethodPost, url, nil, response); err != nil {
		return err
	}

	c.bearerToken = response.ResponseBody.AccessToken
	return nil
}

// generateNewBearerToken generates a new bearer token after the previously token has expired
func (c *Client) generateNewBearerToken() error {
	if err := c.bearerTokenGen(); err != nil {
		return err
	}
	return nil
}
