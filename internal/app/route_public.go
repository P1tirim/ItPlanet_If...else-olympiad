package app

import (
	"api/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(router *fiber.App, handlers *http.Handlers) {
	router.Post("/registration", handlers.Registration)
}
