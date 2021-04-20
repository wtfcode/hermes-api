package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type config struct {
	user   string
	pwd    string
	host   string
	port   string
	dbName string
}

func (c *config) GetConnString() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Managua",
		c.host,
		c.user,
		c.pwd,
		c.dbName,
		c.port,
	)
}

func (c *config) SetPropsConfig() {
	c.user = os.Getenv("POSTGRES_USER")
	c.pwd = os.Getenv("POSTGRES_PASSWORD")
	c.host = os.Getenv("POSTGRES_HOST")
	c.port = os.Getenv("POSTGRES_PORT")
	c.dbName = os.Getenv("POSTGRES_DB")
}

func GetDBConn() *gorm.DB {
	var conf config
	conf.SetPropsConfig()
	dsn := conf.GetConnString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}
	return db
}
