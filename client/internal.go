package client

import (
	"context"
	"net/http"
	"time"

	clientErrors "github.com/aledenza/serco/client/errors"
	"github.com/aledenza/serco/requestId"
	"github.com/go-resty/resty/v2"
)

type request struct {
	RequestParams
	queryParams map[string]string
	headers     map[string]string
	json        any
	body        []byte
	url         string
	httpMethod  string
}

func (c *Client) runRequest(ctx context.Context, request request) (*resty.Response, error) {
	client := c.client

	{
		if request.Timeout != nil {
			client = client.Clone().SetTimeout(time.Duration(*request.Timeout * float64(time.Second)))
		}
		if request.Retry != nil {
			client = client.Clone().SetRetryCount(*request.Retry)
		}
	}

	httpRequest := client.R().
		SetContext(ctx).
		SetQueryParams(request.queryParams).
		SetHeaders(c.prepareHeaders(ctx, request))

	{
		if request.json != nil && request.body != nil {
			return nil, clientErrors.JsonAndBodyConflict{}
		}
		if request.json != nil && request.body == nil {
			httpRequest.SetBody(request.json)
		}
		if request.json == nil && request.body != nil {
			httpRequest.SetBody(request.body)
		}
	}

	var methodFunc func(string) (*resty.Response, error)
	switch request.httpMethod {
	case http.MethodGet:
		methodFunc = httpRequest.Get
	case http.MethodPost:
		methodFunc = httpRequest.Post
	case http.MethodPatch:
		methodFunc = httpRequest.Patch
	case http.MethodPut:
		methodFunc = httpRequest.Put
	case http.MethodDelete:
		methodFunc = httpRequest.Delete
	}
	if httpResponse, err := methodFunc(request.url); err != nil {
		return nil, clientErrors.SendRequestError{Err: err}
	} else {
		return httpResponse, nil
	}
}

func (c *Client) prepareHeaders(ctx context.Context, request request) map[string]string {
	headers := map[string]string{}
	correlation := requestId.Get(ctx)
	if correlation != nil {
		headers[c.config.RequestId] = *correlation
	}
	for k, v := range c.config.Headers {
		headers[k] = v
	}
	for k, v := range request.headers {
		headers[k] = v
	}
	return headers
}
