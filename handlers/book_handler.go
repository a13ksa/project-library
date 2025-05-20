package handlers

import (
	"context"

	"github.com/a13ksa/project-library/config"
	"github.com/a13ksa/project-library/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	collection := config.DB.Collection("books")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var book models.Book
		if err := cursor.Decode(&book); err != nil {
			return err
		}
		books = append(books, book)
	}
	return c.JSON(books)
}

func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if book.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Book name is missing!"})
	}

	collection := config.DB.Collection("books")
	res, err := collection.InsertOne(context.Background(), book)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	book.ID = res.InsertedID.(primitive.ObjectID)
	return c.Status(201).JSON(book)
}
