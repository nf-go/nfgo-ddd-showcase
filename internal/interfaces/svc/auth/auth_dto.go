package auth

import "nfgo-ddd-showcase/internal/domain/auth/entity"

func newRoleDTO(role *entity.AuthRole) *RoleDTO {
	return &RoleDTO{
		Id:          role.ID,
		Name:        role.Name,
		Code:        role.Code,
		Description: role.Description,
	}
}

func newRoleDTOs(roles []*entity.AuthRole) []*RoleDTO {
	dtos := make([]*RoleDTO, 0, len(roles))
	for _, role := range roles {
		dtos = append(dtos, newRoleDTO(role))
	}
	return dtos
}
