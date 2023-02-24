package routing

import (
	"crud/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/", controllers.GetUsers)
	app.Get("/get/:id",controllers.GetUser)
	app.Post("/add", controllers.AddUsers)
	app.Delete("/delete/:id", controllers.DelUsers)
	app.Put("/put/:id", controllers.PutUsers)
}
