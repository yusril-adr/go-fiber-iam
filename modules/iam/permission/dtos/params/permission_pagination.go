package params

import (
	"iam-service/infrastructure/dtos/params"
)

type PermissionPagination struct {
	params.Pagination
}

func (p *PermissionPagination) SetDefaultValue() {
	p.Pagination.SetDefaultValue()
}
