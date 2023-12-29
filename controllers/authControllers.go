package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/project/project-skripsi/go-be/database"
	"github.com/project/project-skripsi/go-be/models"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

// nama func selalu diawali huruf besar
func Register(c *fiber.Ctx) error {
	var data map[string]string

	// throw error
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// password generate random
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	// User
	users := models.Users{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&users)

	return c.JSON(users)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	// throw error
	if err := c.BodyParser(&data); err != nil {
		return err

	}

	var users models.Users

	database.DB.Where("email = ?", data["email"]).First(&users)

	if users.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"Message": "User Not Found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(users.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(users.Id)),
		ExpiresAt: expirationTime.Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  expirationTime,
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success Login",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.Users

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(claims)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Succes Logout",
	})
}
