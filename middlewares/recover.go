package middlewares

import (
	"log/slog"

	"github.com/danielgtaylor/huma/v2"
)

func Recover(ctx huma.Context, next func(huma.Context)) {
	defer func() {
		if r := recover(); r != nil {
			panicMetric(ctx.Operation().OperationID)
			slog.Error("Panic occured", "panic", r)
		}
	}()
	next(ctx)
}
