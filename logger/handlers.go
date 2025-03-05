package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"os"
	"runtime"
	"strings"

	"github.com/aledenza/serco/env"
	"github.com/aledenza/serco/requestId"
)

func source(r slog.Record) *slog.Source {
	fs := runtime.CallersFrames([]uintptr{r.PC})
	f, _ := fs.Next()
	if strings.HasSuffix(f.Function, "logger.logErr") {
		var pcs [1]uintptr
		runtime.Callers(7, pcs[:])
		fs = runtime.CallersFrames([]uintptr{pcs[0]})
		f, _ = fs.Next()
	}
	return &slog.Source{
		Function: f.Function,
		File:     f.File,
		Line:     f.Line,
	}
}

var str2lev = map[string]slog.Level{
	"debug":   slog.LevelDebug,
	"info":    slog.LevelInfo,
	"warning": slog.LevelWarn,
	"warn":    slog.LevelWarn,
	"error":   slog.LevelError,
	"err":     slog.LevelError,
}

func levelToLevel(level string) slog.Level {
	level = strings.ToLower(level)
	lvl, ok := str2lev[level]
	if !ok {
		return slog.LevelDebug
	}
	return lvl
}

type JsonHandler struct {
	slog.Handler
	l *log.Logger
}

type fields struct {
	RequestId *string `json:"request_id,omitempty"`
	Levelname string  `json:"levelname,omitempty"`
	Version   string  `json:"version,omitempty"`
	File      string  `json:"file,omitempty"`
	Lineno    int     `json:"lineno,omitempty"`
	Name      string  `json:"name,omitempty"`
	Message   string  `json:"message,omitempty"`
	Created   int64   `json:"created,omitempty"`
	Error     string  `json:"error,omitempty"`
	Panic     string  `json:"panic,omitempty"`
}

func (h *JsonHandler) Handle(ctx context.Context, r slog.Record) error {
	source := source(r)
	var passedError string
	var passedPanic string
	if r.NumAttrs() != 0 {
		r.Attrs(func(a slog.Attr) bool {
			switch a.Key {
			case "error":
				passedError = fmt.Sprint(a.Value)
			case "panic":
				passedPanic = fmt.Sprint(a.Value)
			}
			return true
		})
	}
	b, err := json.Marshal(fields{
		RequestId: requestId.Get(ctx),
		Levelname: r.Level.String(),
		Version:   env.VERSION(),
		File:      source.File,
		Lineno:    source.Line,
		Name:      source.Function,
		Message:   r.Message,
		Created:   r.Time.UnixMilli(),
		Error:     passedError,
		Panic:     passedPanic,
	})
	if err != nil {
		panic(err)
	}
	h.l.Println(string(b))
	return nil
}

func NewJsonHandler(level string) *JsonHandler {
	h := &JsonHandler{
		Handler: slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{Level: levelToLevel(level)},
		),
		l: log.New(os.Stdout, "", 0),
	}
	return h
}
