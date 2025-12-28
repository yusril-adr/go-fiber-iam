package notifications

import (
	models "iam-service/models/maindb"
	"iam-service/models/maindb/iam"
)

type Notification struct {
	models.Base
	Title   string `gorm:"size:255;not null" json:"title"`
	Content string `gorm:"type:text;not null" json:"content"`
	IsRead  bool   `gorm:"default:false" json:"is_read"`
	Href    string `gorm:"type:text;default:null" json:"href"`
	UserId  string `gorm:"type:uuid;not null" json:"user_id"`

	// Relations
	User iam.User `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE" json:"user"`
}
