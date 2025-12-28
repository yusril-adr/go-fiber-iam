package iam

import models "iam-service/models/maindb"

type Permission struct {
	models.Base
	Name string `gorm:"size:255;not null" json:"name"`
	Key  string `gorm:"size:255;not null" json:"key"`
}
