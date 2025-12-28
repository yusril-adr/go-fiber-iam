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
func RoleSeeder() {
	path := filepath.Join("seeders", "main_db", "roles.json")

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	var roles []iam.Role
	if err := json.Unmarshal(file, &roles); err != nil {
		panic(err.Error())
	}

	jsonRoleKey := utils.Map(roles, func(role iam.Role) string { return role.Key })

	maindb.InitConnection()

	var existingRoles []iam.Role
	maindb.Connection.Where("key In ?", jsonRoleKey).Find(&existingRoles)

	nonExistingRoles := utils.Filter(roles, func(role iam.Role) bool {
		_, isRoleFound := utils.Find(existingRoles, func(r iam.Role) bool { return r.Key == role.Key })
		return !isRoleFound
	})

	if len(nonExistingRoles) > 0 {
		maindb.Connection.Create(&nonExistingRoles)
	}

	log.Println("Roles seeded successfully")
}
