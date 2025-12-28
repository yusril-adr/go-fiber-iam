package params

import (
	"iam-service/constants"
	"iam-service/infrastructure/types"
)

type Pagination struct {
	types.TDefaultValueParam
	Page    *int    `query:"page" validate:"omitempty"`
	PerPage *int    `query:"per_page" validate:"omitempty"`
	SortBy  *string `query:"sort_by" validate:"omitempty"`
	Order   *string `query:"order" validate:"omitempty,pagination_order"`
	Search  *string `query:"search" validate:"omitempty"`
}

func (p *Pagination) SetDefaultValue() {
	if p.Page == nil {
		page := 1
		p.Page = &page
	}

	if p.PerPage == nil {
		perPage := 10
		p.PerPage = &perPage
	}

	if p.SortBy == nil {
		sortBy := "updated_at"
		p.SortBy = &sortBy
	}

	if p.Order == nil {
		order := constants.PAGINATION_ORDER_DESC
		p.Order = &order
	}
}
