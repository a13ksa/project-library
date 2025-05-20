package handlers

import (
	"context"
	"fmt"
	"strconv"

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
		fmt.Println("BodyParser error:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Validate mandatory fields
	if book.Title == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Book title is missing!"})
	}

	if len(book.Genres) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "At least one genre is required!"})
	}

	if book.InventoryNumber == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Inventory number is required and cannot be zero!"})
	}

	if book.Signature == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Signature is required and cannot be empty!"})
	}

	collection := config.DB.Collection("books")

	// 1. Check if inv_number is unique
	countInv, err := collection.CountDocuments(context.Background(), bson.M{"inv_number": book.InventoryNumber})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if countInv > 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Inventory number must be unique"})
	}

	// 2. Check if a book with the same title and authors already exists
	filter := bson.M{
		"title":   book.Title,
		"authors": book.Authors,
	}
	var existingBook models.Book
	err = collection.FindOne(context.Background(), filter).Decode(&existingBook)
	if err == nil {
		// If book exists, signature must match
		if existingBook.Signature != book.Signature {
			return c.Status(400).JSON(fiber.Map{"error": "Signature must match existing book with same title and authors"})
		}
	}

	res, err := collection.InsertOne(context.Background(), book)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	book.ID = res.InsertedID.(primitive.ObjectID)
	return c.Status(201).JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	// Parse inventory number from URL param
	invNumberParam := c.Params("inv_number")
	invNumber, err := strconv.Atoi(invNumberParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid inventory number"})
	}

	updateData := new(models.Book)
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	collection := config.DB.Collection("books")

	// Build update map with only non-empty fields
	update := bson.M{}

	if updateData.Title != "" {
		update["title"] = updateData.Title
	}
	if len(updateData.Authors) > 0 {
		update["authors"] = updateData.Authors
	}
	if len(updateData.Genres) > 0 {
		update["genres"] = updateData.Genres
	}
	// Add other fields if needed...

	if len(update) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "No valid fields to update"})
	}

	result, err := collection.UpdateOne(
		context.Background(),
		bson.M{"inv_number": invNumber},
		bson.M{"$set": update},
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if result.MatchedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}

	return c.SendStatus(204) // No Content, update successful
}

func DeleteBook(c *fiber.Ctx) error {
	invNumberParam := c.Params("inv_number")
	invNumber, err := strconv.Atoi(invNumberParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid inventory number"})
	}

	collection := config.DB.Collection("books")

	result, err := collection.DeleteOne(context.Background(), bson.M{"inv_number": invNumber})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if result.DeletedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Book successfully deleted",
	})
}
