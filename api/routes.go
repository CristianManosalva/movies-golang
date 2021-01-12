package api

import "github.com/gofiber/fiber"

//SetupMoviesRoutes default comment
func SetupMoviesRoutes(app *fiber.App) {
	s := start()
	group := app.Group("/movies")

	group.Get("", s.SearchMovieHandler)
}
