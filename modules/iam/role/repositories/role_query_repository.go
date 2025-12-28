package repositories

import (
	"errors"

	"iam-service/infrastructure/databases/maindb"
	appError "iam-service/infrastructure/errors"
	"iam-service/infrastructure/messages"
	"iam-service/infrastructure/types"
	"iam-service/infrastructure/utils"
	iamModel "iam-service/models/maindb/iam"
	"iam-service/modules/iam/role/dtos/params"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func FindManyByIds(ids []uuid.UUID) []iamModel.Role {
	var roles []iamModel.Role

	res := maindb.Connection.Where("id IN ?", ids).Find(&roles)

	err := res.Error
	if err != nil {

		// Check if record not found
		// Then return nil
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []iamModel.Role{}
		}

		panic(err)
	}

	return roles
}

func Paginate(pagination *params.RolePagination) ([]iamModel.Role, *types.TMetaResult) {
	// Build query
	query := maindb.Connection.
		Model(&iamModel.Role{}).
		Preload("Permissions")

	paginateFilter(query, pagination)
	paginateSearch(query, pagination)
	paginateSort(query, pagination)
	roleCount := paginateQuery(query, pagination)

	var roles []iamModel.Role
	res := query.Find(&roles)

	err := res.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []iamModel.Role{}, nil
		}
		panic(err)
	}

	meta := utils.PaginationMetaBuilder(*pagination.Page, *pagination.PerPage, roleCount)
	return roles, meta
}

func paginateFilter(query *gorm.DB, pagination *params.RolePagination) {
	if pagination.PermissionIds != nil && len(*pagination.PermissionIds) > 0 {
		query.Joins("JOIN role_permissions ON role_permissions.role_id = roles.id").
			Where("role_permissions.permission_id IN ?", *pagination.PermissionIds).
			Distinct()
	}
}

func paginateSearch(query *gorm.DB, pagination *params.RolePagination) {
	if pagination.Search != nil {
		query.Where("name LIKE ? OR key LIKE ?", "%"+*pagination.Search+"%", "%"+*pagination.Search+"%")
	}
}

func paginateSort(query *gorm.DB, pagination *params.RolePagination) {
	if pagination.SortBy != nil {
		var allowedSort = map[string]string{
			"name":       "roles.name",
			"key":        "roles.key",
			"updated_at": "roles.updated_at",
			"created_at": "roles.created_at",
		}
		allowedFields := make([]string, 0, len(allowedSort))
		for k := range allowedSort {
			allowedFields = append(allowedFields, k)
		}

		if col, ok := allowedSort[*pagination.SortBy]; ok {
			query.Order(col + " " + *pagination.Order)
			return
		} else {
			panic(
				appError.
					ValidationError(
						"sort_by",
						messages.ErrSortByColInvalid(*pagination.SortBy, allowedFields),
					),
			)
		}
	}

	if pagination.Order != nil {
		query.Order(*pagination.Order)
	}
}

func paginateQuery(query *gorm.DB, pagination *params.RolePagination) int64 {
	var roleCount int64
	query.Count(&roleCount)

	query.Offset((*pagination.Page - 1) * *pagination.PerPage)
	query.Limit(*pagination.PerPage)

	return roleCount
}

func FindByName(name string) *iamModel.Role {
	var role *iamModel.Role
	res := maindb.Connection.Preload("Permissions").Where("name = ?", name).First(&role)

	err := res.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(err)
	}

	return role
}

func FindByKey(key string) *iamModel.Role {
	var role *iamModel.Role
	res := maindb.Connection.Preload("Permissions").Where("key = ?", key).First(&role)

	err := res.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(err)
	}

	return role
}

func FindById(id uuid.UUID) *iamModel.Role {
	var role *iamModel.Role
	res := maindb.Connection.Preload("Permissions").Where("id = ?", id).First(&role)

	err := res.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(err)
	}

	return role
}

func CountUsersByRoleId(roleId uuid.UUID) int64 {
	var count int64
	maindb.Connection.
		Table("user_roles").
		Where("role_id = ?", roleId).
		Count(&count)
	return count
}
