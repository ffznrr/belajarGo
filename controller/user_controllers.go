package controller

import (
	"fmt"
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerCreate(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{
			"hello": "from contor",
		})

}

func GetUser(ctx *fiber.Ctx)error{
	var user []entity.User

	err := database.DB.Find(&user)

	if err != nil {
		fmt.Errorf("Error Co")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"meessage": user,
	})
}

func CreateUser(c *fiber.Ctx)error{
	user := new(request.UserCreateRequest)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	fmt.Println(user, 2)

	validate := validator.New()

	errValidate := validate.Struct(user)

	if errValidate != nil {
		validationErrors := errValidate.(validator.ValidationErrors)
		var errors []string

		for _, err := range validationErrors {
			errors = append(errors, fmt.Sprintf("Field %s: %s", err.Field(), err.Tag()))
		} 
		return c.Status(400).JSON(fiber.Map{
			"error": errors,
		})
	}

	

	newUser := entity.User{
		Name: user.Name,
		Email: user.Email,
		Address: user.Address,
		Phone: user.Phone,
	}
	

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success Create User",
		"data": newUser,
	})

}

func UpdateUser(c *fiber.Ctx)error{
	id := c.Params("id")
	var user1 entity.User

	errFindOneUser := database.DB.First(&user1, id)

	if errFindOneUser.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Data Not Found",
		})
	}

	user := new(request.UserCreateRequest)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	updateData := map[string]interface{}{
        "name":    user.Name,
        "phone":   user.Phone,
        "email":   user.Email,
        "address": user.Address,
    }

	 result := database.DB.Model(&entity.User{}).Where("id = ?", id).Updates(updateData)
	  if result.Error != nil {
        return c.Status(500).JSON(fiber.Map{
            "message": "Failed to update user",
            "error":   result.Error.Error(),
        })
    }

	return c.Status(200).JSON(fiber.Map{
        "message": "User updated successfully",
    })
}

func UpdateUserEmail(c *fiber.Ctx)error{
	userRequest := new(request.UserEmailRequest)
	id := c.Params("id")
	var user1 entity.User
	var isEmail entity.User

	errFindOneUser := database.DB.First(&user1, id)
	
	if err := c.BodyParser(&userRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if errFindOneUser.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Data Not Found",
		})
	}

	errCheckEmail := database.DB.First(&isEmail, "email = ?", userRequest.Email).Error
	if errCheckEmail == nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Email Sudah Digunakan",
		})
	}


	user1.Email = userRequest.Email

	errUpdate := database.DB.Save(&user1).Error

	if errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": user1,
	})
}

func GetUserId(c *fiber.Ctx)error{
	id := c.Params("id")
	var user entity.User

	result := database.DB.First(&user, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message" : "Data Not Found",
		})
	}

	userResponse := request.UserCreateRequest{
		Name: user.Name,
		Address: user.Address,
		Phone: user.Phone,
		Email: user.Email,
	}

	return c.Status(200).JSON(fiber.Map{
		"Message" : userResponse,
	})
}

func DeleteUser(c *fiber.Ctx)error{
	id := c.Params("id")

	var user entity.User

	err := database.DB.Debug().Find(&user, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"Message": "Data Not Found",
		})
	}

	errVal := database.DB.Debug().Delete(&user,"id = ?", id).Error
	if errVal != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "INTERNAL SERVER ERROR",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "User Success Delete",
	})
}