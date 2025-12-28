package services

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"iam-service/infrastructure/databases/maindb"
	"iam-service/infrastructure/errors"
	"iam-service/infrastructure/types"
	"iam-service/infrastructure/utils"
	iamModel "iam-service/models/maindb/iam"
	permissionService "iam-service/modules/iam/permission/services"
	"iam-service/modules/iam/role/dtos/params"
	"iam-service/modules/iam/role/dtos/results"
	"iam-service/modules/iam/role/messages"
	roleRepository "iam-service/modules/iam/role/repositories"
)

func FindAndValidateExistingRoleIds(roleIds []uuid.UUID) []iamModel.Role {
	roles := roleRepository.FindManyByIds(roleIds)

	notFoundRoleIds := utils.Filter(roleIds, func(roleId uuid.UUID) bool {
		_, isRoleFound := utils.Find(roles, func(role iamModel.Role) bool { return role.Id == roleId })
		return !isRoleFound
	})

	if len(notFoundRoleIds) > 0 {
		panic(errors.NotFound(messages.ErrRolesNotFound(notFoundRoleIds)))
	}

	return roles
}

// generateUniqueRoleKey generates a unique key from a role name
// If the base key already exists, it appends a random 4-character suffix
func generateUniqueRoleKey(baseName string, excludeId *uuid.UUID) string {
	// Generate base key from name
	baseKey := utils.GenerateSlug(baseName)

	// Check if base key is available
	existingRole := roleRepository.FindByKey(baseKey)
	if existingRole == nil || (excludeId != nil && existingRole.Id == *excludeId) {
		return baseKey
	}

	// Key collision - append random suffix and retry recursively
	return generateUniqueRoleKeyWithSuffix(baseKey, excludeId)
}

// generateUniqueRoleKeyWithSuffix recursively tries to find a unique key with random suffix
func generateUniqueRoleKeyWithSuffix(baseKey string, excludeId *uuid.UUID) string {
	// Generate random 4-character suffix
	suffix := utils.GenerateRandomSuffix(4)
	candidateKey := fmt.Sprintf("%s_%s", baseKey, suffix)

	// Check if this key is available
	existingRole := roleRepository.FindByKey(candidateKey)
	if existingRole == nil || (excludeId != nil && existingRole.Id == *excludeId) {
		return candidateKey
	}

	// Still collision (very rare) - recurse with new random suffix
	return generateUniqueRoleKeyWithSuffix(baseKey, excludeId)
}

func CreateRole(dto *params.RoleCreate) results.RoleResult {
	permissions := permissionService.FindAndValidateExistingPermissionIds(dto.PermissionIds)

	// Auto-generate unique key from name with collision handling
	generatedKey := generateUniqueRoleKey(dto.Name, nil)

	role := iamModel.Role{
		Name: dto.Name,
		Key:  generatedKey,
	}

	transaction := maindb.Connection.Begin()
	savedRoleId := func() uuid.UUID {
		defer func(tx *gorm.DB) {
			if rec := recover(); rec != nil {
				tx.Rollback()
				panic(rec)
			}
		}(transaction)

		savedRole := roleRepository.Create(&role, transaction)
		roleRepository.UpdateRolePermission(savedRole, permissions, transaction)

		transaction.Commit()

		return savedRole.Id
	}()

	return GetRoleById(savedRoleId)
}

func PaginateRoles(pagination *params.RolePagination) *types.TPaginationResult[results.RoleResult] {
	rolePaginations, metaPagination := roleRepository.Paginate(pagination)

	mappedRoles := utils.Map(rolePaginations, func(rolePagination iamModel.Role) results.RoleResult {
		result := results.RoleResult{}
		result.MapModel(rolePagination)
		return result
	})
	return utils.PaginationBuilder(mappedRoles, *metaPagination)
}

func GetRoleById(id uuid.UUID) results.RoleResult {
	role := roleRepository.FindById(id)

	if role == nil {
		panic(errors.NotFound(messages.ErrRoleNotFound))
	}

	result := results.RoleResult{}
	result.MapModel(*role)
	return result
}

func UpdateRole(id uuid.UUID, dto *params.RoleUpdate) results.RoleResult {
	role := roleRepository.FindById(id)
	if role == nil {
		panic(errors.NotFound(messages.ErrRoleNotFound))
	}

	permissions := permissionService.FindAndValidateExistingPermissionIds(dto.PermissionIds)

	// Update role name only - key stays unchanged (immutable)
	role.Name = dto.Name
	// role.Key stays as-is (NOT regenerated)

	transaction := maindb.Connection.Begin()

	func() {
		defer func(tx *gorm.DB) {
			if rec := recover(); rec != nil {
				tx.Rollback()
				panic(rec)
			}
		}(transaction)

		roleRepository.Update(role, transaction)
		roleRepository.UpdateRolePermission(role, permissions, transaction)

		transaction.Commit()
	}()

	return GetRoleById(id)
}

func DeleteRole(id uuid.UUID) {
	role := roleRepository.FindById(id)
	if role == nil {
		panic(errors.NotFound(messages.ErrRoleNotFound))
	}

	// Check if role is assigned to any users
	userCount := roleRepository.CountUsersByRoleId(id)
	if userCount > 0 {
		panic(errors.ValidationError("role_id", messages.ErrRoleAssignedToUsers))
	}

	// Delete role
	roleRepository.Delete(role, nil)
}
