package toncenter

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	apikey string
	resty  *resty.Client
}

func NewClient(apiKey string) *Client {

	return &Client{
		apikey: SetApiKey(apiKey),
		resty:  resty.New(),
	}
}

func SetApiKey(apiKey string) string {
	return fmt.Sprint(apiKey)
}
