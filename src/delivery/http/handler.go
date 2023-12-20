package http

import "dating-app/src/delivery/container"

type handler struct{}

func SetupHandler(container container.Container) *handler {
	return &handler{}
}
