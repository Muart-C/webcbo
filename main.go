package main

import (
	"fmt"
	"github.com/Muart-C/webcbo/routes"
	"net/http"
)

func main() {

	router := routes.InitializeRoutes()

	fmt.Println("server is running")
	http.ListenAndServe(":8000",router)

}
