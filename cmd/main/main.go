package main

import (
	"github.com/CristianManosalva/movies-golang/api"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	api.SetupMoviesRoutes(app)

	_ = app.Listen("3000")
}
