package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivesh/crud-app/bootstrap"
)

func main() {
	app := fiber.New()
	bootstrap.InitializeApp(app)
}
