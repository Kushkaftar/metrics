package client

import (
	"io"
	"net/http"
)

func (c Client) Get(url string) (*Response, error) {
	req, err := http.NewRequest(
		http.MethodGet, url, nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set(authHeader, c.token)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := Response{}
	response.Header = resp.Header
	response.Body = body
	response.StatusCode = resp.StatusCode

	return &response, nil
}
