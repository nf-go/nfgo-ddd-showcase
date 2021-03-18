package repo

import (
	"context"
	"nfgo-ddd-showcase/internal/domain/auth/entity"
	"nfgo-ddd-showcase/internal/infra"

	"nfgo.ga/nfgo/ndb"
)

// AuthRoleRepo -
type AuthRoleRepo interface {
	FindAllRoles(ctx context.Context, withResources bool) ([]*entity.AuthRole, error)
	FindRolesByUser(ctx context.Context, user *entity.AuthUser) ([]*entity.AuthRole, error)
	FindRoles(ctx context.Context, cond *entity.FindRoleCond) ([]*entity.AuthRole, error)
}

func NewAuthRoleRepo(dbOper ndb.DBOper) AuthRoleRepo {
	return &authRoleRepo{
		DBOper: dbOper,
	}
}

type authRoleRepo struct {
	ndb.DBOper
}

func (r *authRoleRepo) FindAllRoles(ctx context.Context, withResources bool) ([]*entity.AuthRole, error) {
	roles := []*entity.AuthRole{}
	db := r.DB(ctx)
	if withResources {
		db = db.Preload("Resources")
	}
	if err := db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *authRoleRepo) FindRolesByUser(ctx context.Context, user *entity.AuthUser) ([]*entity.AuthRole, error) {
	roles := []*entity.AuthRole{}
	if err := r.DB(ctx).Model(user).Association("Roles").Find(&roles); err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *authRoleRepo) FindRoles(ctx context.Context, cond *entity.FindRoleCond) ([]*entity.AuthRole, error) {
	roles := []*entity.AuthRole{}
	if cond == nil {
		return nil, infra.ErrLackFindCond
	}

	db := r.DB(ctx).Model(&entity.AuthRole{})
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
