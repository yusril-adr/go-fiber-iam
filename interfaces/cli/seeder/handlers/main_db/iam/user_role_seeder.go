package iam

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"

	"iam-service/infrastructure/databases/maindb"
	"iam-service/infrastructure/utils"
	"iam-service/models/maindb/iam"
)

type userRoleSeed struct {
	UserEmail string `json:"user_email"`
	RoleKey   string `json:"role_key"`
}

/*
- Read json from file
- Resolve user_email -> user_id
- Resolve role_key -> role_id
- Find non existing user_roles
- Insert non existing data
*/
func UserRoleSeeder() {
	path := filepath.Join("seeders", "main_db", "user_roles.json")

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	var seeds []userRoleSeed
	if err := json.Unmarshal(file, &seeds); err != nil {
		panic(err.Error())
	}

	// collect identifiers
	emails := utils.Map(seeds, func(s userRoleSeed) string {
		return s.UserEmail
	})

	roleKeys := utils.Map(seeds, func(s userRoleSeed) string {
		return s.RoleKey
	})

	maindb.InitConnection()

	// load users
	var users []iam.User
	maindb.Connection.Where("email IN ?", emails).Find(&users)

	// load roles
	var roles []iam.Role
	maindb.Connection.Where("key IN ?", roleKeys).Find(&roles)

	// build user_roles candidates
	var userRoles []iam.UserRole

	for _, seed := range seeds {
		user, userFound := utils.Find(users, func(u iam.User) bool {
			return u.Email == seed.UserEmail
		})

		role, roleFound := utils.Find(roles, func(r iam.Role) bool {
			return r.Key == seed.RoleKey
		})

		if !userFound || !roleFound {
			log.Printf(
				"Skipping user_role (user=%s, role=%s) — key not found",
				seed.UserEmail,
				seed.RoleKey,
			)
			continue
		}

		userRoles = append(userRoles, iam.UserRole{
			UserId: user.Id,
			RoleId: role.Id,
		})
	}

	if len(userRoles) == 0 {
		log.Println("No user roles to seed")
		return
	}

	// find existing user_roles
	userIds := utils.Map(userRoles, func(ur iam.UserRole) uuid.UUID {
		return ur.UserId
	})

	roleIds := utils.Map(userRoles, func(ur iam.UserRole) uuid.UUID {
		return ur.RoleId
	})

	var existing []iam.UserRole
	maindb.Connection.
		Where("user_id IN ? AND role_id IN ?", userIds, roleIds).
		Find(&existing)

	// filter non existing
	nonExisting := utils.Filter(userRoles, func(ur iam.UserRole) bool {
		_, found := utils.Find(existing, func(e iam.UserRole) bool {
			return e.UserId == ur.UserId && e.RoleId == ur.RoleId
		})
		return !found
	})

	if len(nonExisting) > 0 {
		maindb.Connection.Create(&nonExisting)
	}

	log.Println("User roles seeded successfully")
}
