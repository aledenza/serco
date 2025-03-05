package middlewares

import (
	"regexp"

	"github.com/aledenza/serco/utils"
	"github.com/creasty/defaults"
	"github.com/danielgtaylor/huma/v2"
)

type MetricConfig struct {
	Path string `default:"/metrics"`
	// ServiceName string
	Whitelist []string
}

type server[I any] interface {
	Handle(path string, handler I)
}

func Metric[I any](server server[I], handler I, cfg ...MetricConfig) func(huma.Context, func(huma.Context)) {
	config := utils.GetOptionalValue(cfg...)
	defaults.MustSet(&config)
	whiteList := make([]*regexp.Regexp, 0, len(config.Whitelist))
	for _, reg := range config.Whitelist {
		whiteList = append(whiteList, regexp.MustCompile(reg))
	}
	server.Handle(config.Path, handler)
	return func(ctx huma.Context, next func(huma.Context)) {
		for _, route := range whiteList {
			if route.Match([]byte(ctx.URL().Path)) {
				next(ctx)
				return
			}
		}
		metric := apiMetric()
		next(ctx)
		metric(ctx.Method(), ctx.Operation().OperationID, ctx.Status())
	}
}
