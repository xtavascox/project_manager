package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/project_management/api/repositories"
)

func Login(app *fiber.App) {
	auth := app.Group("api/v1/auth")

	auth.Post("/login", repositories.Login)
	auth.Post("/logout", repositories.Logout)
}
