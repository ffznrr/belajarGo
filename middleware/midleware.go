package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func UserMiddleware(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	fmt.Println(token,"ini token")
	if token == "" || token != "secret" {
		return ctx.Status(403).JSON(fiber.Map{
			"message": "unaothorized",
		})
	}

	return ctx.Next()
}