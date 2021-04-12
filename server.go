package main

import (
	"context"
	"nfgo-ddd-showcase/internal/infra"
	apiauth "nfgo-ddd-showcase/internal/interfaces/api/v1/auth"
	svcauth "nfgo-ddd-showcase/internal/interfaces/svc/auth"

	"nfgo.ga/nfgo"
	"nfgo.ga/nfgo/nconf"
	"nfgo.ga/nfgo/ndb"
	"nfgo.ga/nfgo/njob"
	"nfgo.ga/nfgo/nlog"
	"nfgo.ga/nfgo/nmetrics"
	"nfgo.ga/nfgo/rpc"
	"nfgo.ga/nfgo/web"
)

func NewConfig() (*infra.Config, func()) {
	config := &infra.Config{}
	nconf.MustLoadConfigCustom(configFile, config)
	nlog.InitLogger(config.Config)
	cleanup := func() {
		nlog.Sync()
	}
	return config, cleanup
}

func NewMetricsServer(config *infra.Config, dbOper ndb.DBOper) nmetrics.Server {
	return nmetrics.MustNewServer(config.Config, nmetrics.DBOption(dbOper.DB(context.Background())))
}

func NewRPCServer(config *infra.Config, metricsServer nmetrics.Server, authSvc *svcauth.AuthSvc) rpc.Server {
	server := rpc.MustNewServer(config.Config, rpc.MetricsServerOption(metricsServer))

	// regester svc services
	svcauth.RegisterAuthSvcServer(server.GRPCServer(), authSvc)
	// ...

	return server
}

func NewWebServer(config *infra.Config, metricsServer nmetrics.Server, authApi *apiauth.AuthAPI) web.Server {

	webServer := web.MustNewServer(config.Config, web.MetricsServerOption(metricsServer), web.MiddlewaresOption(
	// authApi.AuthcAndAuthz,
	))

	// register routers
	rg := webServer.Group("/api/v1")
	{
		rg = rg.Group("/auth")
		{
			rg.POST("/login", authApi.Login)
			rg.POST("/register", authApi.Register)
			rg.POST("/users/:id/avatar", authApi.UploadAvatar)
			rg.GET("/roles", authApi.Roles)

		}
	}
	// ...

	return webServer
}

func NewJobServer(config *infra.Config, jobs njob.Jobs) njob.Server {
	jobServer := njob.MustNewServer(config.Config, njob.JobsOption(jobs))
	return jobServer
}

func NewServer(config *infra.Config, metricsServer nmetrics.Server, webServer web.Server, rpcServer rpc.Server, jobServer njob.Server) nfgo.Server {
	server := nfgo.MustNewServer(config.Config,
		metricsServer,
		webServer,
		rpcServer,
		jobServer)
	return server
}
