package middlewares

import (
	"lazy-platform-auth/config"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

type AuthMiddleware interface {
	TokenMiddleware() fiber.Handler
	RefreshTokenMiddleware() fiber.Handler
}

type authMiddleware struct {
	configService config.Config
}

func NewAuthMiddleware(configService config.Config) AuthMiddleware {
	return authMiddleware{
		configService: configService,
	}
}

func (m authMiddleware) TokenMiddleware() fiber.Handler {
	return JwtMiddleware(m.configService.JwtTokenSecret)
}

func (m authMiddleware) RefreshTokenMiddleware() fiber.Handler {
	return JwtMiddleware(m.configService.JwtRefreshTokenSecret)
}

func JwtMiddleware(jwtSecret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(jwtSecret),
		SuccessHandler: func(c *fiber.Ctx) error {
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			return fiber.ErrUnauthorized
		},
	})
}
