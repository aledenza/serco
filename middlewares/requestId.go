package middlewares

import (
	"github.com/aledenza/serco/requestId"
	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
)

type LookUp string

const (
	Header LookUp = "Header"
	Query  LookUp = "Query"
)

type SearchParam struct {
	LookUp LookUp
	Key    string
}

func RequestId(
	// Use to search for multiple places for requestId in priority of order
	requestIdKeysInp []SearchParam,
	requestIdCtxKey string) func(huma.Context, func(huma.Context)) {
	requestId.SetKey(requestIdCtxKey)
	return func(ctx huma.Context, next func(huma.Context)) {
		requestId := ""
		for _, param := range requestIdKeysInp {
			var reqId string
			switch param.LookUp {
			case Header:
				reqId = ctx.Header(param.Key)
			case Query:
				reqId = ctx.Query(param.Key)
			}
			if reqId != "" {
				requestId = reqId
				break
			}
		}
		if requestId == "" {
			requestId = uuid.NewString()
		}
		ctx = huma.WithValue(ctx, requestIdCtxKey, requestId)
		next(ctx)
	}
}
