package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func ValidateJwt(c *fiber.Ctx) error {
	if c.Path() == "/api/v1/auth/login" || c.Path() == "/api/v1/auth/register" {
		return c.Next()
	}
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	return c.Next()
}

func ValidationErrors(err error) map[string]string {
	errors := err.(validator.ValidationErrors)
	errorMessages := make(map[string]string)
	for _, err := range errors {
		log.Println(err.Field())
		log.Println(err.Value())
		errorMessages[err.Field()] = fmt.Sprintf("%s is %s", err.Field(), err.Tag())
	}
	return errorMessages
}
