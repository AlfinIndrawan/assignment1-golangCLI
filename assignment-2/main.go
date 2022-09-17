package main

import (
	"assignment-2/config"
	"assignment-2/controllers"
	"assignment-2/repository"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func getEnv() *gorm.DB {
	dbUsername := os.Getenv("DB_USERNAME")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	db := config.DBInit(dbUsername, dbPassword, dbPort, dbName, dbHost)
	return db
}
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	serverAddress := os.Getenv("SERVICE_PORT")
	db := getEnv()
	Order := repository.NewRepository(db)
	inDB := &controllers.InDB{
		Repository: Order,
	}

	router := gin.Default()

	router.GET("/orders", inDB.GetOrders)
	router.POST("/orders", inDB.CreateOrder)
	router.GET("/orders/id", inDB.GetOrder)
	router.PUT("/orders/:id", inDB.UpdateOrder)
	router.DELETE("/orders/:id", inDB.DeleteOrder)
	router.Run(serverAddress)
}
