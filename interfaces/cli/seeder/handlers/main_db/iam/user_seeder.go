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
- Hash password
- Insert non existing data
*/
func UserSeeder() {
	path := filepath.Join("seeders", "main_db", "users.json")

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	var users []iam.User
	if err := json.Unmarshal(file, &users); err != nil {
		panic(err.Error())
	}

	jsonUserEmail := utils.Map(users, func(user iam.User) string { return user.Email })

	maindb.InitConnection()
	var existingUsers []iam.User
	maindb.Connection.Where("email In ?", jsonUserEmail).Find(&existingUsers)

	nonExistingUsers := utils.Filter(users, func(user iam.User) bool {
		_, isUserFound := utils.Find(existingUsers, func(u iam.User) bool { return u.Email == user.Email })
		return !isUserFound
	})

	if len(nonExistingUsers) > 0 {
		for i := range nonExistingUsers {
			users[i].Password = utils.HashPassword(users[i].Password)
		}

		maindb.Connection.Create(&nonExistingUsers)
	}

	log.Println("Users seeded successfully")
}
