package params

import "github.com/google/uuid"

type RoleCreate struct {
	Name          string      `json:"name" validate:"required"`
	PermissionIds []uuid.UUID `json:"permission_ids" validate:"min=1,dive,uuid"`
}
