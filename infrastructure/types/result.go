package types

import "iam-service/infrastructure/errors"

type TBaseResult struct {
	StatusCode int `json:"status_code"`
	Data       any `json:"data"`
}

type TErrorResult struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Trace      string `json:"trace,omitempty"`
}

type TValidationResult struct {
	Status  int                        `json:"status"`
	Errors  []*errors.TValidationError `json:"errors,omitempty"`
	Message string                     `json:"message,omitempty"`
}

type TPaginationResult[T any] struct {
	Items *[]T         `json:"items"`
	Meta  *TMetaResult `json:"meta"`
}

type TMetaResult struct {
	Page      int   `json:"page"`
	PerPage   int   `json:"per_page"`
	Total     int64 `json:"total"`
	TotalPage int   `json:"total_page"`
}
