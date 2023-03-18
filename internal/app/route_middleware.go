package app

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func MiddlewareRoutes(router *fiber.App) {
	// Last middleware to match anything
	router.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusNotFound)
		// => 404 "Not Found"
	})
}
