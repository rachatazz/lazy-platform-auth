package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok", "time_stamp": time.Now()})
}
