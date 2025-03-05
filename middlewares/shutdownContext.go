package middlewares

import (
	"os/signal"
	"syscall"

	"github.com/danielgtaylor/huma/v2"
)

func ShutdownContext(ctx huma.Context, next func(huma.Context)) {
	c, stop := signal.NotifyContext(ctx.Context(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	ctx = huma.WithContext(ctx, c)
	next(ctx)
}
