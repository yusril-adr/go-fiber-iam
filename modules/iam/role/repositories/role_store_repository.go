package repositories

import (
	"iam-service/infrastructure/databases/maindb"
	"iam-service/infrastructure/utils"
	iamModel "iam-service/models/maindb/iam"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Create(role *iamModel.Role, transaction *gorm.DB) *iamModel.Role {
	connection := utils.If(transaction != nil, transaction, maindb.Connection)

	res := connection.Create(&role)

	err := res.Error
	if err != nil {
		panic(err)
	}

	return role
}

func Update(role *iamModel.Role, transaction *gorm.DB) *iamModel.Role {
	connection := utils.If(transaction != nil, transaction, maindb.Connection)

	// Clear relations
	role.Permissions = nil

	res := connection.Save(&role)
	err := res.Error
	if err != nil {
		panic(err)
	}

	return role
}

func UpdateRolePermission(role *iamModel.Role, permissions []iamModel.Permission, transaction *gorm.DB) {
	connection := utils.If(transaction != nil, transaction, maindb.Connection)
	utils.LogJson(role, logrus.InfoLevel)
	utils.LogJson(permissions, logrus.InfoLevel)

	// Clear relations
	connection.Model(role).Association("Permissions").Clear()
	err := connection.Model(role).Association("Permissions").Replace(permissions)
	if err != nil {
		panic(err)
	}
}

func Delete(role *iamModel.Role, transaction *gorm.DB) {
	connection := utils.If(transaction != nil, transaction, maindb.Connection)

	res := connection.Delete(&role)

	err := res.Error
	if err != nil {
		panic(err)
	}
}
