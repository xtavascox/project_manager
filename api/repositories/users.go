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

	if err := validador.Struct(formattedData); err != nil {
		validationErrors := utils.ValidationErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validationErrors,
		})
	}
	dbResult := database.DB.Create(&formattedData)
	if dbResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": dbResult.Error.Error(),
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
	id := c.Params("id")
	var record model.User
	dbResult := database.DB.Where("id = ?", id).First(&record)
	if dbResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": dbResult.Error.Error(),
		})
	}
	if record.Id == uuid.Nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}
	response := model.UserDto{
		Id:        record.Id,
		FullName:  record.FirstName + " " + record.LastName,
		Email:     record.Email,
		UserName:  record.UserName,
		Birthdate: record.Birthdate,
	}
	context := fiber.Map{
		"statusText": "Ok",
		"data":       response,
	}
	c.Status(200)

	return c.JSON(context)
}
func UsersList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "User List",
	}

	var records []model.User
	database.DB.Find(&records)

	formattedRecords := make([]model.UserDto, len(records))
	for index, record := range records {
		formattedRecords[index] = model.UserDto{
			Id:        record.Id,
			FullName:  record.FirstName + " " + record.LastName,
			Email:     record.Email,
			UserName:  record.UserName,
			Birthdate: record.Birthdate,
		}
	}

	context["data"] = formattedRecords

	c.Status(200)
	return c.JSON(context)

}
func UpdateUser(c *fiber.Ctx) error {
	context := fiber.Map{}

	id := c.Params("id")
	var record model.User
	dbResult := database.DB.Where("id = ?", id).First(&record)
	if dbResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": dbResult.Error.Error(),
		})
	}
	if record.Id == uuid.Nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}
	if err := c.BodyParser(&record); err != nil {
		context["statusText"] = "Error"
		context["msg"] = "Something went wrong. "
		return c.JSON(err)
	}
	result := database.DB.Save(&record)

	if result.Error != nil {
		log.Println("Error in User updating record. ")
		context["msg"] = "Something went wrong. "
		context["statusText"] = "Error"
		return c.JSON(result.Error)
	}
	context["msg"] = "User Updated Successfully"

	c.Status(200)
	return c.JSON(context)
}
func DeleteUser(c *fiber.Ctx) error {
	context := fiber.Map{}
	id := c.Params("id")
	var record model.User
	dbResult := database.DB.Where("id = ?", id).First(&record)
	if dbResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": dbResult.Error.Error(),
		})
	}
	if record.Id == uuid.Nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	result := database.DB.Delete(record)
	if result.Error != nil {
		context["msg"] = "Something went wrong. "
		context["statusText"] = "Error"
		return c.JSON(result.Error)
	}
	context["msg"] = "User Deleted Successfully"

	return c.JSON(context)
}
