package main

import (
	"log"
	"main/database"
	"main/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))
	if app != nil {
		routes.Setup(app)
	}

	log.Fatal(app.Listen(":3000"))
}
