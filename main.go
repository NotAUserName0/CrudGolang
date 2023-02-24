package main

import (
	"crud/databases"
	"crud/models"
	"crud/routing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	databases.Connection() //conecta
	models.InitDatabase()
	app.Use(cors.New(cors.Config{
        AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
        AllowOrigins:     "*",
        AllowCredentials: true,
        AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
    }))
	routing.UserRoutes(app)
	app.Listen("127.0.0.1:3000")
}