package main

import (
	"guisantosalves/apigolang/models"
	"guisantosalves/apigolang/routes"
)

// w -> response / r -> request
func main() {
	// db config
	models.InitialMigration()

	// routes config
	var e = *routes.E

	// initializing routes
	routes.Routing()

	// starting the server
	e.Start(":3000")
}
