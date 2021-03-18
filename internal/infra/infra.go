package infra

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/google/wire"
	"nfgo.ga/nfgo/ndb"
	"nfgo.ga/nfgo/nlog"
	"nfgo.ga/nfgo/x/nsecurity"
)

var ProviderSet = wire.NewSet(NewDBOper, NewTransactional, NewRedisPool, NewRedisOper, NewReplayChecker, NewSignKeyStore, NewEnforcer)

func NewDBOper(config *Config) ndb.DBOper {
	return ndb.NewDBOper(ndb.MustNewDB(config.DB))
}

func NewTransactional(dbOper ndb.DBOper) ndb.Transactional {
	return dbOper.Transactional
}

func NewRedisPool(config *Config) (ndb.RedisPool, func()) {
	pool := ndb.MustNewRedisPool(config.Redis)
	cleanup := func() {
		if err := pool.Close(); err != nil {
			nlog.Error(err)
		}
	}
	return pool, cleanup
}

func NewRedisOper(pool ndb.RedisPool) ndb.RedisOper {
	return ndb.NewRedisOper(pool)
}

func NewReplayChecker(config *Config, redisOper ndb.RedisOper) nsecurity.ReplayChecker {
	return nsecurity.NewRedisReplayChecker(redisOper, config.Security)
}

func NewSignKeyStore(config *Config, redisOper ndb.RedisOper) nsecurity.SignKeyStore {
	return nsecurity.NewRedisSignKeyStore(redisOper, config.Security)
}

func NewEnforcer(config *Config, dbOper ndb.DBOper) casbin.IEnforcer {
	return nsecurity.MustNewEnforcer(config.Security, dbOper.DB(context.Background()))
}
