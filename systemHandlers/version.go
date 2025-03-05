package systemHandlers

import (
	"context"

	"github.com/aledenza/serco/env"
)

type response struct {
	GoVersion      string `json:"go_version"`
	ServiceVersion string `json:"service_version"`
}

func Version(_ context.Context, _ *struct{}) (*struct{ Body response }, error) {
	return &struct{ Body response }{Body: response{GoVersion: env.GOVERSION(), ServiceVersion: env.VERSION()}}, nil
}
