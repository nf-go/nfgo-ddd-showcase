package repo

import (
	"context"
	"errors"
	"nfgo-ddd-showcase/internal/domain/auth/entity"

	"gorm.io/gorm"
	"nfgo.ga/nfgo/ndb"
)

// AuthUserRepo -
type AuthUserRepo interface {
	GetUserByUsername(ctx context.Context, username string) (*entity.AuthUser, error)

	GetUserByUserID(ctx context.Context, userID int64) (*entity.AuthUser, error)

	UpdateUser(ctx context.Context, user *entity.AuthUser) error

	InsertUser(ctx context.Context, user *entity.AuthUser) error

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

func (r *authUserRepo) GetUserByUserID(ctx context.Context, userID int64) (*entity.AuthUser, error) {
	user := &entity.AuthUser{}
	err := r.DB(ctx).Where("id = ?", userID).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *authUserRepo) GetUserByUsername(ctx context.Context, username string) (*entity.AuthUser, error) {
	user := &entity.AuthUser{}
	err := r.DB(ctx).Where("username = ?", username).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *authUserRepo) UpdateUser(ctx context.Context, user *entity.AuthUser) error {
	return r.DB(ctx).Save(user).Error
}

func (r *authUserRepo) InsertUser(ctx context.Context, user *entity.AuthUser) error {
	return r.DB(ctx).Create(user).Error
}

func (r *authUserRepo) UpdateAvatarImage(ctx context.Context, userID int64, avatarImageURL string) error {
	return r.DB(ctx).Model(&entity.AuthUser{}).Where("id = ?", userID).Update("avatar_image", avatarImageURL).Error
}
