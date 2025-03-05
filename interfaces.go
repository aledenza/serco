package serco

import "context"

type Scanner interface {
	Scan(ptr any) error
}

type OnStart func(ctx context.Context) error
type OnStop func()

type CanStartUp interface {
	Startup(ctx context.Context) error
}
type CanShutDown interface {
	Shutdown()
}
