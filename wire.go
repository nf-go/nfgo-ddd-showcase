// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"nfgo-ddd-showcase/internal/domain"
	"nfgo-ddd-showcase/internal/infra"
	"nfgo-ddd-showcase/internal/interfaces"

	"github.com/google/wire"
	"nfgo.ga/nfgo"
)

func NewShowcaseServer() (nfgo.Server, func()) {
	panic(wire.Build(
		infra.ProviderSet,
		domain.ProviderSet,
		interfaces.ProviderSet,
		NewConfig,
		NewMetricsServer,
		NewRPCServer,
		NewWebServer,
		NewJobServer,
		NewServer,
	))
}
