package client

import (
	"context"
	"net/http"

	"github.com/aledenza/serco/utils"
	"github.com/go-resty/resty/v2"
)

func (c *Client) Get(method string) func(ctx context.Context, url string, params ...Get) (Response, error) {
	return func(ctx context.Context, url string, params ...Get) (Response, error) {
		param := utils.GetOptionalValue(params...)
		return c.request(method, http.MethodGet)(
			ctx,
			url,
			request{
				RequestParams: param.RequestParams,
				queryParams:   param.QueryParams,
				headers:       param.Headers,
				url:           url,
				httpMethod:    http.MethodGet,
			},
		)
	}
}

func (c *Client) Post(method string) func(context.Context, string, Post) (Response, error) {
	return func(ctx context.Context, url string, params Post) (Response, error) {
		return c.request(method, http.MethodPost)(
			ctx,
			url,
			request{
				RequestParams: params.RequestParams,
				queryParams:   params.QueryParams,
				headers:       params.Headers,
				json:          params.Json,
				body:          params.Body,
				url:           url,
				httpMethod:    http.MethodPost,
			},
		)
	}
}

func (c *Client) Patch(method string) func(context.Context, string, Patch) (Response, error) {
	return func(ctx context.Context, url string, params Patch) (Response, error) {
		return c.request(method, http.MethodPatch)(
			ctx,
			url,
			request{
				RequestParams: params.RequestParams,
				queryParams:   params.QueryParams,
				headers:       params.Headers,
				json:          params.Json,
				url:           url,
				httpMethod:    http.MethodPatch,
			},
		)
	}
}

func (c *Client) Put(method string) func(context.Context, string, Put) (Response, error) {
	return func(ctx context.Context, url string, params Put) (Response, error) {
		return c.request(method, http.MethodPut)(
			ctx,
			url,
			request{
				RequestParams: params.RequestParams,
				queryParams:   params.QueryParams,
				headers:       params.Headers,
				json:          params.Json,
				body:          params.Body,
				url:           url,
				httpMethod:    http.MethodPut,
			},
		)
	}
}

func (c *Client) Delete(method string) func(context.Context, string, Delete) (Response, error) {
	return func(ctx context.Context, url string, params Delete) (Response, error) {
		return c.request(method, http.MethodDelete)(
			ctx,
			url,
			request{
				RequestParams: params.RequestParams,
				queryParams:   params.QueryParams,
				headers:       params.Headers,
				url:           url,
				httpMethod:    http.MethodDelete,
			},
		)
	}
}

func (c *Client) request(method, httpMethod string) func(context.Context, string, request) (Response, error) {
	var err error
	var statusCode int = 500
	metric := clientMetric(c.name, httpMethod, method)
	return func(ctx context.Context, url string, request request) (Response, error) {
		var response *resty.Response
		defer func() { metric(statusCode, err) }()
		response, err = c.runRequest(ctx, request)
		if response != nil {
			statusCode = response.StatusCode()
		}
		return Response{response: response, StatusCode: statusCode}, err
	}
}
