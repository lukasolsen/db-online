package routers

import "github.com/gofiber/fiber/v2"

func SetupRouter(app *fiber.App) {
	api := app.Group("/api/v1")

	// Pokemon
	pokemonRoutes(api)
}
