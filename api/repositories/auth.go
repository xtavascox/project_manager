package repositories

import (
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/project_management/api/model"
	"github.com/project_management/database"
	"github.com/project_management/utils"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func Login(c *fiber.Ctx) error {
	validador := validator.New()
	var data model.UserLogin

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := validador.Struct(data); err != nil {
		validationErrors := utils.ValidationErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validationErrors,
		})
	}
	var user model.User
	database.DB.Where("email = ?", data.Email).First(&user)
	if user.Id == uuid.Nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if error := bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password)); error != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}
	expireAt := time.Unix(time.Now().Add(time.Minute*24).Unix(), 0)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Id.String(),
		ExpiresAt: jwt.At(expireAt),
	})
	key := os.Getenv("JWT_SECRET")
	token, err := claims.SignedString([]byte(key))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not login",
		})
	}
	cookie := fiber.Cookie{Name: "jwt", Value: token, Expires: expireAt, HTTPOnly: true}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
		"token":   token,
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{Name: "jwt", Value: "", Expires: time.Now().Add(-time.Hour), HTTPOnly: true}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Logout success",
	})
}
