package database

import (
	"github.com/faagerholm/auth-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=database user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Helsinki"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	DB = db
	db.AutoMigrate(&models.User{})
}
