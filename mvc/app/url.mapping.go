package app

import (
	"github.com/saravase/golang_fiber_mvc/mvc/controller"
)

func mapURLs() {
	app.Get("/user/:id", controller.UserController.Get)
	app.Post("/user", controller.UserController.Create)
}
