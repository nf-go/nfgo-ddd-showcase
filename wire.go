//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"nfgo-ddd-showcase/internal/domain"
	"nfgo-ddd-showcase/internal/infra"
	"nfgo-ddd-showcase/internal/interfaces"

	"github.com/google/wire"
	"github.com/nf-go/nfgo"
)

func NewShowcaseServer() (nfgo.Server, func()) {
	panic(wire.Build(
		infra.ProviderSet,
		domain.ProviderSet,
		interfaces.ProviderSet,
		newConfig,
	))
}
