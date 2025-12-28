package permission

import (
	"net/http"

	"iam-service/constants"
	"iam-service/infrastructure/utils"
	"iam-service/modules/iam/permission/dtos/params"
	permissionService "iam-service/modules/iam/permission/services"

	"github.com/gofiber/fiber/v2"
)

func CreateHandler(ctx *fiber.Ctx) error {
	dto := ctx.Locals(constants.CTX_DTO_KEY).(*params.PermissionCreate)

	result := permissionService.CreatePermission(dto)
	return ctx.Status(http.StatusOK).JSON(utils.SuccessResult(result))
}

func PaginateHandler(ctx *fiber.Ctx) error {
	dto := ctx.Locals(constants.CTX_DTO_KEY).(*params.PermissionPagination)

	result := permissionService.PaginatePermissions(dto)

	return ctx.Status(http.StatusOK).JSON(utils.SuccessResult(result))
}

func GetByIdHandler(ctx *fiber.Ctx) error {
	paramId := ctx.Params("id")
	id := utils.UUIDChecker(paramId, "id")

	result := permissionService.GetPermissionById(id)
	return ctx.Status(http.StatusOK).JSON(utils.SuccessResult(result))
}

func UpdateByIdHandler(ctx *fiber.Ctx) error {
	paramId := ctx.Params("id")
	id := utils.UUIDChecker(paramId, "id")

	dto := ctx.Locals(constants.CTX_DTO_KEY).(*params.PermissionUpdate)

	result := permissionService.UpdatePermission(id, dto)
	return ctx.Status(http.StatusOK).JSON(utils.SuccessResult(result))
}

func DeleteByIdHandler(ctx *fiber.Ctx) error {
	paramId := ctx.Params("id")
	id := utils.UUIDChecker(paramId, "id")

	permissionService.DeletePermission(id)

	return ctx.Status(http.StatusOK).JSON(utils.SuccessResult(nil))
}
