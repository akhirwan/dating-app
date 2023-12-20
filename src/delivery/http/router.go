package http

import "github.com/gofiber/fiber/v2"

func RouterGroup(app *fiber.App, handler *handler) {

	check := app.Group("/ping")
	{
		check.Get("/", func(c *fiber.Ctx) error {
			return c.JSON("pong!")
		})
	}

}
