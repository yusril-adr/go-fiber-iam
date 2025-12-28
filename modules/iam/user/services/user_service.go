package services

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"iam-service/infrastructure/databases/maindb"
	"iam-service/infrastructure/errors"
	"iam-service/infrastructure/types"
	"iam-service/infrastructure/utils"
	iamModel "iam-service/models/maindb/iam"
	roleService "iam-service/modules/iam/role/services"
	"iam-service/modules/iam/user/dtos/params"
	"iam-service/modules/iam/user/dtos/results"
	"iam-service/modules/iam/user/messages"
	userRepository "iam-service/modules/iam/user/repositoires"
)

func CreateUser(dto *params.UserCreate) results.UserResult {
	roles := roleService.FindAndValidateExistingRoleIds(dto.RoleIds)

	// Validate email uniqueness
	existingUser := userRepository.FindByEmail(dto.Email)
	if existingUser != nil {
		panic(errors.ValidationError("email", messages.ErrUserEmailAlreadyExists))
	}

	// Hash password
	// var hashedPassword = utils.HashPassword(dto.Password)
	var hashedPassword = "dummy"
	user := iamModel.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: hashedPassword,
	}

	transaction := maindb.Connection.Begin()
	savedUserId := func() uuid.UUID {
		defer func(tx *gorm.DB) {
			if rec := recover(); rec != nil {
				tx.Rollback()
				panic(rec)
			}
		}(transaction)

		savedUser := userRepository.Create(&user, nil)
		userRepository.UpdateUserRole(savedUser, roles, transaction)

		transaction.Commit()

		return savedUser.Id
	}()

	return GetUserById(savedUserId)
}

func PaginateUsers(pagination *params.UserPagination) *types.TPaginationResult[results.UserResult] {
	userPaginations, metaPagination := userRepository.Paginate(pagination)

	mappedUsers := utils.Map(userPaginations, func(userPagination iamModel.User) results.UserResult {
		result := results.UserResult{}
		result.MapModel(userPagination)
		return result
	})
	return utils.PaginationBuilder(mappedUsers, *metaPagination)
}

func GetUserById(id uuid.UUID) results.UserResult {
	user := userRepository.FindById(id)

	if user == nil {
		panic(errors.NotFound(messages.ErrUserNotFound))
	}

	result := results.UserResult{}
	result.MapModel(*user)
	return result
}

func UpdateUser(id uuid.UUID, dto *params.UserUpdate) results.UserResult {
	user := userRepository.FindById(id)
	if user == nil {
		panic(errors.NotFound(messages.ErrUserNotFound))
	}

	existingUserEmail := userRepository.FindByEmail(dto.Email)
	if existingUserEmail != nil && existingUserEmail.Id != id {
		panic(errors.ValidationError("email", messages.ErrUserEmailAlreadyExists))
	}

	roles := roleService.FindAndValidateExistingRoleIds(dto.RoleIds)

	// Update user
	hashedPassword := utils.HashPassword(dto.Password)
	user.Name = dto.Name
	user.Email = dto.Email
	user.Password = hashedPassword

	transaction := maindb.Connection.Begin()

	func() {
		defer func(tx *gorm.DB) {
			if rec := recover(); rec != nil {
				tx.Rollback()
				panic(rec)
			}
		}(transaction)

		userRepository.Update(user, transaction)
		userRepository.UpdateUserRole(user, roles, transaction)

		transaction.Commit()
	}()

	return GetUserById(id)
}

func DeleteUser(id uuid.UUID) {
	user := userRepository.FindById(id)
	if user == nil {
		panic(errors.NotFound(messages.ErrUserNotFound))
	}

	// Delete user
	userRepository.Delete(user, nil)
}
