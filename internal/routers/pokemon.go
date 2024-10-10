package routers

import "github.com/gofiber/fiber/v2"

func pokemonRoutes(app fiber.Router) {
	pokemon := app.Group("/pokemon")

	pokemon.Get("/", getPokemon)
	pokemon.Get("/:id", getPokemonByID)
}

func getPokemon(c *fiber.Ctx) error {
	return c.SendString("Get all Pokemon")
}

func getPokemonByID(c *fiber.Ctx) error {
	return c.SendString("Get Pokemon by ID")
}
