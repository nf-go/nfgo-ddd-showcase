package api

import (
	"nfgo-ddd-showcase/internal/infra"
	"nfgo-ddd-showcase/internal/interfaces/api/v1/auth"

	"github.com/google/wire"
	"github.com/nf-go/nfgo/nmetrics"
	"github.com/nf-go/nfgo/web"
)

var ProviderSet = wire.NewSet(
	NewWebServer,
	wire.Struct(new(APIs), "*"),
	auth.NewMiddleWare,
	auth.NewAuthAPI,
)

type APIs struct {
	*auth.AuthAPI
}

func NewWebServer(config *infra.Config, metricsServer nmetrics.Server, middleWare *auth.MiddleWare, apis *APIs) web.Server {
	webServer := web.MustNewServer(config.Config, web.MetricsServerOption(metricsServer), web.MiddlewaresOption(
		middleWare.AuthcAndAuthz,
	))
	// register routes
	rootRg := webServer.Group("/api/v1")
	routerRegistrars := web.ReflectToRouterRegistrars(apis)
	for _, registrar := range routerRegistrars {
		registrar.RegisterRoutes(rootRg)
	}
	return webServer
}
