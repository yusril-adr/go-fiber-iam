package user

import (
	"iam-service/interfaces/http/middlewares"
	"iam-service/modules/iam/user/dtos/params"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	rg := router.Group("api/v1/iam/users")

	rg.Use(
		middlewares.BearerAccessTokenAuth(),
	)

	rg.Post("/", middlewares.ValidateDto[params.UserCreate](), CreateHandler)
	rg.Get("/", middlewares.ValidateDto[params.UserPagination](), PaginateHandler)
	rg.Get("/:id", GetByIdHandler)
	rg.Put("/:id", middlewares.ValidateDto[params.UserUpdate](), UpdateByIdHandler)
	rg.Delete("/:id", DeleteByIdHandler)
}
