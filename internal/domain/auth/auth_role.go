package auth

import (
	"nfgo-ddd-showcase/internal/infra/util"
)

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
