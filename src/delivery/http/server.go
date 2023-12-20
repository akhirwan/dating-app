package http

import (
	"dating-app/src/delivery/container"

	"github.com/gofiber/fiber/v2"
)

type FiberApp struct {
	Server *fiber.App
}

func ServeHttp(container container.Container) FiberApp {
	// handler := SetupHandler(container)

	app := fiber.New()

	// iniate router v1
	// RouterGroupV1(app, handler)

	return FiberApp{app}
}

// Shutdown
func (f FiberApp) Shutdown() {
	f.Server.Shutdown()
}
