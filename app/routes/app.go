package routes

import (
	handler "lazy-platform-auth/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func AppRouter(router fiber.Router, appHandler handler.AppHandler) {
	router.Get("", appHandler.GetApp)
	router.Get("/:id", appHandler.GetAppById)
	router.Post("", appHandler.CreateApp)
	router.Put("/:id", appHandler.UpdateApp)
	router.Delete("/:id", appHandler.DeleteApp)
}
