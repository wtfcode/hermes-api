package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wtfcode/hermes-api/pkg/controllers"
	"gorm.io/gorm"
)

func SetUserRouter(r fiber.Router, db *gorm.DB) {
	controller := controllers.NewUserController(db)
	r.Post("/", controller.Save())
	r.Get("/all/:page", controller.GetAll())
	r.Get("/:id", controller.GetOne())
	r.Delete("/:id", controller.Delete())
	r.Put("/:id", controller.Update())
}
