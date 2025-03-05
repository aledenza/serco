package systemHandlers

import (
	"context"
	"fmt"

	"github.com/aledenza/serco/env"
)

func Hello(ctx context.Context, _ *struct{}) (*struct{ Body string }, error) {
	return &struct{ Body string }{Body: fmt.Sprintf("Hello, it's a %s.", env.APP_NAME())}, nil
}
