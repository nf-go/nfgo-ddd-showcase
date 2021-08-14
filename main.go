//go:generate swag init -g main.go -o ./internal/interfaces/api/docs

package main

import (
	"flag"
	"nfgo-ddd-showcase/internal/infra"
	_ "nfgo-ddd-showcase/internal/interfaces/api/docs"

	_ "github.com/go-resty/resty/v2"
	"nfgo.ga/nfgo/nconf"
	"nfgo.ga/nfgo/nlog"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "c", "", "the config file")
}

// @title NFGO DDD Showcase API
// @version 1.0
// @description This is a sample server go-ddd-showcase server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @BasePath /api

// @securityDefinitions.apikey Token
// @in header
// @name X-Token

// @securityDefinitions.apikey Sub
// @in header
// @name X-Sub

// @securityDefinitions.apikey ClientType
// @in header
// @name X-ClientType

// @securityDefinitions.apikey TraceID
// @in header
// @name X-Trace-ID

// @securityDefinitions.apikey TS
// @in header
// @name X-Ts

// @securityDefinitions.apikey Sig
// @in header
// @name X-Sig
func main() {
	flag.Parse()
	server, cleanup := NewShowcaseServer()
	server.RegisterOnShutdown(func() error {
		cleanup()
		return nil
	})
	server.MustServe()
}

func newConfig() (*infra.Config, func()) {
	config := &infra.Config{}
	nconf.MustLoadConfigCustom(configFile, config)
	nlog.InitLogger(config.Config)
	cleanup := func() {
		nlog.Sync()
	}
	return config, cleanup
}
