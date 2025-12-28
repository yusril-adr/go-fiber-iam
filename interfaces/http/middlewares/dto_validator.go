package middlewares

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"iam-service/constants"
	"iam-service/infrastructure/errors"
	"iam-service/infrastructure/types"
	"iam-service/infrastructure/utils"
	"iam-service/infrastructure/validators"
)

func registerCustomValidator(validator *validator.Validate) {
	validator.RegisterValidation(validators.PASSWORD_KEY, validators.PasswordValidator)
	validator.RegisterValidation(validators.PAGINATION_ORDER_KEY, validators.PaginationOrderValidator)
}

func ValidateDto[T any]() fiber.Handler {
	var validate = validator.New()
	registerCustomValidator(validate)

	return func(ctx *fiber.Ctx) error {
		obj := new(T)

		_ = ctx.BodyParser(obj)
		_ = ctx.QueryParser(obj)

		if err := validate.Struct(obj); err != nil {
			// Should map to ValidationErrors
			errs, isValid := err.(validator.ValidationErrors)
			if isValid {
				mappedErrors := make([]*errors.TValidationError, 0)

				for _, validationErr := range errs {
					mappedErrors = append(mappedErrors, errors.ValidationError(
						utils.StringToSnakeCase(validationErr.Field()), // Fully qualified field name
						validationMessage(validationErr),
					))
				}

				return ctx.Status(http.StatusBadRequest).JSON(utils.ValidationErrorsResult(mappedErrors))
			}

			panic(err.Error())
		}

		if dv, ok := any(obj).(types.TDefaultValueParam); ok {
			dv.SetDefaultValue()
		}

		ctx.Locals(constants.CTX_DTO_KEY, obj)
		return ctx.Next()
	}
}

func validationMessage(e validator.FieldError) string {
	field := utils.StringToSnakeCase(e.Field())

	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s items", field, e.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s items", field, e.Param())
	case validators.PASSWORD_KEY:
		return fmt.Sprintf("%s %s", field, validators.PASSWORD_MESSAGE)
	case validators.PAGINATION_ORDER_KEY:
		return fmt.Sprintf("%s %s", field, validators.PAGINATION_ORDER_MESSAGE)
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}
