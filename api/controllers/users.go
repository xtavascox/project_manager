package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/project_management/api/repositories"
)

func Users(app *fiber.App) {
	user := app.Group("/api/v1/user")

	user.Get("/", repositories.UsersList)
	user.Get("/:id", repositories.UserById)
	user.Post("/register", repositories.RegisterUser)
	user.Put("/:id", repositories.UpdateUser)
	user.Delete("/:id", repositories.DeleteUser)
}
