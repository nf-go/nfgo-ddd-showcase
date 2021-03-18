package svc

import (
	"nfgo-ddd-showcase/internal/interfaces/svc/auth"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(auth.NewAuthSvc)
