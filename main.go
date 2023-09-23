package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/zanamira43/todo/database"
	"github.com/zanamira43/todo/routes"
)

func main() {
	// database connection
	database.Connect()

	// new app instance with fiber new methods
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	// use route system
	routes.Setup(app)

	// runnig app on port number
	app.Listen("0.0.0.0:3000")
}
