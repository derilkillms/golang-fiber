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
	nganu := database.DB.Raw("SELECT id, name, email, address, phone, created_at, updated_at FROM users WHERE 1=?", 1).Scan(&results)
	if nganu.Error != nil {
		panic(nganu.Error)
	}
	return ctx.JSON(results)
}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	//INITIAL STRUCT REQUEST USER
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

	//INITIAL FROM STRUCT REQUEST TO STRUCT FIELD TABLE
	newUser := entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Address:  user.Address,
		Phone:    user.Phone,
		Password: user.Password,
	}

	//CREATE TO DATABASE
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

func UserHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User
	err := database.DB.First(&user, "id=?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// userResponse := response.UserResponse{
	// 	ID:        user.ID,
	// 	Name:      user.Name,
	// 	Email:     user.Email,
	// 	Address:   user.Address,
	// 	Phone:     user.Phone,
	// 	CreatedAt: user.CreatedAt,
	// 	UpdatedAt: user.UpdatedAt,
	// }

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})

}

// update data by id
func UserHandlerUpdate(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)
	//PARSING FROM REQUEST FORM
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user entity.User
	userId := ctx.Params("id")
	//CHECK AVAILABLE
	err := database.DB.First(&user, "id=?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})

	}

	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}

	user.Address = userRequest.Address
	user.Phone = userRequest.Phone

	errUpdate := database.DB.Save(&user).Error

	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	//UPDATE DATA
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

// update email
func UserHandlerUpdateEmail(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUEmailRequest)
	//PARSING FROM REQUEST FORM
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	//VALIDATING FORM INPUT
	validate := validator.New()
	errValidate := validate.Struct(userRequest)
	if errValidate != nil {
		return ctx.JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var user entity.User
	userId := ctx.Params("id")
	//CHECK AVAILABLE
	err := database.DB.First(&user, "id=?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})

	}

	// CHECK AVAILABLE EMAIL
	errCheckEmail := database.DB.First(&user, "email = ?", userRequest.Email).Error

	if errCheckEmail == nil {
		return ctx.Status(402).JSON(fiber.Map{
			"message": "user already use",
		})
	}

	// update user Data

	user.Email = userRequest.Email

	errUpdate := database.DB.Save(&user).Error

	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	//UPDATE DATA
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserHandlerDelete(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User

	//AVAILABLE ID
	err := database.DB.Debug().First(&user, "id=?", userId).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&user).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})

	}
	return ctx.JSON(fiber.Map{
		"message": "user was deleted",
	})
}
