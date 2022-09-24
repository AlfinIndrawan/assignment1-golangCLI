package database

import (
	"fmt"
	"jwt-assignment/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     string
	user     string
	password string
	dbName   string
	dbPort   string
	db       *gorm.DB
)

// take from env
func DBinit() {
	host = os.Getenv("host")
	user = os.Getenv("user")
	password = os.Getenv("password")
	dbName = os.Getenv("dbName")
	dbPort = os.Getenv("dbPort")
}

func StartDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	DBinit()

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ", host, user, password, dbName, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected")
	db.AutoMigrate(&models.User{}, &models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
