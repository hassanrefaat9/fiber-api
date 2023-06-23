package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hassanrefaat9/handler"
)

func SetUpRouter(app *fiber.App) {

	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	product := api.Group("/product")
	product.Get("/", handler.GetAllProducts)
	product.Get("/:id", handler.GetProduct)
	product.Post("/", handler.CreateProduct)
	product.Delete("/:id", handler.DeleteProduct)
}
