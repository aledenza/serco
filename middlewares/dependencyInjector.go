package middlewares

import (
	"github.com/danielgtaylor/huma/v2"
)

func InjectDependencies(dependencyKey string, dependencies map[string]any) func(huma.Context, func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		ctx = huma.WithValue(ctx, dependencyKey, dependencies)
		next(ctx)
	}
}
