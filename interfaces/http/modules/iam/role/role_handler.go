package role

import (
	"net/http"

	"iam-service/constants"
	"iam-service/infrastructure/utils"
	"iam-service/modules/iam/role/dtos/params"
	roleService "iam-service/modules/iam/role/services"

	"github.com/gofiber/fiber/v2"
)

func CreateHandler(ctx *fiber.Ctx) error {
	dto := ctx.Locals(constants.CTX_DTO_KEY).(*params.RoleCreate)

	result := roleService.CreateRole(dto)
	return ctx.Status(http.StatusOK).JSON(utils.SuccessResult(result))
}

func PaginateHandler(ctx *fiber.Ctx) error {
	dto := ctx.Locals(constants.CTX_DTO_KEY).(*params.RolePagination)

	result := roleService.PaginateRoles(dto)

	return ctx.Status(http.StatusOK).JSON(utils.SuccessResult(result))
}

func GetByIdHandler(ctx *fiber.Ctx) error {
	paramId := ctx.Params("id")
	id := utils.UUIDChecker(paramId, "id")

	result := roleService.GetRoleById(id)
	return ctx.Status(http.StatusOK).JSON(utils.SuccessResult(result))
}

func UpdateByIdHandler(ctx *fiber.Ctx) error {
	paramId := ctx.Params("id")
	id := utils.UUIDChecker(paramId, "id")

	dto := ctx.Locals(constants.CTX_DTO_KEY).(*params.RoleUpdate)

	result := roleService.UpdateRole(id, dto)
	return ctx.Status(http.StatusOK).JSON(utils.SuccessResult(result))
}

func DeleteByIdHandler(ctx *fiber.Ctx) error {
	paramId := ctx.Params("id")
	id := utils.UUIDChecker(paramId, "id")

	roleService.DeleteRole(id)

	return ctx.Status(http.StatusOK).JSON(utils.SuccessResult(nil))
}
