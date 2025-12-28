package role

import (
	"iam-service/interfaces/http/middlewares"
	"iam-service/modules/iam/role/dtos/params"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	rg := router.Group("api/v1/iam/roles")

	rg.Use(
		middlewares.BearerAccessTokenAuth(),
	)

	rg.Post("/", middlewares.ValidateDto[params.RoleCreate](), CreateHandler)
	rg.Get("/", middlewares.ValidateDto[params.RolePagination](), PaginateHandler)
	rg.Get("/:id", GetByIdHandler)
	rg.Put("/:id", middlewares.ValidateDto[params.RoleUpdate](), UpdateByIdHandler)
	rg.Delete("/:id", DeleteByIdHandler)
}
