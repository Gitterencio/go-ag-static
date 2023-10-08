package main

import (
	"github.com/gitterencio/go-ag-static/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/django/v3"
)

func main() {

	engine := django.New("./public/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//MIDLEWARE
	app.Use(logger.New())

	apiGroup := app.Group("api")
	apiGroup.Route("/", routes.APIRouter)

	serverGroup := app.Group("lc")
	serverGroup.Route("/", routes.ServerRouter)

	//STATIC SPA ANGULAR

	app.Static("/", "front_wha_tar")

	app.Get("/*", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./front_wha_tar/index.html")
	})

	app.Listen(":3000")
}
