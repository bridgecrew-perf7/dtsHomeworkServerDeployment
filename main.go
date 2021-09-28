package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/sammy/detail", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"username":  "SammyShark",
			"followers": 987,
		})
	})

	app.Get("/jesse/detail", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"username":  "JesseOctopus",
			"followers": 432,
		})
	})

	app.Get("/drew/detail", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"username":  "DrewSquid",
			"followers": 321,
		})
	})

	app.Get("/jamie/detail", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"username":  "JamieMantisShrimp",
			"followers": 654,
		})
	})

	// -----------------------------------------------------------------//

	app.Get("/follower/SammyShark", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"followers": 987,
		})
	})

	app.Get("/follower/JesseOctopus", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"followers": 432,
		})
	})

	app.Get("/follower/DrewSquid", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"followers": 321,
		})
	})

	app.Get("/follower/JamieMantisShrimp", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"username":  "JamieMantisShrimp",
			"followers": 654,
		})
	})

	app.Listen(":" + os.Getenv("PORT"))
}
