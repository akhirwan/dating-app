package app

import (
	"context"
	"dating-app/src/delivery/container"
	"dating-app/src/delivery/http"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	Container container.Container
	Http      http.FiberApp
}

func Execute() *App {
	// start init container
	container := container.SetupContainer()

	// start http service
	http := http.ServeHttp(container)
	go http.Fiber.Listen(fmt.Sprintf(":%s", container.EnvironmentConfig.App.Port))

	return &App{
		Container: container,
		Http:      http,
	}
}

/* WithGracefulShutdown ... */
func (app *App) WithGracefulShutdown() error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go log.Println("Press ctrl+c to stop the service")

	sig := <-c /* Blocking for graceful */
	log.Printf("\n⛔️ Got %v signal. Shutting down app... \n", sig)

	/* Give some time to finish ongoing request */
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()
	defer ctx.Done()

	app.Http.Shutdown()
	log.Println("App has been completely shutdown. Asta Lavista! ")

	os.Exit(0)
	return nil
}
