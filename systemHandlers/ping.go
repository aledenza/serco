package systemHandlers

import "context"

func Ping(_ context.Context, _ *struct{}) (*struct{ Body string }, error) {
	return &struct{ Body string }{Body: "OK"}, nil
}

func JWTPing(_ context.Context, _ *struct{}) (*struct{ Body string }, error) {
	return &struct{ Body string }{Body: "OK"}, nil
}
