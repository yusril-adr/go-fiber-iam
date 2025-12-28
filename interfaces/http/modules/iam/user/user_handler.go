package user

import (
	"iam-service/constants"
	"iam-service/infrastructure/utils"
	"iam-service/modules/iam/user/dtos/params"
	userService "iam-service/modules/iam/user/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateHandler(ctx *fiber.Ctx) error {
	dto := ctx.Locals(constants.CTX_DTO_KEY).(*params.UserCreate)

	result := userService.CreateUser(dto)
	return ctx.Status(http.StatusOK).JSON(utils.SuccessResult(result))
}

func PaginateHandler(ctx *fiber.Ctx) error {
	dto := ctx.Locals(constants.CTX_DTO_KEY).(*params.UserPagination)

	result := userService.PaginateUsers(dto)

	return ctx.Status(http.StatusOK).JSON(utils.SuccessResult(result))
}

func GetByIdHandler(ctx *fiber.Ctx) error {
	paramId := ctx.Params("id")
	id := utils.UUIDChecker(paramId, "id")

	result := userService.GetUserById(id)
	return ctx.Status(http.StatusOK).JSON(utils.SuccessResult(result))
}

func UpdateByIdHandler(ctx *fiber.Ctx) error {
	paramId := ctx.Params("id")
	id := utils.UUIDChecker(paramId, "id")

	dto := ctx.Locals(constants.CTX_DTO_KEY).(*params.UserUpdate)

	result := userService.UpdateUser(id, dto)
	return ctx.Status(http.StatusOK).JSON(utils.SuccessResult(result))
}

func DeleteByIdHandler(ctx *fiber.Ctx) error {
	paramId := ctx.Params("id")
	id := utils.UUIDChecker(paramId, "id")

	userService.DeleteUser(id)

	return ctx.Status(http.StatusOK).JSON(utils.SuccessResult(nil))
}
