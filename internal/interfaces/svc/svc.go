package svc

import (
	"nfgo-ddd-showcase/internal/interfaces/svc/auth"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(auth.NewAuthSvc)
