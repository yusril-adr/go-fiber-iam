package iam

import (
	models "iam-service/models/maindb"
)

type User struct {
	models.Base
	Name     string `gorm:"size:255;not null" json:"name"`
	Email    string `gorm:"size:255;not null" json:"email"`
	Password string `gorm:"size:255;not null" json:"password"`

	// Relations
	Roles []Role `gorm:"many2many:user_roles;foreignKey:ID;joinForeignKey:user_id;References:ID;joinReferences:role_id" json:"roles"`
}
