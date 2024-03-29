package monnify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type (
	// Client object config
	Client struct {
		http        *http.Client
		basicToken  string
		bearerToken string
		baseURL     string
		apiKey      string
		secretKey   string
		isBasic     bool
	}
)

// New is the monnify config initializer
func New(http http.Client, baseURL, apiKey, secretKey string) *Client {
	baseUrl := fmt.Sprintf("%s/api/v1", baseURL)
	return &Client{
		http:      &http,
		apiKey:    apiKey,
		secretKey: secretKey,
		baseURL:   baseUrl,
	}
}

// newRequest makes a http request to the monnify's server and decodes the server response into the reqBody parameter passed into the newRequest method
func (c *Client) newRequest(method, reqURL string, reqBody, resp interface{}) error {
	newURL := c.baseURL + reqURL
	var body io.Reader

	if reqBody != nil {
		bb, err := json.Marshal(reqBody)
		if err != nil {
			return err
		}
		body = bytes.NewReader(bb)
	}

	req, err := http.NewRequest(method, newURL, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.isBasic {
		req.Header.Set("Authorization", fmt.Sprintf("Basic %s", c.basicToken))
	} else {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	}

	res, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return err
	}

	return nil
}
