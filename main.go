package main

import (
	"guisantosalves/apigolang/models"
	"guisantosalves/apigolang/routes"

	"github.com/joho/godotenv"
)

// w -> response / r -> request
func main() {
	// initializing dotenv file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// db config
	models.InitialMigration()

	// routes config
	var e = *routes.E

	// initializing routes
	routes.Routing()

	// starting the server
	e.Start(":3000")
}
