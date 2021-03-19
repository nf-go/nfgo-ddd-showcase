package domain

import (
	authrepo "nfgo-ddd-showcase/internal/domain/auth/repo"
	authservice "nfgo-ddd-showcase/internal/domain/auth/service"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(authrepo.ProviderSet, authservice.ProviderSet)
