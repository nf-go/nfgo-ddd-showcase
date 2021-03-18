package entity

import "nfgo-ddd-showcase/internal/infra/util"

// ResourceType -
type ResourceType int8

const (
	// ResTypeMenu -
	ResTypeMenu ResourceType = 1
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
