package repositories

import (
	"iam-service/infrastructure/databases/maindb"
	"iam-service/infrastructure/utils"
	iamModel "iam-service/models/maindb/iam"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Create(userToken *iamModel.UserToken, transaction *gorm.DB) {
	connection := utils.If(transaction != nil, transaction, maindb.Connection)
	res := connection.Create(&userToken)

	err := res.Error
	if err != nil {
		panic(err)
	}
}

func DeleteTokenWithUserId(token string, userId uuid.UUID, transaction *gorm.DB) {
	connection := utils.If(transaction != nil, transaction, maindb.Connection)
	res := connection.
		Where("user_id = ? AND token = ?", userId, token).
		Delete(&iamModel.UserToken{})

	err := res.Error
	if err != nil {
		panic(err)
	}
}

func DeleteTokenWithExpiredAt(expiredAt time.Time, transaction *gorm.DB) {
	connection := utils.If(transaction != nil, transaction, maindb.Connection)
	res := connection.
		Unscoped().
		Where("expires_at <= ?", expiredAt).
		Delete(&iamModel.UserToken{})

	err := res.Error
	if err != nil {
		panic(err)
	}
}
