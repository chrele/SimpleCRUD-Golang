package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	router := fiber.New()
	app := fiber.New()

	//Set main end-point
	app.Mount("/api", router)

	//Set sub end-points
	router.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "server is online",
		})
	})

	log.Fatal(app.Listen(":8000"))
}
