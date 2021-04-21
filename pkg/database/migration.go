package database

import (
	"github.com/wtfcode/hermes-api/pkg/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
