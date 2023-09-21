package main

import (
	"fmt"

	"lazy-platform-auth/app/adapter"
	handler "lazy-platform-auth/app/handlers"
	"lazy-platform-auth/app/middlewares"
	repository "lazy-platform-auth/app/repositories"
	"lazy-platform-auth/app/routes"
	service "lazy-platform-auth/app/services"
	"lazy-platform-auth/config"
	"lazy-platform-auth/database"
	"lazy-platform-auth/logs"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configService := config.ConfigService()
	// fmt.Printf("config: %v ", configService.DataBaseUri)
	db := database.InitDatabase(configService)

	appRepository := repository.NewAppRepository(db)
	userRepository := repository.NewUserRepository(db)
	sessionRepository := repository.NewSessionRepository(db)
	elasticMailAdapter := adapter.NewElasticEmailAdapter(configService)

	appService := service.NewAppService(appRepository)
	userService := service.NewAccountService(
		configService,
		appRepository,
		userRepository,
		sessionRepository,
		elasticMailAdapter,
	)

	appHandler := handler.NewAppHandler(appService)
	userHandler := handler.NewUserHandler(userService)
	authMiddleware := middlewares.NewAuthMiddleware(configService)

	app := fiber.New(fiber.Config{Prefork: false})

	app.Use(logs.ApiLogger())
	app.Get("/health", handler.Health)
	app.Get("/templat", userHandler.GetTemplate)

	appRouter := app.Group("/api/app")
	authRouter := app.Group("/api/auth")

	routes.AppRouter(appRouter, appHandler)
	routes.AuthRouter(authRouter, authMiddleware, userHandler)

	err := app.Listen(fmt.Sprintf(":%v", configService.Port))
	if err != nil {
		panic(err)
	}
}
