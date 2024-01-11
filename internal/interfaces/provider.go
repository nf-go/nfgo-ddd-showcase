package interfaces

import (
	"context"
	"nfgo-ddd-showcase/internal/infra"
	"nfgo-ddd-showcase/internal/interfaces/api"
	"nfgo-ddd-showcase/internal/interfaces/job"
	"nfgo-ddd-showcase/internal/interfaces/svc"

	"github.com/nf-go/nfgo/rpc"

	"github.com/google/wire"
	"github.com/nf-go/nfgo"
	"github.com/nf-go/nfgo/ndb"
	"github.com/nf-go/nfgo/njob"
	"github.com/nf-go/nfgo/nmetrics"
	"github.com/nf-go/nfgo/web"
)

var ProviderSet = wire.NewSet(svc.ProviderSet, api.ProviderSet, job.ProviderSet, NewMetricsServer, NewServer)

func NewMetricsServer(config *infra.Config, dbOper ndb.DBOper) nmetrics.Server {
	return nmetrics.MustNewServer(config.Config, nmetrics.DBOption(dbOper.DB(context.Background())))
}

func NewServer(config *infra.Config, metricsServer nmetrics.Server, webServer web.Server, rpcServer rpc.Server, jobServer njob.Server) nfgo.Server {
	server := nfgo.MustNewServer(config.Config,
		metricsServer,
		webServer,
		rpcServer,
		jobServer)
	return server
}
