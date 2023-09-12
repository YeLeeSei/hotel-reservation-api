package api

import (
	"github.com/YeLeeSei/hotel-reservation-api/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "Jake",
		LastName:  "Smith",
	}
	return c.JSON(user)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("Lebron James")
}
