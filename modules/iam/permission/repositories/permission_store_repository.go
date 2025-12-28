package repositories

import (
	"iam-service/infrastructure/databases/maindb"
	"iam-service/infrastructure/utils"
	iamModel "iam-service/models/maindb/iam"

	"gorm.io/gorm"
)

func Create(permission *iamModel.Permission, transaction *gorm.DB) *iamModel.Permission {
	connection := utils.If(transaction != nil, transaction, maindb.Connection)

	res := connection.Create(&permission)

	err := res.Error
	if err != nil {
		panic(err)
	}

	return permission
}

func Update(permission *iamModel.Permission, transaction *gorm.DB) *iamModel.Permission {
	connection := utils.If(transaction != nil, transaction, maindb.Connection)

	res := connection.Save(&permission)
	err := res.Error
	if err != nil {
		panic(err)
	}

	return permission
}

func Delete(permission *iamModel.Permission, transaction *gorm.DB) {
	connection := utils.If(transaction != nil, transaction, maindb.Connection)

	res := connection.Delete(&permission)

	err := res.Error
	if err != nil {
		panic(err)
	}
}
