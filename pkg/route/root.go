package route

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/ivandrenjanin/go-fiber-htmx-boilerplate/platform/database"
)

func GeneralRoute(a *fiber.App) {
	a.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Api!",
			"docs":    "/swagger/index.html",
			"status":  "/health",
		})
	})

	a.Get("/health", func(c *fiber.Ctx) error {
		err := database.GetDB().Ping()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
				"success": false,
			})
		}

		return c.JSON(fiber.Map{
			"message": "Health Check",
			"success": true,
		})
	})
}

func SwaggerRoute(a *fiber.App) {
	route := a.Group("/swagger")
	route.Get("*", swagger.HandlerDefault)
}

func NotFoundRoute(a *fiber.App) {
	a.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Route not found",
		})
	})
}
