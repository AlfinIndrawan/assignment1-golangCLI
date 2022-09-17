package config

import (
	"assignment-2/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInit(dbUsername, dbPassword, dbPort, dbName, dbHost string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(model.Order{})
	db.AutoMigrate(model.Item{})
	return db
}
