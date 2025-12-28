package results

import (
	"github.com/google/uuid"

	"iam-service/infrastructure/utils"
	iamModel "iam-service/models/maindb/iam"
	permissionResult "iam-service/modules/iam/permission/dtos/results"
)

type RoleResult struct {
	Id          uuid.UUID                            `json:"id"`
	Name        string                               `json:"name"`
	Key         string                               `json:"key"`
	Permissions []*permissionResult.PermissionResult `json:"permissions"`
}

func (r *RoleResult) MapModel(role iamModel.Role) {
	r.Id = role.Id
	r.Name = role.Name
	r.Key = role.Key

	r.Permissions = utils.Map(role.Permissions, func(permission iamModel.Permission) *permissionResult.PermissionResult {
		result := &permissionResult.PermissionResult{}
		result.MapModel(permission)
		return result
	})
}
