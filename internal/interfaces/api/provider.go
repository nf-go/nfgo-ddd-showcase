package api

import (
	"nfgo-ddd-showcase/internal/infra"
	"nfgo-ddd-showcase/internal/interfaces/api/v1/auth"

	"github.com/google/wire"
	"nfgo.ga/nfgo/nmetrics"
	"nfgo.ga/nfgo/web"
)

var ProviderSet = wire.NewSet(NewWebServer, auth.NewAuthAPI)

func NewWebServer(config *infra.Config, metricsServer nmetrics.Server, authApi *auth.AuthAPI) web.Server {

	webServer := web.MustNewServer(config.Config, web.MetricsServerOption(metricsServer), web.MiddlewaresOption(
		authApi.AuthcAndAuthz,
	))

	// register routers
	rg := webServer.Group("/api/v1")
	{
		rg = rg.Group("/auth")
		{
			rg.POST("/login", authApi.Login)
			rg.POST("/register", authApi.Register)
			rg.POST("/users/:id/avatar", authApi.UploadAvatar)
			rg.GET("/roles", authApi.Roles)

		}
	}
	// ...

	return webServer
}
