package database

import (
	"sync"

	"gorm.io/gorm"
)

var once sync.Once
var conn *gorm.DB

func Connect() *gorm.DB {
	once.Do(func() {
		conn = GetDBConn()
	})

	return conn
}
