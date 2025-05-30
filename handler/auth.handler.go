package handler

import (
	"fmt"
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"
	"go-fiber-gorm/model/response"
	"go-fiber-gorm/utils"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx)error {
	userResponse := new(request.UserLoginRequest)

	var user entity.User
	errLogin := ctx.BodyParser(userResponse)
	

	if errLogin != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message" : errLogin.Error(),
		})
	}

	checkName := database.DB.First(&user, "Name = ?",userResponse.Name ).Error
	

	if checkName != nil {
		return ctx.Status(fiber.ErrNotFound.Code).JSON(fiber.Map{
			"message": "Username/Password Salah",
		})
	}

	// validasi password

	isValid :=  utils.CheckPasswordHash(userResponse.Password, user.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong Credential",
		})
	}


	return ctx.Status(200).JSON(fiber.Map{
		"message": "Login Success",
	})
}


func Register(ctx *fiber.Ctx)error{
	user := &request.UserCreateRequest{}
	if err := ctx.BodyParser(user); err != nil {
		return err
	}
	fmt.Println(user, "ini user")

	validate := validator.New()

	errValidate := validate.Struct(user)

	if errValidate != nil {
		validationErrors := errValidate.(validator.ValidationErrors)
		var errors []string

		for _, err := range validationErrors {
			errors = append(errors, fmt.Sprintf("Field %s: %s", err.Field(), err.Tag()))
		} 
		return ctx.Status(400).JSON(fiber.Map{
			"error": errors,
		})
	}


	

	newUser := entity.User{
		Name: user.Name,
		Password: user.Password,
		Email: user.Email,
		Address: user.Address,
		Phone: user.Phone,
	}

	hashedPass, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	newUser.Password = hashedPass

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	userResponse := response.RegisterResponse{
		Name: newUser.Name,
		Email: newUser.Email,
		Address: newUser.Address,
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Success Create User",
		"data": userResponse,
	})

}