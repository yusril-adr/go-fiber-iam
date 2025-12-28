package permission

import (
	"iam-service/interfaces/http/middlewares"
	"iam-service/modules/iam/permission/dtos/params"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	rg := router.Group("api/v1/iam/permissions")

	rg.Use(
		middlewares.BearerAccessTokenAuth(),
	)

	rg.Post("/", middlewares.ValidateDto[params.PermissionCreate](), CreateHandler)
	rg.Get("/", middlewares.ValidateDto[params.PermissionPagination](), PaginateHandler)
	rg.Get("/:id", GetByIdHandler)
	rg.Put("/:id", middlewares.ValidateDto[params.PermissionUpdate](), UpdateByIdHandler)
	rg.Delete("/:id", DeleteByIdHandler)
}
