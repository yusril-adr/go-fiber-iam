package repositories

import (
	"errors"
	"iam-service/infrastructure/databases/maindb"
	iamModel "iam-service/models/maindb/iam"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserToken(userId uuid.UUID, token string) *iamModel.UserToken {
	var userToken *iamModel.UserToken
	res := maindb.
		Connection.
		Where("user_id = ? AND token = ? AND expires_at > ?", userId, token, time.Now()).
		First(&userToken)

	err := res.Error
	if err != nil {
		// Check if record not found
		// Then return nil
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		panic(err)
	}

	return userToken
}
