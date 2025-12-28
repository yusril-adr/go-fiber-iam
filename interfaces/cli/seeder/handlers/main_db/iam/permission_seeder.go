package iam

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"iam-service/infrastructure/databases/maindb"
	"iam-service/infrastructure/utils"
	"iam-service/models/maindb/iam"
)

/*
- Read json from file
- Find non existing data
- Insert non existing data
*/
func PermissionSeeder() {
	path := filepath.Join("seeders", "main_db", "permissions.json")

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	var permissions []iam.Permission
	if err := json.Unmarshal(file, &permissions); err != nil {
		panic(err.Error())
	}

	jsonPermissionKey := utils.Map(permissions, func(permission iam.Permission) string { return permission.Key })

	maindb.InitConnection()

	var existingPermissions []iam.Permission
	maindb.Connection.Where("key In ?", jsonPermissionKey).Find(&existingPermissions)

	nonExistingPermissions := utils.Filter(permissions, func(permission iam.Permission) bool {
		_, isPermissionFound := utils.Find(existingPermissions, func(p iam.Permission) bool { return p.Key == permission.Key })
		return !isPermissionFound
	})

	if len(nonExistingPermissions) > 0 {
		maindb.Connection.Create(&nonExistingPermissions)
	}

	log.Println("Permissions seeded successfully")
}
