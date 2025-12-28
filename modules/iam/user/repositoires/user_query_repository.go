package repositories

import (
	"errors"
	"iam-service/infrastructure/databases/maindb"
	appError "iam-service/infrastructure/errors"
	"iam-service/infrastructure/messages"
	"iam-service/infrastructure/types"
	"iam-service/infrastructure/utils"
	iamModel "iam-service/models/maindb/iam"
	"iam-service/modules/iam/user/dtos/params"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Paginate(pagination *params.UserPagination) ([]iamModel.User, *types.TMetaResult) {
	// Build query
	query := maindb.Connection.
		Model(&iamModel.User{}).
		Preload("Roles")

	paginateFilter(query, pagination)
	paginateSearch(query, pagination)
	paginateSort(query, pagination)
	userCount := paginateQuery(query, pagination)

	var users []iamModel.User
	res := query.Find(&users)

	err := res.Error
	if err != nil {
		// Check if record not found
		// Then return nil
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []iamModel.User{}, nil
		}

		panic(err)
	}

	meta := utils.PaginationMetaBuilder(*pagination.Page, *pagination.PerPage, userCount)
	return users, meta
}

func paginateFilter(query *gorm.DB, pagination *params.UserPagination) {
	if pagination.RoleIds != nil && len(*pagination.RoleIds) > 0 {
		query = query.Joins(`
			JOIN user_roles ur ON ur.user_id = users.id
		`).Where(`
			ur.role_id IN ?
		`, *pagination.RoleIds)
	}
}

func paginateSearch(query *gorm.DB, pagination *params.UserPagination) {
	if pagination.Search != nil {
		query = query.Where("name LIKE ? OR email LIKE ?", "%"+*pagination.Search+"%", "%"+*pagination.Search+"%")
	}
}

func paginateSort(query *gorm.DB, pagination *params.UserPagination) {
	if pagination.SortBy != nil {
		var allowedSort = map[string]string{
			"name":       "users.name",
			"email":      "users.email",
			"updated_at": "users.updated_at",
			"created_at": "users.created_at",
		}
		allowedFields := make([]string, 0, len(allowedSort))
		for k := range allowedSort {
			allowedFields = append(allowedFields, k)
		}

		if col, ok := allowedSort[*pagination.SortBy]; ok {
			query = query.Order(col + " " + *pagination.Order)
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
		query = query.Order(*pagination.Order)
	}
}

func paginateQuery(query *gorm.DB, pagination *params.UserPagination) int64 {
	var userCount int64
	query.Count(&userCount)

	query.Offset((*pagination.Page - 1) * *pagination.PerPage)
	query.Limit(*pagination.PerPage)

	return userCount
}

func FindByEmail(email string) *iamModel.User {
	var user *iamModel.User
	res := maindb.Connection.Preload("Roles").Where("email = ?", email).First(&user)

	err := res.Error
	if err != nil {

		// Check if record not found
		// Then return nil
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		panic(err)
	}

	return user
}

func FindById(id uuid.UUID) *iamModel.User {
	var user *iamModel.User
	res := maindb.Connection.Preload("Roles").Where("id = ?", id).First(&user)

	err := res.Error
	if err != nil {

		// Check if record not found
		// Then return nil
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		panic(err)
	}

	return user
}
