package results

import (
	"iam-service/infrastructure/utils"
	iamModel "iam-service/models/maindb/iam"
	"iam-service/modules/iam/role/dtos/results"

	"github.com/google/uuid"
)

type UserResult struct {
	Id    uuid.UUID             `json:"id"`
	Name  string                `json:"name"`
	Email string                `json:"email"`
	Roles []*results.RoleResult `json:"roles"`
}

func (u *UserResult) MapModel(user iamModel.User) {
	u.Id = user.Id
	u.Name = user.Name
	u.Email = user.Email

	u.Roles = utils.Map(user.Roles, func(role iamModel.Role) *results.RoleResult {
		result := &results.RoleResult{}
		result.MapModel(role)
		return result
	})
}
