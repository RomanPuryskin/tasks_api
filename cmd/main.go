package main

// @title tasks_api
// @version 1.0
// @description API для управления задачами
// @host localhost:8080

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/tasks_api/docs"
	"github.com/tasks_api/pkg/routes"
	"github.com/tasks_api/pkg/storage"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	if _, err := storage.NewStorage(); err != nil {
		log.Fatal("ошибка инициализации хранилища:", err)
	}

	routes.RoutesInitialization(app)

	log.Fatal(app.Listen(":8080"))
}
