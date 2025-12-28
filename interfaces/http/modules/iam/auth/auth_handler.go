package auth

import (
	"iam-service/constants"
	"iam-service/infrastructure/utils"
	"iam-service/modules/iam/auth/dtos/params"
	authService "iam-service/modules/iam/auth/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SignInHandler(ctx *fiber.Ctx) error {
	dto := ctx.Locals(constants.CTX_DTO_KEY).(*params.AuthSignInParam)

	result := authService.SignIn(dto)
	return ctx.Status(http.StatusOK).JSON(utils.SuccessResult(result))
}

func ProfileHandler(ctx *fiber.Ctx) error {
	user := ctx.Locals(constants.CTX_TOKEN_KEY).(string)

	result := authService.Profile(user)
	return ctx.Status(http.StatusOK).JSON(utils.SuccessResult(result))
}

func RenewTokenHandler(ctx *fiber.Ctx) error {
	refreshToken := ctx.Locals(constants.CTX_TOKEN_KEY).(string)

	result := authService.RenewToken(refreshToken)
	return ctx.Status(http.StatusOK).JSON(utils.SuccessResult(result))
}

func SignOutHandler(ctx *fiber.Ctx) error {
	refreshToken := ctx.Locals(constants.CTX_TOKEN_KEY).(string)

	authService.SignOut(refreshToken)

	return ctx.Status(http.StatusOK).JSON(utils.SuccessResult(nil))
}
