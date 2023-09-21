package routes

import (
	handler "lazy-platform-auth/app/handlers"
	"lazy-platform-auth/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(
	router fiber.Router,
	middleware middlewares.AuthMiddleware,
	userHandler handler.UserHandler,
) {
	router.Get("/me", middleware.TokenMiddleware(), userHandler.GetMe)
	router.Put("/me", middleware.TokenMiddleware(), userHandler.UpdateUser)
	router.Post("/register", userHandler.CreateUser)
	router.Post("/token", userHandler.Token)
	router.Post("/refresh-token", userHandler.RefreshToken)
	router.Post("/change-password", middleware.TokenMiddleware(), userHandler.ChangePassword)
	router.Post("/forgot-password", userHandler.ForgotPassword)
	router.Post("/reset-password", userHandler.ResetPassword)
	router.Post("/request-verify", userHandler.ResetPassword)
	router.Delete("/token", middleware.RefreshTokenMiddleware(), userHandler.DeleteToken)
}
