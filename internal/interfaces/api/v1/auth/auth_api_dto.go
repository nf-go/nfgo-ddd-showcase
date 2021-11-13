package auth

import (
	"nfgo-ddd-showcase/internal/domain/auth"
	"nfgo-ddd-showcase/internal/infra/util"
)

// UploadAvatarResp -
type UploadAvatarResp struct {
	AvatarURL string `json:"avatarURL,omitempty"`
}

// LoginReq -
type LoginReq struct {
	Username string `form:"username" binding:"required,min=5,max=10"`
	Password string `form:"password" binding:"required,min=5,max=10"`
}

// LoginResp -
type LoginResp struct {
	Token   string `json:"token,omitempty"`
	Sub     string `josn:"sub,omitempty"`
	SignKey string `json:"signKey,omitempty"`
}

// RegisterReq -
type RegisterReq struct {
	Username string `json:"username" binding:"required,min=5,max=10"`
	Password string `json:"password" binding:"required,min=5,max=10"`
}

// RegisterResp -
type RegisterResp struct {
}

// RoleDTO -
type RoleDTO struct {
	ID          int64
	Code        string
	Name        string
	Description string
}

func newRoleDTO(role *auth.AuthRole) *RoleDTO {
	return &RoleDTO{
		ID:          role.ID,
		Name:        role.Name,
		Code:        role.Code,
		Description: role.Description,
	}
}

func newRoleDTOs(roles []*auth.AuthRole) []*RoleDTO {
	dtos := make([]*RoleDTO, 0, len(roles))
	for _, role := range roles {
		dtos = append(dtos, newRoleDTO(role))
	}
	return dtos
}

// RolesReq -
type RolesReq struct {
	util.Page
	Code string `form:"code"`
	Name string `form:"name"`
}

// RolesResp -
type RolesResp struct {
	Total int64      `json:"total,omitempty"`
	Roles []*RoleDTO `json:"roles,omitempty"`
}
