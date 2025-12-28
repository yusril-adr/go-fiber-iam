package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"iam-service/constants"
	"iam-service/infrastructure/config"
	"iam-service/infrastructure/errors"
	"iam-service/infrastructure/types"
	"iam-service/infrastructure/utils"
)

func BearerAccessTokenAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			logrus.Warn("Authorization header is empty")
			panic(errors.Unauthorized("Unauthorized"))
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			logrus.Warn("Invalid Authorization header")
			panic("Unauthorized")
		}

		token := strings.TrimSpace(parts[1])
		if token == "" {
			logrus.Warn("Token is empty")
			panic(errors.Unauthorized("Unauthorized"))
		}

		userPayload := utils.ParseToken[*types.TUserPayload](token, config.JWT_ACCESS_TOKEN_SECRET)

		ctx.Locals(constants.CTX_USER_KEY, userPayload)
		ctx.Locals(constants.CTX_TOKEN_KEY, token)

		return ctx.Next()
	}
}

func BearerRefreshTokenAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			logrus.Warn("Authorization header is empty")
			panic(errors.Unauthorized("Unauthorized"))
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			logrus.Warn("Invalid Authorization header")
			panic("Unauthorized")
		}

		token := strings.TrimSpace(parts[1])
		if token == "" {
			logrus.Warn("Token is empty")
			panic(errors.Unauthorized("Unauthorized"))
		}

		userPayload := utils.ParseToken[*types.TUserPayload](token, config.JWT_REFRESH_TOKEN_SECRET)

		ctx.Locals(constants.CTX_USER_KEY, userPayload)
		ctx.Locals(constants.CTX_TOKEN_KEY, token)

		return ctx.Next()
	}
}
