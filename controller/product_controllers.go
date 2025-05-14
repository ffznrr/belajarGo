package controller

import (
	"fmt"
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetProducts(ctx *fiber.Ctx)error{
	var product []entity.Product

	result := database.DB.Find(&product)

	if result.Error != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": product,
		"woi": "ok",
	})

}

func CreateProduct(ctx *fiber.Ctx)error{
	product := new(request.ProductCreateRequest)

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validate := validator.New()

	errValidate := validate.Struct(product)

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

	newProduct := entity.Product{
		Name: product.Name,
		Stock: product.Stock,
		Price: product.Price,
		
	}

	result := database.DB.Create(&newProduct)

	if result.Error != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"Message": product,
	})
}
