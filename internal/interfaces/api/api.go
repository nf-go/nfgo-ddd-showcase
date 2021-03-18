package api

import (
	"nfgo-ddd-showcase/internal/interfaces/api/v1/auth"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(auth.NewAuthAPI)
