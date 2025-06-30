package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func RoutesInitialization(app *fiber.App) {
	api := app.Group("/")

	api.Post("/tasks", postTask)
	api.Get("/tasks", getAllTasks)
	api.Get("/tasks/:title", getTask)
	api.Delete("/tasks/:title", deleteTask)
	app.Get("/swagger/*", swagger.HandlerDefault)

}
