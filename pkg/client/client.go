package client

import (
	"fmt"
	"net/http"
)

type Client struct {
	HttpClient *http.Client
	token      string
}

const (
	typeHeader             = "Authorization"
	authHeaderContentValue = "application/x-yametrika+json"
	authHeader             = "Authorization"
)

func NewClient(token string) *Client {
	client := Client{}
	client.HttpClient = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}

	client.token = fmt.Sprintf("OAuth %s", token)
	return &client
}
