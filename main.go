package main

import (
	"fmt"
	"log"
	"os"

	"github.com/a13ksa/project-library/config"
	"github.com/a13ksa/project-library/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Starting server...")

	app := fiber.New()

	config.ConnectDB()
	routes.RegisterBookRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Fatal(app.Listen(":" + port))
}
