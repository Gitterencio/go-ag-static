package main

import (
	"github.com/gitterencio/go-ag-static/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	//MIDLEWARE
	app.Use(logger.New())

	apiGroup := app.Group("api")
	apiGroup.Route("/", routes.APIRouter)

	//STATIC SPA ANGULAR
	app.Static("/*", "./front_wha_tar")

	app.Listen(":3000")
}
