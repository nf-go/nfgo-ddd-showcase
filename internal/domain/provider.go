package domain

import (
	"nfgo-ddd-showcase/internal/domain/auth"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(auth.ProviderSet)
