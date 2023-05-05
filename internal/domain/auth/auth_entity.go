package auth

import (
	"fmt"
	"nfgo-ddd-showcase/internal/infra"
	"nfgo-ddd-showcase/internal/infra/util"
	"time"

	"nfgo.ga/nfgo/nutil/ncrypto"
	"nfgo.ga/nfgo/x/nsecurity"
)

// AuthResource -
type AuthResource struct {
	util.Model
	Name         string `gorm:"type:varchar(50);"`
	Type         ResourceType
	URL          string `gorm:"type:varchar(100);"`
	Method       string `gorm:"type:varchar(50);"`
	Icon         string `gorm:"type:varchar(50);"`
	TreeLevel    int
	IsLeaf       bool
	DisplayOrder int
	Description  string `gorm:"type:varchar(500);"`
	ParentID     int
	Children     []*AuthResource `gorm:"foreignKey:ParentID"`
}

// AuthRole -
type AuthRole struct {
	util.Model
	Code        string          `gorm:"type:varchar(50);unique_index"`
	Name        string          `gorm:"type:varchar(50);"`
	Description string          `gorm:"type:varchar(500);"`
	Resources   []*AuthResource `gorm:"many2many:auth_role_resource"`
}

// FindRoleCond -
type FindRoleCond struct {
	util.Page
	Code string
	Name string
}

// AuthUser -
type AuthUser struct {
	util.Model
	Username    string `gorm:"type:varchar(50);unique_index"`
	Password    string `gorm:"type:varchar(64);"`
	Salt        string `gorm:"type:varchar(10);"`
	Status      UserStatus
	AvatarImage string `gorm:"type:varchar(100);"`

	Roles []*AuthRole `gorm:"many2many:auth_user_role"`

	SignKey string `gorm:"-"`
	Token   string `gorm:"-"`
}

// Login - 用户登录
func (u *AuthUser) Login(username string, plainPwd string, jwtOper nsecurity.JWTOper) error {
	// 判断用户名是否正确
	if u.Username != username {
		return infra.ErrUsernameOrPwd
	}

	hashedPwd := ncrypto.Sha256(u.Username + u.Salt + plainPwd)
	// 判断密码是否正确
	if u.Password != hashedPwd {
		return infra.ErrUsernameOrPwd
	}
	return u.refreshTokenAndSingKey(jwtOper)
}

// refreshTokenAndSingKey -
func (u *AuthUser) refreshTokenAndSingKey(jwtOper nsecurity.JWTOper) error {
	// 生成token
	token, err := jwtOper.IssueToken(fmt.Sprint(u.ID), time.Now().Add(time.Hour*24*365), map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("fail to generate token for user %s, %w: ", u.Username, err)
	}
	u.Token = token
	// 生成signKey
	u.SignKey, err = ncrypto.UUID()

	return err
}

// Register - 用户注册
func (u *AuthUser) Register() error {
	if u.Password == "" {
		u.Password = "012456"
	}
	u.Status = UserStatusEnabled

	// 生成密码盐值
	u.Salt = ncrypto.RandString(10)
	// Hash存储密码
	u.Password = ncrypto.Sha256(u.Username + u.Salt + u.Password)

	return nil
}
