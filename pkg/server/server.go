package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wtfcode/hermes-api/pkg/router"
	"gorm.io/gorm"
)

func NewServer(db *gorm.DB) *fiber.App {
	app := fiber.New()
	router.SetUserRouter(app.Group("/users"), db)

	return app
}
