package routes

import (
	"github.com/a13ksa/project-library/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterBookRoutes(app *fiber.App) {
	api := app.Group("/api")
	books := api.Group("/books")

	books.Get("/", handlers.GetBooks)
	books.Post("/", handlers.CreateBook)
	books.Patch("/:inv_number", handlers.UpdateBook)
	books.Delete("/:inv_number", handlers.DeleteBook)
}
