package handler

import (
	"golang-fiber/database"
	"golang-fiber/model/entity"
	"golang-fiber/model/request"
	"golang-fiber/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)

	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
	}

	//validasi request
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	//check available user
	var user entity.User
	err := database.DB.First(&user, "email=?", loginRequest.Email).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	//check validation password
	isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}
	//generate jwt
	claims := jwt.MapClaims{}

	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["address"] = user.Address
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix() // set the expiration time for the token
	claims["iat"] = time.Now().Unix()                      // set the time the token was issued
	claims["role"] = "user"
	if user.Email == "seizuro@gmail.com" {
		claims["role"] = "admin"
	}
	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "error generate token",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})

}
