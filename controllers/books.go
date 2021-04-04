package controllers

import (
	"github.com/a-johanes/go_rest_api/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CreateBookInput struct {
	Title  string `validate:"required"`
	Author string `validate:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func FindBooks(c *fiber.Ctx) error {
	var books []models.Book
	models.DB.Find(&books)

	return c.JSON(fiber.Map{"data": books})
}

func FindBook(c *fiber.Ctx) error { // Get model if exist
	var book models.Book

	models.DB.Where("id = ?", c.Params("id")).Limit(1).Find(&book)

	if book == (models.Book{}) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found!"})
	}

	return c.JSON(fiber.Map{"data": book})
}

func CreateBook(c *fiber.Ctx) error {
	input := new(CreateBookInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	validate := c.Locals("validator").(*validator.Validate)
	if err := validate.Struct(*input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	return c.JSON(fiber.Map{"data": book})
}

func UpdateBook(c *fiber.Ctx) error {
	// Get model if exist
	var book models.Book
	models.DB.Where("id = ?", c.Params("id")).Limit(1).Find(&book)

	if book == (models.Book{}) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found!"})
	}

	// Validate input
	input := new(UpdateBookInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	models.DB.Model(&book).Updates(models.Book{Title: input.Title, Author: input.Author})

	return c.JSON(fiber.Map{"data": book})
}

func DeleteBook(c *fiber.Ctx) error {
	// Get model if exist
	var book models.Book
	models.DB.Where("id = ?", c.Params("id")).Limit(1).Find(&book)

	if book == (models.Book{}) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found!"})
	}

	models.DB.Delete(&book)

	return c.JSON(fiber.Map{"data": book})
}
