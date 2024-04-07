package main

import (
	"fmt"
	"log"
)

// START OMIT

func main() {
	app := fiber.New()

	api := app.Group("/api")
	api.Get("/list", func(c fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})
	api.Post("/register", func(c fiber.Ctx) error {
		return c.SendString("I'm a POST request!")
	})

	app.Get("/user/:name/books/:title", func(c fiber.Ctx) error {
		fmt.Fprintf(c, "%s\n", c.Params("name"))
		fmt.Fprintf(c, "%s\n", c.Params("title"))
		return nil
	})
	log.Fatal(app.Listen(":3000"))
}

// END OMIT
