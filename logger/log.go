package logger

import (
	"context"
	"log/slog"

	"github.com/aledenza/serco/env"
	"github.com/aledenza/serco/utils"
)

func logErr(ctx context.Context, msg string, err error, logFunc func(ctx context.Context, msg string, args ...any)) {
	if err != nil {
		logFunc(ctx, msg, "error", err)
	} else {
		logFunc(ctx, msg)
	}
}

func Debug(ctx context.Context, msg string, err ...error) {
	logErr(ctx, msg, utils.GetOptionalValue(err...), slog.DebugContext)
}

func Info(ctx context.Context, msg string, err ...error) {
	logErr(ctx, msg, utils.GetOptionalValue(err...), slog.InfoContext)
}
func Warning(ctx context.Context, msg string, err ...error) {
	logErr(ctx, msg, utils.GetOptionalValue(err...), slog.WarnContext)
}

func Error(ctx context.Context, msg string, err error) {
	logErr(ctx, msg, err, slog.ErrorContext)
}

func init() {
	var logger slog.Handler = NewJsonHandler(env.LOG_LEVEL())
	slog.SetDefault(slog.New(logger))
}
