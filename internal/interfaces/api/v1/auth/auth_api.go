package auth

import (
	"fmt"
	"nfgo-ddd-showcase/internal/domain/auth/entity"
	"nfgo-ddd-showcase/internal/domain/auth/service"
	"strconv"

	"github.com/casbin/casbin/v2"
	"nfgo.ga/nfgo/nlog"
	"nfgo.ga/nfgo/web"
)

type AuthAPI struct {
	authService service.AuthService
	enforcer    casbin.IEnforcer
}

func NewAuthAPI(authService service.AuthService, enforcer casbin.IEnforcer) *AuthAPI {
	return &AuthAPI{authService: authService, enforcer: enforcer}
}

// UploadAvatar -
// @Summary 上传头像
// @Description 上传头像
// @Tags auth
// @Accept  multipart/form-data
// @Produce  json
// @Param id path int true "user id"
// @Param file formData file true "avatar image"
// @Success 200 {object} util.APIResult{data=UploadAvatarResp}
// @Router /v1/auth/users/{id}/avatar [post]
// @Security Token
// @Security Sub
func (a *AuthAPI) UploadAvatar(c *web.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Fail(err)
		return
	}

	bytes, filename, err := c.FormFileBytes("file")
	if err != nil {
		c.Fail(err)
		return
	}

	nlog.Logger(c).Infof("id = %d filename = %s", userID, filename)

	avatarURL, err := a.authService.UploadAvatar(c, int64(userID), bytes)
	if err != nil {
		c.Fail(err)
		return
	}

	c.Success(&UploadAvatarResp{
		AvatarURL: avatarURL,
	})

}

// Login -
// @Summary 登录
// @Description 用户登录
// @Tags auth
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   username formData string true "username"
// @Param   password formData string true "password"
// @Success 200 {object} util.APIResult{data=LoginResp}
// @Router /v1/auth/login [post]
func (a *AuthAPI) Login(c *web.Context) {
	req := &LoginReq{}
	if c.Bind(req) != nil {
		return
	}

	user, err := a.authService.Login(c, req.Username, req.Password)
	if err != nil {
		c.Fail(err)
		return
	}

	resp := &LoginResp{
		Token:   user.Token,
		SignKey: user.SignKey,
		Sub:     fmt.Sprint(user.ID),
	}
	c.Success(resp)
}

// Register -
// @Summary 注册
// @Description 用户注册
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   body body RegisterReq true "body"
// @Success 200 {object} util.APIResult{data=RegisterResp}
// @Router /v1/auth/register [post]
func (a *AuthAPI) Register(c *web.Context) {
	req := &RegisterReq{}
	if c.BindJSON(req) != nil {
		return
	}

	user := &entity.AuthUser{
		Username: req.Username,
		Password: req.Password,
	}

	if err := a.authService.Register(c, user); err != nil {
		c.Fail(err)
		return
	}

	c.Success(&RegisterResp{})
}

// Roles -
// @Summary 角色列表
// @Description 查询角色列表
// @Tags auth
// @Produce  json
// @Param name query string false "name"
// @Param code query string false "code"
// @Param pageNo query int true "code"
// @Param pageSize query int true "code"
// @Success 200 {object} util.APIResult{data=RolesResp}
// @Security Token
// @Security Sub
// @Router /v1/auth/roles [get]
func (a *AuthAPI) Roles(c *web.Context) {
	req := &RolesReq{}
	if c.Bind(req) != nil {
		return
	}

	cond := &entity.FindRoleCond{
		Page: req.Page,
		Code: req.Code,
		Name: req.Name,
	}
	roles, err := a.authService.FindRoles(c, cond)
	if err != nil {
		c.Fail(err)
		return
	}
	resp := &RolesResp{
		Total: cond.Total,
		Roles: newRoleDTOs(roles),
	}
	c.Success(resp)
}
