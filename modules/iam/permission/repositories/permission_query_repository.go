package repositories

import (
	"errors"

	"iam-service/infrastructure/databases/maindb"
	appError "iam-service/infrastructure/errors"
	"iam-service/infrastructure/messages"
	"iam-service/infrastructure/types"
	"iam-service/infrastructure/utils"
	iamModel "iam-service/models/maindb/iam"
	"iam-service/modules/iam/permission/dtos/params"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func FindManyByIds(ids []uuid.UUID) []iamModel.Permission {
	var permissions []iamModel.Permission

	res := maindb.Connection.Where("id IN ?", ids).Find(&permissions)

	err := res.Error
	if err != nil {

		// Check if record not found
		// Then return nil
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []iamModel.Permission{}
		}

		panic(err)
	}

	return permissions
}

func FindById(id uuid.UUID) *iamModel.Permission {
	var permission *iamModel.Permission
	res := maindb.Connection.Where("id = ?", id).First(&permission)

	err := res.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(err)
	}

	return permission
}

func FindByKey(key string) *iamModel.Permission {
	var permission *iamModel.Permission
	res := maindb.Connection.Where("key = ?", key).First(&permission)

	err := res.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(err)
	}

	return permission
}

func FindByName(name string) *iamModel.Permission {
	var permission *iamModel.Permission
	res := maindb.Connection.Where("name = ?", name).First(&permission)

	err := res.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(err)
	}

	return permission
}

func Paginate(pagination *params.PermissionPagination) ([]iamModel.Permission, *types.TMetaResult) {
	// Build query
	query := maindb.Connection.Model(&iamModel.Permission{})

	paginateSearch(query, pagination)
	paginateSort(query, pagination)
	permissionCount := paginateQuery(query, pagination)

	var permissions []iamModel.Permission
	res := query.Find(&permissions)

	err := res.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []iamModel.Permission{}, nil
		}
		panic(err)
	}

	meta := utils.PaginationMetaBuilder(*pagination.Page, *pagination.PerPage, permissionCount)
	return permissions, meta
}

func paginateSearch(query *gorm.DB, pagination *params.PermissionPagination) {
	if pagination.Search != nil {
		query.Where("name LIKE ? OR key LIKE ?", "%"+*pagination.Search+"%", "%"+*pagination.Search+"%")
	}
}

func paginateSort(query *gorm.DB, pagination *params.PermissionPagination) {
	if pagination.SortBy != nil {
		var allowedSort = map[string]string{
			"name":       "permissions.name",
			"key":        "permissions.key",
			"updated_at": "permissions.updated_at",
			"created_at": "permissions.created_at",
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

func paginateQuery(query *gorm.DB, pagination *params.PermissionPagination) int64 {
	var permissionCount int64
	query.Count(&permissionCount)

	query.Offset((*pagination.Page - 1) * *pagination.PerPage)
	query.Limit(*pagination.PerPage)

	return permissionCount
}

func CountRolesByPermissionId(permissionId uuid.UUID) int64 {
	var count int64
	maindb.Connection.
		Table("role_permissions").
		Where("permission_id = ?", permissionId).
		Count(&count)
	return count
}
