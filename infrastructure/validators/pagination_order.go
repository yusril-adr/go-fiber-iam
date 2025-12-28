package validators

import (
	"fmt"
	"iam-service/constants"

	"github.com/go-playground/validator/v10"
)

var (
	PAGINATION_ORDER_KEY     = "pagination_order"
	PAGINATION_ORDER_MESSAGE = fmt.Sprintf(
		"must be %s or %s", constants.PAGINATION_ORDER_ASC, constants.PAGINATION_ORDER_DESC,
	)
)

func PaginationOrderValidator(fl validator.FieldLevel) bool {
	orderValue := fl.Field().String()

	return orderValue == constants.PAGINATION_ORDER_ASC || orderValue == constants.PAGINATION_ORDER_DESC
}
