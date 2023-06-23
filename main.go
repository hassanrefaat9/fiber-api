package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hassanrefaat9/database"
	"github.com/hassanrefaat9/router"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}
	app := fiber.New()
	database.ConnectDB()

	router.SetUpRouter(app)
	log.Fatal(app.Listen(":3000"))
}
