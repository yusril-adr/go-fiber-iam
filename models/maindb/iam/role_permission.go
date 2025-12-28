package iam

import (
	"github.com/google/uuid"

	models "iam-service/models/maindb"
)

type RolePermission struct {
	models.Base
	RoleId       uuid.UUID  `gorm:"type:uuid" json:"role_id"`
	PermissionId uuid.UUID  `gorm:"type:uuid" json:"permission_id"`
	Role         Role       `gorm:"foreignKey:RoleId;constraint:OnDelete:CASCADE" json:"role"`
	Permission   Permission `gorm:"foreignKey:permissionId;constraint:OnDelete:CASCADE" json:"permission"`
}
