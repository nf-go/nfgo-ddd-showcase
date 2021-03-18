package repo

import (
	"context"
	"nfgo-ddd-showcase/internal/infra"
	"nfgo-ddd-showcase/internal/infra/util"
	"strings"

	"nfgo.ga/nfgo/ndb"
)

// CacheRepo -
type CacheRepo interface {
	SetUserRoles(ctx context.Context, userID int, roles []string) error
	GetUserRoles(ctx context.Context, userID int) ([]string, error)
}

func NewCacheRepo(redisOper ndb.RedisOper, config *infra.Config) CacheRepo {
	return &cacheRepo{
		config:    config,
		RedisOper: redisOper,
	}
}

type cacheRepo struct {
	ndb.RedisOper
	config *infra.Config
}

func (r *cacheRepo) SetUserRoles(ctx context.Context, userID int, roles []string) error {
	key := util.RedisKeyUserRoles.String(userID)
	value := strings.Join(roles, ",")
	return r.SetStringOpts(key, value, false, false, r.config.Security.SignKeyLifeTime)
}

func (r *cacheRepo) GetUserRoles(ctx context.Context, userID int) ([]string, error) {
	str, err := r.GetString(util.RedisKeyUserRoles.String(userID))
	if err != nil {
		return nil, err
	}
	return strings.Split(str, ","), nil
}
