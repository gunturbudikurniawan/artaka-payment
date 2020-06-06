package adapter

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var database *gorm.DB

func Connect() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=ihsan dbname=artaka sslmode=disable")
	if err != nil {
		panic(err)
	}

	database = db

	return database
}

func GetDB() *gorm.DB {
	return database
}
