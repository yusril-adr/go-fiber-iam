package params

import (
	"iam-service/infrastructure/dtos/params"

	"github.com/google/uuid"
)

type UserPagination struct {
	params.Pagination
	RoleIds *[]uuid.UUID `query:"role_ids" validate:"omitempty,dive,uuid"`
}

func (p *UserPagination) SetDefaultValue() {
	p.Pagination.SetDefaultValue()

	if p.RoleIds == nil {
		p.RoleIds = &[]uuid.UUID{}
	}
}
