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

type rolePermissionSeed struct {
	RoleKey       string `json:"role_key"`
	PermissionKey string `json:"permission_key"`
}

/*
- Read json from file
- Resolve role_key -> role_id
- Resolve permission_key -> permission_id
- Find non existing role_permissions
- Insert non existing data
*/
func RolePermissionSeeder() {
	path := filepath.Join("seeders", "main_db", "role_permissions.json")

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	var seeds []rolePermissionSeed
	if err := json.Unmarshal(file, &seeds); err != nil {
		panic(err.Error())
	}

	// collect keys
	roleKeys := utils.Map(seeds, func(s rolePermissionSeed) string {
		return s.RoleKey
	})

	permissionKeys := utils.Map(seeds, func(s rolePermissionSeed) string {
		return s.PermissionKey
	})

	maindb.InitConnection()

	// load roles
	var roles []iam.Role
	maindb.Connection.Where("key IN ?", roleKeys).Find(&roles)

	// load permissions
	var permissions []iam.Permission
	maindb.Connection.Where("key IN ?", permissionKeys).Find(&permissions)

	// build role_permissions candidates
	var rolePermissions []iam.RolePermission

	for _, seed := range seeds {
		role, roleFound := utils.Find(roles, func(r iam.Role) bool {
			return r.Key == seed.RoleKey
		})

		permission, permFound := utils.Find(permissions, func(p iam.Permission) bool {
			return p.Key == seed.PermissionKey
		})

		if !roleFound || !permFound {
			log.Printf(
				"Skipping role_permission (role=%s, permission=%s) — key not found",
				seed.RoleKey,
				seed.PermissionKey,
			)
			continue
		}

		rolePermissions = append(rolePermissions, iam.RolePermission{
			RoleId:       role.Id,
			PermissionId: permission.Id,
		})
	}

	if len(rolePermissions) == 0 {
		log.Println("No role permissions to seed")
		return
	}

	// find existing role_permissions
	roleIds := utils.Map(rolePermissions, func(rp iam.RolePermission) uuid.UUID {
		return rp.RoleId
	})

	permissionIds := utils.Map(rolePermissions, func(rp iam.RolePermission) uuid.UUID {
		return rp.PermissionId
	})

	var existing []iam.RolePermission
	maindb.Connection.
		Where("role_id IN ? AND permission_id IN ?", roleIds, permissionIds).
		Find(&existing)

	// filter non existing
	nonExisting := utils.Filter(rolePermissions, func(rp iam.RolePermission) bool {
		_, found := utils.Find(existing, func(e iam.RolePermission) bool {
			return e.RoleId == rp.RoleId && e.PermissionId == rp.PermissionId
		})
		return !found
	})

	if len(nonExisting) > 0 {
		maindb.Connection.Create(&nonExisting)
	}

	log.Println("Role permissions seeded successfully")
}
