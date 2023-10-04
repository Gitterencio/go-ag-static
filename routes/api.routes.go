package routes

import (
	"github.com/gitterencio/go-ag-static/models/users"
	"github.com/gofiber/fiber/v2"
)

func APIRouter(router fiber.Router) {

	userGroup := router.Group("users")
	userGroup.Route("/", users.UserRouter)

	group := router.Group("same")
	group.Get("/", HandleIndex)
	group.Post("/", HandleCreateUser)

}
