//go:generate wire .

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
	metricsOpt := &nmetrics.ServerOption{
		DB: dbOper.DB(context.Background()),
	}
	return nmetrics.MustNewServer(config.Config, metricsOpt)
}

func NewRPCServer(config *infra.Config, metricsServer nmetrics.Server,
	authSvc *svcauth.AuthSvc) rpc.Server {
	rpcServerOpt := &rpc.ServerOption{MetricsServer: metricsServer}
	server := rpc.MustNewServer(config.Config, rpcServerOpt)

	svcauth.RegisterAuthSvcServer(server.GRPCServer(), authSvc)

	return server
}

func NewWebServer(config *infra.Config, metricsServer nmetrics.Server,
	authApi *apiauth.AuthAPI) web.Server {
	webServerOpts := &web.ServerOption{
		MetricsServer: metricsServer,
		Middlewares:   []web.HandlerFunc{
			// authApi.AuthcAndAuthz,
		},
	}
	webServer := web.MustNewServer(config.Config, webServerOpts)

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

	return webServer
}

func NewJobServer(config *infra.Config, jobs njob.JobFuncs) njob.Server {
	jobServer := njob.MustNewServer(config.Config, &njob.ServerOption{
		JobFuncs: jobs,
	})
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
