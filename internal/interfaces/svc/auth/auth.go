//go:generate protoc -I . -I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.5.1 --go_out=. --go-grpc_out=. --validate_out=lang=go:. auth.proto
package auth

import (
	"context"
	"fmt"
	"nfgo-ddd-showcase/internal/domain/auth"
	"nfgo-ddd-showcase/internal/infra/util"
)

type AuthSvc struct {
	authService auth.AuthService
	UnimplementedAuthSvcServer
}

func NewAuthSvc(authService auth.AuthService) *AuthSvc {
	return &AuthSvc{
		authService: authService,
	}
}

func (s *AuthSvc) Login(ctx context.Context, req *LoginReq) (*LoginResp, error) {

	user, err := s.authService.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	resp := &LoginResp{
		Token:   user.Token,
		SignKey: user.SignKey,
		Sub:     fmt.Sprint(user.ID),
	}
	return resp, nil
}

func (s *AuthSvc) Register(ctx context.Context, req *RegisterReq) (*ReisterResp, error) {
	user := &auth.AuthUser{
		Username: req.Username,
		Password: req.Password,
	}

	if err := s.authService.Register(ctx, user); err != nil {
		return nil, err
	}

	return &ReisterResp{}, nil
}

func (s *AuthSvc) UploadAvatar(ctx context.Context, req *UploadAvatarReq) (*UploadAvatarResp, error) {

	avatarURL, err := s.authService.UploadAvatar(ctx, req.UserID, req.File)

	return &UploadAvatarResp{
		AvatarURL: avatarURL,
	}, err
}

func (s *AuthSvc) FindRoles(ctx context.Context, req *FindRolesReq) (*FindRolesResp, error) {
	cond := &auth.FindRoleCond{
		Page: util.NewPage(req.PageNo, req.PageSize),
		Code: req.Code,
		Name: req.Name,
	}

	roles, err := s.authService.FindRoles(ctx, cond)
	if err != nil {
		return nil, err
	}

	resp := &FindRolesResp{
		Total: cond.Total,
		Roles: newRoleDTOs(roles),
	}

	return resp, nil
}
