package handler

import (
	"golang-fiber/database"
	"golang-fiber/model/entity"
	"log"

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

// func UserHandlerCreate(ctx *fiber.Ctx) {
// 	return ctx.SendString("Hello, World!")
// }
