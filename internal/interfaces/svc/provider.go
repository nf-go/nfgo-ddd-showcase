package svc

import (
	"nfgo-ddd-showcase/internal/infra"
	"nfgo-ddd-showcase/internal/interfaces/svc/auth"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	"nfgo.ga/nfgo/nmetrics"
	"nfgo.ga/nfgo/rpc"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewRPCServer,
	wire.Struct(new(Svcs), "*"),
	auth.NewAuthSvc,
)

type Svcs struct {
	*auth.AuthSvc
}

func NewRPCServer(config *infra.Config, metricsServer nmetrics.Server, svcs *Svcs) rpc.Server {
	server := rpc.MustNewServer(config.Config, rpc.MetricsServerOption(metricsServer))

	// regester svc services
	auth.RegisterAuthSvcServer(server.GRPCServer(), svcs.AuthSvc)
	// ...

	return server
}
