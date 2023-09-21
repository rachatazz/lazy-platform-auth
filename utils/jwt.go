package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(id string, expiredInTime int64, secret string) (string, error) {
	cliamsToken := jwt.StandardClaims{
		Subject:   id,
		ExpiresAt: int64(expiredInTime),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, cliamsToken)

	return jwtToken.SignedString([]byte(secret))
}

func ValidateToken(tokenString string, secret string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(
				fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]),
			)
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	}

	return nil, false
}

func CliamsToken(c *fiber.Ctx) string {
	userToken := c.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	return claims["sub"].(string)
}
