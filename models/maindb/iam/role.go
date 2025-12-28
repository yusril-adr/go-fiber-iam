package iam

import (
	models "iam-service/models/maindb"
)

type Role struct {
	models.Base
	Name string `gorm:"size:255;not null" json:"name"`
	Key  string `gorm:"size:255;not null" json:"key"`

	// Relations
	Permissions []Permission `gorm:"many2many:role_permissions;foreignKey:ID;joinForeignKey:role_id;References:ID;joinReferences:permission_id" json:"permissions,omitempty"`
}
