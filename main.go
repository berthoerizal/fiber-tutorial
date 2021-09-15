package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/erizal", func(c *fiber.Ctx) error {
		return c.Send([]byte("Hello, erizal!"))
	})

	log.Fatal(app.Listen(":3000"))
}
