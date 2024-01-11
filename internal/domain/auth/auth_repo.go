package auth

import (
	"context"
	"errors"
	"nfgo-ddd-showcase/internal/infra"
	"nfgo-ddd-showcase/internal/infra/util"
	"strings"

	"github.com/nf-go/nfgo/ndb"
	"gorm.io/gorm"
)

// AuthUserRepo -
type AuthUserRepo interface {
	GetUserByUsername(ctx context.Context, username string) (*AuthUser, error)

	GetUserByUserID(ctx context.Context, userID int64) (*AuthUser, error)

	UpdateUser(ctx context.Context, user *AuthUser) error

	InsertUser(ctx context.Context, user *AuthUser) error

	UpdateAvatarImage(ctx context.Context, userID int64, avatarImageURL string) error
}

func NewAuthUserRepo(dbOper ndb.DBOper) AuthUserRepo {
	return &authUserRepo{
		DBOper: dbOper,
	}
}

type authUserRepo struct {
	ndb.DBOper
}

func (r *authUserRepo) GetUserByUserID(ctx context.Context, userID int64) (*AuthUser, error) {
	user := &AuthUser{}
	err := r.DB(ctx).Where("id = ?", userID).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *authUserRepo) GetUserByUsername(ctx context.Context, username string) (*AuthUser, error) {
	user := &AuthUser{}
	err := r.DB(ctx).Where("username = ?", username).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *authUserRepo) UpdateUser(ctx context.Context, user *AuthUser) error {
	return r.DB(ctx).Save(user).Error
}

func (r *authUserRepo) InsertUser(ctx context.Context, user *AuthUser) error {
	return r.DB(ctx).Create(user).Error
}

func (r *authUserRepo) UpdateAvatarImage(ctx context.Context, userID int64, avatarImageURL string) error {
	return r.DB(ctx).Model(&AuthUser{}).Where("id = ?", userID).Update("avatar_image", avatarImageURL).Error
}

// AuthRoleRepo -
type AuthRoleRepo interface {
	FindAllRoles(ctx context.Context, withResources bool) ([]*AuthRole, error)
	FindRolesByUser(ctx context.Context, user *AuthUser) ([]*AuthRole, error)
	FindRoles(ctx context.Context, cond *FindRoleCond) ([]*AuthRole, error)
}

func NewAuthRoleRepo(dbOper ndb.DBOper) AuthRoleRepo {
	return &authRoleRepo{
		DBOper: dbOper,
	}
}

type authRoleRepo struct {
	ndb.DBOper
}

func (r *authRoleRepo) FindAllRoles(ctx context.Context, withResources bool) ([]*AuthRole, error) {
	roles := []*AuthRole{}
	db := r.DB(ctx)
	if withResources {
		db = db.Preload("Resources")
	}
	if err := db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *authRoleRepo) FindRolesByUser(ctx context.Context, user *AuthUser) ([]*AuthRole, error) {
	roles := []*AuthRole{}
	if err := r.DB(ctx).Model(user).Association("Roles").Find(&roles); err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *authRoleRepo) FindRoles(ctx context.Context, cond *FindRoleCond) ([]*AuthRole, error) {
	roles := []*AuthRole{}
	if cond == nil {
		return nil, infra.ErrLackFindCond
	}

	db := r.DB(ctx).Model(&AuthRole{})
	if cond.Code != "" {
		db = db.Where("code like ?", "%"+cond.Code+"%")
	}
	if cond.Name != "" {
		db = db.Where("name like ?", "%"+cond.Name+"%")
	}
	if err := db.Count(&cond.Total).Error; err != nil {
		return nil, err
	}
	db = db.Offset(cond.Offset()).Limit(cond.Limit())
	if err := db.Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}

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
