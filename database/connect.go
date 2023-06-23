package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hassanrefaat9/config"
	"github.com/hassanrefaat9/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic(err)
	}
	DB, err = gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"),
		port,
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
	)), &gorm.Config{})

	if err != nil {
		panic("faild to connect to database")
	}

	log.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.Product{})
	log.Println("Database Migrated")
}
