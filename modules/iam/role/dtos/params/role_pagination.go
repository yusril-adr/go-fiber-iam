package params

import (
	"iam-service/infrastructure/dtos/params"

	"github.com/google/uuid"
)

type RolePagination struct {
	params.Pagination
	PermissionIds *[]uuid.UUID `query:"permission_ids"`
}

func (p *RolePagination) SetDefaultValue() {
	p.Pagination.SetDefaultValue()
}
