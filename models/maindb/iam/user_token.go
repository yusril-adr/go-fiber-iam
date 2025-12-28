package iam

import (
	"time"

	"github.com/google/uuid"

	models "iam-service/models/maindb"
)

type UserToken struct {
	models.Base
	UserId    uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Token     string    `gorm:"size:255" json:"token"`
	ExpiresAt time.Time `gorm:"type:timestamp" json:"expires_at"`

	User User `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE" json:"user"`
}
