package route

import (
	"go-fiber-gorm/config"
	"go-fiber-gorm/controller"
	"go-fiber-gorm/handler"
	"go-fiber-gorm/middleware"

	"github.com/gofiber/fiber/v2"
)


func Router(app *fiber.App) {
	app.Static("/public",config.ProjectRootPath+"./public/asset")
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)


app.Get("/main", controller.UserHandlerCreate )
app.Use(middleware.UserMiddleware)
app.Get("/user", controller.GetUser)
app.Get("/user/:id", controller.GetUserId)
app.Get("/product", controller.GetProducts)
app.Post("/createproduct", controller.CreateProduct)
app.Put("/updateuser/:id", controller.UpdateUser)
app.Put("/updateuseremail/:id", controller.UpdateUserEmail)
app.Delete("/deleteuser/:id", controller.DeleteUser)
}