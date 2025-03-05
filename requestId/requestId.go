package requestId

import "context"

type requestIdType string

var RequestIdKey requestIdType

func Get(ctx context.Context) *string {
	requestId := ctx.Value(string(RequestIdKey))
	if requestId == nil {
		return nil
	}
	requestIdString, ok := requestId.(string)
	if !ok {
		return nil
	}
	return &requestIdString
}

func Set(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, RequestIdKey, requestId)
}

func SetKey(key string) {
	RequestIdKey = requestIdType(key)
}

func init() {
	RequestIdKey = "requestId"
}
