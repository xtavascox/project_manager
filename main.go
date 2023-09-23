package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/project_management/api/controllers"
	"github.com/project_management/database"
	"github.com/project_management/utils"
	"log"
)

func init() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Environment variables loaded successfully")
	database.Connect()
}

func main() {

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Use(utils.ValidateJwt)

	controllers.Login(app)
	controllers.Users(app)

	log.Fatal(app.Listen(":4000"))
}
