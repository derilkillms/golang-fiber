package handler

import (
	"golang-fiber/database"
	"golang-fiber/model/entity"
	"golang-fiber/model/request"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerRead(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"data": "user",
	})

}

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(users)

}

func UserHandlerRaw(ctx *fiber.Ctx) error {
	var results []map[string]interface{}
	nganu := database.DB.Raw("SELECT * FROM users WHERE 1=?", 1).Scan(&results)
	if nganu.Error != nil {
		panic(nganu.Error)
	}
	return ctx.JSON(results)
}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	//INITIAL STRUCT INPUT FORM
	user := new(request.UserCreateRequest)

	//PARSING FROM REQUEST FORM
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	//VALIDATING FORM INPUT
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	//INITIAL FROM STRUCT INPUT TO STRUCT FIELD TABLE
	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	//CRETE TO DATABASE
	errCreate := database.DB.Create(&newUser).Error
	if errCreate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})
}
