package api

import (
	"nfgo-ddd-showcase/internal/infra"
	"nfgo-ddd-showcase/internal/interfaces/api/v1/auth"

	"github.com/google/wire"
	"nfgo.ga/nfgo/nmetrics"
	"nfgo.ga/nfgo/web"
)

var ProviderSet = wire.NewSet(
	NewWebServer,
	NewRouterRegistrars,
	auth.NewMiddleWare,
	auth.NewAuthAPI,
)

type RouterRegistrar interface {
	RegisterRoutes(rg web.RouterGroup)
}

func NewWebServer(config *infra.Config, metricsServer nmetrics.Server, middleWare *auth.MiddleWare, routerRegistrars []RouterRegistrar) web.Server {
	webServer := web.MustNewServer(config.Config, web.MetricsServerOption(metricsServer), web.MiddlewaresOption(
		middleWare.AuthcAndAuthz,
	))
	// register routes
	rootRg := webServer.Group("/api/v1")
	for _, registrar := range routerRegistrars {
		registrar.RegisterRoutes(rootRg)
	}
	return webServer
}

func NewRouterRegistrars(
	authApi *auth.AuthAPI,
) []RouterRegistrar {
	return []RouterRegistrar{
		authApi,
	}
}
