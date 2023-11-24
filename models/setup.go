package models

import (
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=db user=postgres password=password dbname=bookstore_go_development port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Faild to conect to database!")
	}

	err = database.AutoMigrate(&Book{})

	if err != nil {
		return
	}

	DB = database
}
