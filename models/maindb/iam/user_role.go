package iam

import (
	"github.com/google/uuid"

	models "iam-service/models/maindb"
)

type UserRole struct {
	models.Base
	RoleId uuid.UUID `gorm:"type:uuid" json:"role_id"`
	UserId uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Role   Role      `gorm:"foreignKey:RoleId;constraint:OnDelete:CASCADE" json:"role"`
	User   User      `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE" json:"user"`
}
