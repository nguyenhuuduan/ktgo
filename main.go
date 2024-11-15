package main

import (
	"customer-management/config"
	"customer-management/routes"
)

func main() {
	config.ConnectDB()
	router := routes.SetupRoutes()
	router.Run(":8080") // Run on localhost:8080
}
