package auth

import (
	"iam-service/interfaces/http/middlewares"
	"iam-service/modules/iam/auth/dtos/params"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	rg := router.Group("api/v1/iam/auth")

	rg.Post("/sign-in", middlewares.ValidateDto[params.AuthSignInParam](), SignInHandler)
	rg.Get("/profile", middlewares.BearerAccessTokenAuth(), ProfileHandler)
	rg.Get("/renew-token", middlewares.BearerRefreshTokenAuth(), RenewTokenHandler)
	rg.Post("/sign-out", middlewares.BearerRefreshTokenAuth(), SignOutHandler)
}
