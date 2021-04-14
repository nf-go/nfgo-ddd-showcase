package interfaces

import (
	"context"
	"nfgo-ddd-showcase/internal/infra"
	"nfgo-ddd-showcase/internal/interfaces/api"
	"nfgo-ddd-showcase/internal/interfaces/job"
	"nfgo-ddd-showcase/internal/interfaces/svc"

	"nfgo.ga/nfgo/rpc"

	"github.com/google/wire"
	"nfgo.ga/nfgo"
	"nfgo.ga/nfgo/ndb"
	"nfgo.ga/nfgo/njob"
	"nfgo.ga/nfgo/nmetrics"
	"nfgo.ga/nfgo/web"
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
