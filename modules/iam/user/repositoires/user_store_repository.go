package repositories

import (
	"iam-service/infrastructure/databases/maindb"
	"iam-service/infrastructure/utils"
	iamModel "iam-service/models/maindb/iam"

	"gorm.io/gorm"
)

func Create(user *iamModel.User, transaction *gorm.DB) *iamModel.User {
	connection := utils.If(transaction != nil, transaction, maindb.Connection)

	res := connection.Create(&user)

	err := res.Error
	if err != nil {
		panic(err)
	}

	return user
}

func Update(user *iamModel.User, transaction *gorm.DB) *iamModel.User {
	connection := utils.If(transaction != nil, transaction, maindb.Connection)

	// Clear relations
	user.Roles = nil

	res := connection.Save(&user)
	err := res.Error
	if err != nil {
		panic(err)
	}

	return user
}

func UpdateUserRole(user *iamModel.User, roles []iamModel.Role, transaction *gorm.DB) {
	connection := utils.If(transaction != nil, transaction, maindb.Connection)

	// Clear relations
	connection.Model(user).Association("Roles").Clear()
	err := connection.Model(user).Association("Roles").Replace(roles)
	if err != nil {
		panic(err)
	}
}

func Delete(user *iamModel.User, transaction *gorm.DB) {
	connection := utils.If(transaction != nil, transaction, maindb.Connection)

	res := connection.Delete(&user)

	err := res.Error
	if err != nil {
		panic(err)
	}
}
