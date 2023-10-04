package routes

import "github.com/gofiber/fiber/v2"

func ServerRouter(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		// Render with and extends
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	router.Get("/embed", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("embed", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main2")
	})

}
