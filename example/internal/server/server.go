package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func NewServer(log *logrus.Logger) *fiber.App {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		log.Infof("%s %s", c.Method(), c.Path())
		return c.Next()
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	return app
}
