package client

import (
	"bytes"
	"io"
	"net/http"
)

func (c Client) Post(url string, body []uint8) (*Response, error) {

	req, err := http.NewRequest(
		http.MethodPost, url, bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set(authHeader, c.token)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := Response{}
	response.Header = resp.Header
	response.Body = bodyResp
	response.StatusCode = resp.StatusCode

	return &response, nil
}
