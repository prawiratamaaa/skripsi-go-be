package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/project/project-skripsi/go-be/database"
	"github.com/project/project-skripsi/go-be/routes"
	"github.com/sirupsen/logrus"
)

func main() {
	database.Connection()

	//log below
	log := logrus.New()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	log.Warn("Warning Message")
	log.Error("Error Message")

	app.Listen(":8080")

}
