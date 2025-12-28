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
	"iam-service/modules/iam/permission/dtos/params"
	"iam-service/modules/iam/permission/dtos/results"
	"iam-service/modules/iam/permission/messages"
	permissionRepository "iam-service/modules/iam/permission/repositories"
)

func FindAndValidateExistingPermissionIds(permissionIds []uuid.UUID) []iamModel.Permission {
	permissions := permissionRepository.FindManyByIds(permissionIds)

	notFoundPermissionIds := utils.Filter(permissionIds, func(permissionId uuid.UUID) bool {
		_, isPermissionFound := utils.Find(permissions, func(permission iamModel.Permission) bool {
			return permission.Id == permissionId
		})
		return !isPermissionFound
	})

	if len(notFoundPermissionIds) > 0 {
		panic(errors.NotFound(messages.ErrPermissionsNotFound(notFoundPermissionIds)))
	}

	return permissions
}

// generateUniquePermissionKey generates a unique key from a permission name
// If the base key already exists, it appends a random 4-character suffix
func generateUniquePermissionKey(baseName string, excludeId *uuid.UUID) string {
	// Generate base key from name
	baseKey := utils.GenerateSlug(baseName)

	// Check if base key is available
	existingPermission := permissionRepository.FindByKey(baseKey)
	if existingPermission == nil || (excludeId != nil && existingPermission.Id == *excludeId) {
		return baseKey
	}

	// Key collision - append random suffix and retry recursively
	return generateUniquePermissionKeyWithSuffix(baseKey, excludeId)
}

// generateUniquePermissionKeyWithSuffix recursively tries to find a unique key with random suffix
func generateUniquePermissionKeyWithSuffix(baseKey string, excludeId *uuid.UUID) string {
	// Generate random 4-character suffix
	suffix := utils.GenerateRandomSuffix(4)
	candidateKey := fmt.Sprintf("%s_%s", baseKey, suffix)

	// Check if this key is available
	existingPermission := permissionRepository.FindByKey(candidateKey)
	if existingPermission == nil || (excludeId != nil && existingPermission.Id == *excludeId) {
		return candidateKey
	}

	// Still collision (very rare) - recurse with new random suffix
	return generateUniquePermissionKeyWithSuffix(baseKey, excludeId)
}

func CreatePermission(dto *params.PermissionCreate) results.PermissionResult {
	// Auto-generate unique key from name with collision handling
	generatedKey := generateUniquePermissionKey(dto.Name, nil)

	permission := iamModel.Permission{
		Name: dto.Name,
		Key:  generatedKey,
	}

	transaction := maindb.Connection.Begin()
	savedPermissionId := func() uuid.UUID {
		defer func(tx *gorm.DB) {
			if rec := recover(); rec != nil {
				tx.Rollback()
				panic(rec)
			}
		}(transaction)

		savedPermission := permissionRepository.Create(&permission, transaction)

		transaction.Commit()

		return savedPermission.Id
	}()

	return GetPermissionById(savedPermissionId)
}

func PaginatePermissions(pagination *params.PermissionPagination) *types.TPaginationResult[results.PermissionResult] {
	permissionPaginations, metaPagination := permissionRepository.Paginate(pagination)

	mappedPermissions := utils.Map(permissionPaginations, func(permissionPagination iamModel.Permission) results.PermissionResult {
		result := results.PermissionResult{}
		result.MapModel(permissionPagination)
		return result
	})
	return utils.PaginationBuilder(mappedPermissions, *metaPagination)
}

func GetPermissionById(id uuid.UUID) results.PermissionResult {
	permission := permissionRepository.FindById(id)

	if permission == nil {
		panic(errors.NotFound(messages.ErrPermissionNotFound))
	}

	result := results.PermissionResult{}
	result.MapModel(*permission)
	return result
}

func UpdatePermission(id uuid.UUID, dto *params.PermissionUpdate) results.PermissionResult {
	permission := permissionRepository.FindById(id)
	if permission == nil {
		panic(errors.NotFound(messages.ErrPermissionNotFound))
	}

	// Update permission name only - key stays unchanged (immutable)
	permission.Name = dto.Name
	// permission.Key stays as-is (NOT regenerated)

	transaction := maindb.Connection.Begin()

	func() {
		defer func(tx *gorm.DB) {
			if rec := recover(); rec != nil {
				tx.Rollback()
				panic(rec)
			}
		}(transaction)

		permissionRepository.Update(permission, transaction)

		transaction.Commit()
	}()

	return GetPermissionById(id)
}

func DeletePermission(id uuid.UUID) {
	permission := permissionRepository.FindById(id)
	if permission == nil {
		panic(errors.NotFound(messages.ErrPermissionNotFound))
	}

	// Check if permission is assigned to any roles
	roleCount := permissionRepository.CountRolesByPermissionId(id)
	if roleCount > 0 {
		panic(errors.ValidationError("permission_id", messages.ErrPermissionAssignedToRoles))
	}

	// Delete permission
	permissionRepository.Delete(permission, nil)
}
