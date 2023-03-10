package main

import (
	//"fmt"
	"log"

	chat "github.com/asolpshinning/chaingpt/tools/gpt-tools"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	// let all cors be allowed
	app.Use(cors.New())
	app.Get("/", Home)
	app.Get("/hello/:name", HelloHandler)
	app.Post("/chat", ChatHandler)

	app.Listen(":5176")
}

func Home(c *fiber.Ctx) error {
	return c.SendString("This is just the home page")
}

func HelloHandler(c *fiber.Ctx) error {
	name := c.Params("name")
	res := "Hello " + name
	return c.SendString(res)
}

func ChatHandler(c *fiber.Ctx) error {
	//get request body from the fiber context

	req := c.Body()

	// convert the request body to a string
	prompt := string(req)

	res, err := chat.ChatGPT(prompt)
	if err != nil {
		return c.SendString("Error")
	}

	log.Println(res)
	return c.JSON(res)
}
