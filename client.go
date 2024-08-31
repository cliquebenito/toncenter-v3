package toncenter

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	url   string
	resty *resty.Client
}

func NewClient(apiKey string) *Client {

	return &Client{
		url:   SetApiKey(apiKey),
		resty: resty.New(),
	}
}

func SetApiKey(apiKey string) string {
	return fmt.Sprint("https://toncenter.com/api/v3/masterchainInfo?api_key=", apiKey)
}
