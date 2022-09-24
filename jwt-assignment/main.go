package main

import (
	"jwt-assignment/database"
	"jwt-assignment/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
