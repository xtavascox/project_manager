package repositories

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/project_management/api/model"
	"github.com/project_management/database"
	"github.com/project_management/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
	"time"
)

func RegisterUser(c *fiber.Ctx) error {
	validador := validator.New()
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	log.Println(data)
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	birthdate, _ := strconv.ParseInt(data["birthdate"], 10, 64)

	formattedData := model.User{
		Id:         uuid.New(),
		FirstName:  data["first_name"],
		LastName:   data["last_name"],
		Email:      data["email"],
		Password:   password,
		UserName:   data["user_name"],
		Birthdate:  birthdate,
		CratedAt:   time.Now().Unix(),
		ModifiedAt: time.Now().Unix(),
	}
	log.Println(formattedData)
	if err := validador.Struct(formattedData); err != nil {
		validationErrors := utils.ValidationErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validationErrors,
		})
	}
	db_result := database.DB.Create(&formattedData)
	if db_result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": db_result.Error.Error(),
		})
	}
	response := model.UserDto{
		Id:        formattedData.Id,
		FullName:  formattedData.FirstName + " " + formattedData.LastName,
		Email:     formattedData.Email,
		UserName:  formattedData.UserName,
		Birthdate: formattedData.Birthdate,
	}
	context := fiber.Map{
		"message": "User created successfully",
		"data":    response,
	}
	c.Status(fiber.StatusCreated)
	return c.JSON(context)
}
func UserById(c *fiber.Ctx) error {
	return nil
}
func UsersList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "BlogList",
	}

	var records []model.User
	database.DB.Find(&records)

	context["data"] = records

	c.Status(200)
	return c.JSON(context)

}
func UpdateUser(c *fiber.Ctx) error {
	return nil
}
func DeleteUser(c *fiber.Ctx) error {
	return nil
}
