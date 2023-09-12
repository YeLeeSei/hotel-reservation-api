package main

import (
	"flag"

	"github.com/YeLeeSei/hotel-reservation-api/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "the listen address of API server")
	flag.Parse()
	app := fiber.New()
	apiv1 := app.Group("/api/v1")
	apiv1.Get("/foo", handleFoo)
	apiv1.Get("/user/:id", api.HandleGetUser)
	apiv1.Get("/user", api.HandleGetUsers)
	app.Listen(*listenAddr)
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "allright"})
}
