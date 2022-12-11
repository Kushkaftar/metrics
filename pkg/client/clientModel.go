package client

import "net/http"

type Response struct {
	Header     http.Header
	Body       []uint8
	StatusCode int
}
