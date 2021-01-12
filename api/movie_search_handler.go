package api

import "github.com/gofiber/fiber"

// SearchMovieHandler one comment
func (w *WebServices) SearchMovieHandler(c *fiber.Ctx) {
	res, err := w.s.search.Search(MovieFilter{})

	if err != nil {
		err = fiber.NewError(400, "cannot bring movies")
		c.Next(err)
	}

	c.JSON(res)
}
