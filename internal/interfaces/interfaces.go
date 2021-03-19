package interfaces

import (
	"nfgo-ddd-showcase/internal/interfaces/api"
	"nfgo-ddd-showcase/internal/interfaces/job"
	"nfgo-ddd-showcase/internal/interfaces/svc"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(svc.ProviderSet, api.ProviderSet, job.ProviderSet)
