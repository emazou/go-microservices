package main

import (
	"user-service/config"
	"user-service/routes"
)

func main() {
	// Connect to the database
	config.ConnectDatabase()
	// Set up the routes
	r := routes.SetupRouter()
	// Run the server
	r.Run(":8080")
}
