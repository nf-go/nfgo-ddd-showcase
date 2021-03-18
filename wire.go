// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"nfgo-ddd-showcase/internal/domain/auth/repo"
	"nfgo-ddd-showcase/internal/domain/auth/service"
	"nfgo-ddd-showcase/internal/infra"
	"nfgo-ddd-showcase/internal/interfaces/api"
	"nfgo-ddd-showcase/internal/interfaces/job"
	"nfgo-ddd-showcase/internal/interfaces/svc"
	"nfgo.ga/nfgo"
)

func NewShowcaseServer() (nfgo.Server, func()) {
	panic(wire.Build(
		infra.ProviderSet,
		repo.ProviderSet,
		service.ProviderSet,
		svc.ProviderSet,
		api.ProviderSet,
		job.ProviderSet,
		NewConfig,
		NewMetricsServer,
		NewRPCServer,
		NewWebServer,
		NewJobServer,
		NewServer,
	))
}
