package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"iam-service/infrastructure/config"
	"iam-service/interfaces/http/middlewares"
	"iam-service/interfaces/http/modules/iam/auth"
	"iam-service/interfaces/http/modules/iam/permission"
	"iam-service/interfaces/http/modules/iam/role"
	"iam-service/interfaces/http/modules/iam/user"
)

func StartListen() {
	app := fiber.New()

	app.Use(
		cors.New(),
		middlewares.ErrorHandler(),
	)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Service is running, and Healty",
		})
	})

	/* ----------------------------- Register Router ---------------------------- */
	auth.RegisterRoutes(app)
	user.RegisterRoutes(app)
	role.RegisterRoutes(app)
	permission.RegisterRoutes(app)

	app.Listen(fmt.Sprintf(":%s", config.PORT))
}
