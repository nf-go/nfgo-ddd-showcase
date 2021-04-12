// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"nfgo-ddd-showcase/internal/domain/auth/repo"
	"nfgo-ddd-showcase/internal/domain/auth/service"
	"nfgo-ddd-showcase/internal/infra"
	"nfgo-ddd-showcase/internal/interfaces/api/v1/auth"
	"nfgo-ddd-showcase/internal/interfaces/job"
	auth2 "nfgo-ddd-showcase/internal/interfaces/svc/auth"
	"nfgo.ga/nfgo"
)

import (
	_ "github.com/go-resty/resty/v2"
	_ "nfgo-ddd-showcase/internal/interfaces/api/docs"
)

// Injectors from wire.go:

func NewShowcaseServer() (nfgo.Server, func()) {
	config, cleanup := NewConfig()
	dbOper := infra.NewDBOper(config)
	server := NewMetricsServer(config, dbOper)
	transactional := infra.NewTransactional(dbOper)
	authUserRepo := repo.NewAuthUserRepo(dbOper)
	authRoleRepo := repo.NewAuthRoleRepo(dbOper)
	redisPool, cleanup2 := infra.NewRedisPool(config)
	redisOper := infra.NewRedisOper(redisPool)
	cacheRepo := repo.NewCacheRepo(redisOper, config)
	replayChecker := infra.NewReplayChecker(config, redisOper)
	signKeyStore := infra.NewSignKeyStore(config, redisOper)
	iEnforcer := infra.NewEnforcer(config, dbOper)
	authService := service.NewAuthService(transactional, authUserRepo, authRoleRepo, cacheRepo, replayChecker, signKeyStore, iEnforcer, config)
	authAPI := auth.NewAuthAPI(authService, iEnforcer)
	webServer := NewWebServer(config, server, authAPI)
	authSvc := auth2.NewAuthSvc(authService)
	rpcServer := NewRPCServer(config, server, authSvc)
	demoJob := job.NewDemoJob(authService)
	jobs := job.NewJobs(demoJob)
	njobServer := NewJobServer(config, jobs)
	nfgoServer := NewServer(config, server, webServer, rpcServer, njobServer)
	return nfgoServer, func() {
		cleanup2()
		cleanup()
	}
}
