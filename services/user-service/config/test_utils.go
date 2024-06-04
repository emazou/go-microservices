package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"user-service/models"
)

func ConnectTestDatabase() {
	err := godotenv.Load(filepath.Join("..", ".env.test"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dns := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v",
		os.Getenv("TEST_DB_USER"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_HOST"),
		os.Getenv("TEST_DB_PORT"),
		os.Getenv("TEST_DB_NAME"),
	)
	fmt.Println(dns)
	database, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}
	database.AutoMigrate(&models.User{})
	DB = database
}

func TeardownTestDatabase() {
	err := DB.Migrator().DropTable(&models.User{})
	if err != nil {
		log.Fatalf("Failed to drop table: %v", err)
	}
	fmt.Println("Test database teardown completed")
}
