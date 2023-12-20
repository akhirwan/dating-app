package http

import (
	"dating-app/src/delivery/container"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FiberApp struct {
	Fiber *fiber.App
}

func ServeHttp(container container.Container) FiberApp {

	app := fiber.New(
		fiber.Config{
			Prefork:       false,
			CaseSensitive: true,
			StrictRouting: true,
			ServerHeader:  container.EnvironmentConfig.App.Name,
			AppName:       fmt.Sprintf("%s v%s", container.EnvironmentConfig.App.Name, container.EnvironmentConfig.App.Version),
		},
	)

	/* Set global middleware */
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		MaxAge:       0,
	}))

	RouterGroup(app, SetupHandler(container))

	return FiberApp{app}
}

// Shutdown
func (f FiberApp) Shutdown() {
	f.Fiber.Shutdown()
}
