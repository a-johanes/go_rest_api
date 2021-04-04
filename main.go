package main

import (
	"fmt"
	"log"

	"github.com/a-johanes/go_rest_api/controllers"
	"github.com/a-johanes/go_rest_api/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	if err := models.ConnectDatabase(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(recover.New()) // recover from panic
	app.Use(logger.New())  // log request and response
	app.Use(validatorMiddleware())

	app.Get("/books", controllers.FindBooks)
	app.Get("/books/:id", controllers.FindBook)
	app.Post("/books", controllers.CreateBook)
	app.Patch("/books/:id", controllers.UpdateBook)
	app.Delete("/books/:id", controllers.DeleteBook)

	log.Fatal(app.Listen(":4321"))
}

func validatorMiddleware() func(*fiber.Ctx) error {
	validate := validator.New()
	fmt.Println("called")
	return func(c *fiber.Ctx) error {
		c.Locals("validator", validate)
		return c.Next()
	}
}
