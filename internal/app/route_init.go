package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func InitRoutes(router *fiber.App) {
	router.Use(recover.New())

	router.Use(cors.New(cors.Config{
		AllowHeaders:  "*",
		ExposeHeaders: "POST,GET,PUT,DELETE",
	}))
}
