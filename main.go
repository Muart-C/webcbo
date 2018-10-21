package main

import (

	"github.com/Muart-C/webcbo/models"
	"github.com/Muart-C/webcbo/routes"
	"github.com/Muart-C/webcbo/views"

)

func main() {
	// Connect to database
	db := models.Connect()
	defer db.Close()

	// Setup templates
	views.LoadTemplates()

	// Get started
	routes.ServeRouter()
}
