package utils

import (
	"math"
	"net/http"

	"iam-service/infrastructure/errors"
	"iam-service/infrastructure/types"
)

func SuccessResult(data interface{}) *types.TBaseResult {
	return &(types.TBaseResult{
		StatusCode: http.StatusOK,
		Data:       &data,
	})
}

func SuccessCreatedResult(data interface{}) *types.TBaseResult {
	return &(types.TBaseResult{
		StatusCode: http.StatusCreated,
		Data:       &data,
	})
}

func ErrorResult(status int, err error) *types.TErrorResult {
	return &(types.TErrorResult{
		StatusCode: status,
		Message:    err.Error(),
	})
}

func ErrorResultWithMessage(status int, message string) *types.TErrorResult {
	return &(types.TErrorResult{
		StatusCode: status,
		Message:    message,
	})
}

func ValidationErrorsResult(errs []*errors.TValidationError) *types.TValidationResult {
	return &(types.TValidationResult{
		Status: http.StatusBadRequest,
		Errors: errs,
	})
}

func PaginationBuilder[T any](items []T, meta types.TMetaResult) *types.TPaginationResult[T] {
	return &(types.TPaginationResult[T]{
		Items: &items,
		Meta:  &meta,
	})
}

func PaginationMetaBuilder(page int, perPage int, total int64) *types.TMetaResult {
	// Calculate total page
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))

	return &(types.TMetaResult{
		Page:      page,
		PerPage:   perPage,
		Total:     total,
		TotalPage: totalPage,
	})
}
