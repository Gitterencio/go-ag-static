package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type User struct {
	Id        string
	Firstname string
	Lastname  string
}

func HandleIndex(c *fiber.Ctx) error {
	user := User{
		Firstname: "ja",
		Lastname:  "Reja",
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func HandleCreateUser(c *fiber.Ctx) error {
	user := User{}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.Id = uuid.NewString()
	return c.Status(fiber.StatusOK).JSON(user)
}
