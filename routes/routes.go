package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/project/project-skripsi/go-be/controllers"
)

// setup routes
func Setup(app *fiber.App) {

	// API Post
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/logout", controllers.Logout)

	// API Get
	app.Get("/api/user", controllers.User)

}
