package main

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/database/migration"
	"go-fiber-gorm/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseInit()

	migration.RunMigration()


	app := fiber.New()

	route.Router(app)

	app.Listen(":3000")
}