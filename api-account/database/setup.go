package database

import (
	"api-account/models"

	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() (err error) {

	HOST := "localhost"
	PORT := "5432"
	DB_USER := "postgres"
	DB_PASS := "postgres"
	DB_NAME := "fp-account"

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		HOST, DB_USER, DB_PASS, DB_NAME, PORT)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.Account{})

	if err != nil {
		panic(err)
	}

	return nil
}

func Get() *gorm.DB {
	return db
}
