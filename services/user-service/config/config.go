package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"user-service/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dns := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	database, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}
	database.AutoMigrate(&models.User{})
	DB = database
}
