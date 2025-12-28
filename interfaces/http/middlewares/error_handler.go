package middlewares

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"iam-service/constants"
	"iam-service/infrastructure/config"
	appErr "iam-service/infrastructure/errors"
	"iam-service/infrastructure/utils"
)

func ErrorHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer func() error {
			if rec := recover(); rec != nil {

				// For most Http errors
				if err, ok := rec.(*appErr.AppError); ok {
					return ctx.Status(err.StatusCode).JSON(utils.ErrorResult(err.StatusCode, err))
				}

				if err, ok := rec.(*appErr.TValidationError); ok {
					mapErrs := make([]*appErr.TValidationError, 0)
					mapErrs = append(mapErrs, err)
					return ctx.Status(http.StatusBadRequest).JSON(utils.ValidationErrorsResult(mapErrs))
				}

				resp := utils.ErrorResultWithMessage(http.StatusInternalServerError, "Internal server error")

				trace := string(debug.Stack())
				errString := fmt.Sprintf("%s\n%s", rec, trace)
				logrus.Error(errString)

				// 🔥 dev only
				if config.APP_ENV == constants.APP_ENV_DEV {
					resp.Trace = trace
				}

				return ctx.Status(http.StatusInternalServerError).JSON(
					resp,
				)
			}

			return nil
		}()

		return ctx.Next()
	}
}
