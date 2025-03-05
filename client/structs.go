package client

import (
	"database/sql"
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

type Response struct {
	sql.Scanner
	StatusCode int
	response   *resty.Response
}

func (s Response) Scan(output any) error {
	return json.Unmarshal(s.response.Body(), output)
}

type RequestParams struct {
	Timeout *float64
	Backoff *float64
	Retry   *int
}

type Get struct {
	QueryParams map[string]string
	Headers     map[string]string
	RequestParams
}

type Post struct {
	QueryParams map[string]string
	Headers     map[string]string
	Json        any
	Body        []byte
	RequestParams
}

type Patch struct {
	QueryParams map[string]string
	Headers     map[string]string
	Json        any
	RequestParams
}

type Put struct {
	QueryParams map[string]string
	Headers     map[string]string
	Json        any
	Body        []byte
	RequestParams
}

type Delete struct {
	QueryParams map[string]string
	Headers     map[string]string
	RequestParams
}
