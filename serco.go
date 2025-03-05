package serco

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/aledenza/serco/systemHandlers"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/danielgtaylor/huma/v2/humacli"
)

type HandlerDescriprion struct {
	Description string
	Tags        []string
}

type HumaHandler[I, O any] = func(ctx context.Context, input *I) (output *O, err error)

type ServerConfig struct {
	Router      humago.Mux
	Port        int
	AppName     string
	Version     string
	OnStart     []OnStart
	OnStop      []OnStop
	Middlewares huma.Middlewares
	DocsType    DocsType
}

type Server struct {
	Port   int
	app    humacli.CLI
	Api    huma.API
	router http.Handler
}

func (slf *Server) Run() {
	fmt.Printf("---Running app on localhost:%d---", slf.Port)
	slf.app.Run()
}

func startup(ctx context.Context, repositories []OnStart) {
	for _, repo := range repositories {
		err := repo(ctx)
		if err != nil {
			panic(errors.Join(fmt.Errorf("can't start server"), err))
		}
	}
}
func shutdown(repositories []OnStop) {
	for _, repo := range repositories {
		repo()
	}
}

func NewServer(config ServerConfig) *Server {
	if config.Router == nil {
		panic("router required")
	}
	humaConfig := huma.DefaultConfig(config.AppName, config.Version)
	humaConfig.DocsPath = ""
	humaConfig.CreateHooks = nil
	api := humago.New(config.Router, humaConfig)
	var docsPage []byte
	switch config.DocsType {
	case Swagger:
		docsPage = []byte(swaggerUi)
	case ScalarDocs:
		docsPage = []byte(scalarDocs)
	default:
		docsPage = []byte(spotlightElements)
	}
	config.Router.HandleFunc("GET /docs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write(docsPage)
	})
	// register system handlers
	{
		huma.Register(api, huma.Operation{Method: "GET", Path: "/", Hidden: true}, systemHandlers.Hello)
		huma.Register(api, huma.Operation{Method: "GET", Path: "/version", Hidden: true}, systemHandlers.Version)
		huma.Register(api, huma.Operation{Method: "GET", Path: "/ping", Hidden: true}, systemHandlers.Ping)
		huma.Register(api, huma.Operation{Method: "GET", Path: "/jwtping", Hidden: true}, systemHandlers.JWTPing)
	}
	for _, middleware := range config.Middlewares {
		api.UseMiddleware(middleware)
	}
	cli := humacli.New(func(hooks humacli.Hooks, _ *struct{}) {
		server := http.Server{
			Addr:    fmt.Sprintf(":%d", config.Port),
			Handler: config.Router,
		}

		hooks.OnStart(func() {
			startup(context.Background(), config.OnStart)
			server.ListenAndServe()
		})

		hooks.OnStop(func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			shutdown(config.OnStop)
			server.Shutdown(ctx)
		})
	})
	return &Server{Port: config.Port, app: cli, Api: api, router: config.Router}
}
