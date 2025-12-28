package maindb

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	Id        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime:true index" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:true" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
