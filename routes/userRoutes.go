package routes

import (
	"github.com/Muart-C/webcbo/controllers"
	"github.com/gorilla/mux"
)

func SetUserRouter(router *mux.Router) *mux.Router  {
	router.HandleFunc("/users",controllers.CreateUserController).Methods("GET")
	return router
}