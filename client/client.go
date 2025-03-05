package client

import (
	"fmt"
	"log/slog"
	"net/http"
	"slices"
	"time"

	clientErrors "github.com/aledenza/serco/client/errors"
	"github.com/aledenza/serco/utils"
	"github.com/creasty/defaults"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	baseUrl string
	name    string
	config  ClientConfig
	client  *resty.Client
}

func NewClient(baseUrl, name string, config ...ClientConfig) Client {
	cfg := utils.GetOptionalValue(config...)
	if cfg.Bearer != "" {
		cfg.Headers["Authorization"] = fmt.Sprintf("Bearer %s", cfg.Bearer)
	}
	client := Client{baseUrl: baseUrl, name: name, config: cfg}
	client.init()
	return client
}

func (c *Client) init() error {
	if err := defaults.Set(&c.config); err != nil {
		return clientErrors.ClientSetupError{Err: err}
	}
	c.client = resty.NewWithClient(http.DefaultClient).SetBaseURL(c.baseUrl)
	if c.config.Bearer != "" {
		c.client.SetAuthToken(c.config.Bearer)
	}
	if c.config.Headers != nil {
		c.client.SetHeaders(c.config.Headers)
	}
	if c.config.Timeout != 0 {
		c.client.SetTimeout(time.Duration(c.config.Timeout * float64(time.Second)))
	}
	if c.config.Retry != 0 {
		c.client.SetRetryCount(c.config.Retry).AddRetryCondition(func(resp *resty.Response, err error) bool {
			if resp.IsSuccess() || slices.Contains(c.config.IgnoreStatusCodes, resp.StatusCode()) {
				return false
			}
			return true
		}).AddRetryHook(func(resp *resty.Response, err error) {
			slog.WarnContext(resp.Request.Context(), "request error")
		})
	}
	if c.config.Backoff != 0 {
		c.client.SetRetryWaitTime(time.Duration(c.config.Backoff * float64(time.Second)))
	}
	return nil
}

func (c *Client) Shutdown() {
	c.client.GetClient().CloseIdleConnections()
}
