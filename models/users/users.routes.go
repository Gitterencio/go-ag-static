package users

import "github.com/gofiber/fiber/v2"

func UserRouter(router fiber.Router) {

	router.Get("/", HandleIndex)
	router.Post("/", HandleCreateUser)
}
