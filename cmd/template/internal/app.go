package app

import (
	"net/http"
	"service/config"
	"service/internal/service"
	"service/internal/view"
	"strconv"

	"github.com/aledenza/serco"
	"github.com/aledenza/serco/database"
	databaseDrivers "github.com/aledenza/serco/database/drivers"
	"github.com/aledenza/serco/env"
	"github.com/aledenza/serco/middlewares"
	"github.com/aledenza/serco/requestId"
	"github.com/aledenza/serco/web"
	"github.com/danielgtaylor/huma/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func CreateServer() *serco.Server {
	config := serco.NewConfig[config.Config]("settings", "config")

	conn, err := database.NewConnection(config.DbConf.URL, databaseDrivers.Pgx, config.DbConf.Options)
	if err != nil {
		panic(err)
	}

	router := &http.ServeMux{}

	port, err := strconv.Atoi(env.APP_PORT())
	if err != nil {
		panic(err)
	}
	userService := service.NewUserService(config.ClientConf, conn)
	userView := view.NewUserView(userService)
	server := serco.NewServer(
		serco.ServerConfig{
			Router:   router,
			AppName:  env.APP_NAME(),
			Version:  env.VERSION(),
			Port:     port,
			DocsType: serco.Swagger,
			Middlewares: huma.Middlewares{
				middlewares.Recover,
				middlewares.ShutdownContext,
				middlewares.RequestId(
					[]middlewares.SearchParam{{LookUp: middlewares.Header, Key: "request_id"}},
					string(requestId.RequestIdKey),
				),
				middlewares.JWTAuth(
					middlewares.JWTAuthConfig{
						PublicKey:       config.Token,
						AllowEmptyToken: true,
						ServiceName:     env.APP_NAME(),
					},
				),
				middlewares.Metric(
					&middlewares.MuxWrapper{M: router},
					promhttp.Handler(),
					middlewares.MetricConfig{Whitelist: []string{"/jwtping"}},
				),
			},
		},
	)

	// register service handlers
	{
		web.Get(
			server,
			"/user/{user_id}",
			userView.GetUser,
			"GetUser",
			serco.HandlerDescriprion{Tags: []string{"User"}},
		)
	}

	return server
}
