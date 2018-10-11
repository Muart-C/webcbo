package routes

import (
	"github.com/Muart-C/webcbo/controllers"
	"github.com/gorilla/mux"
)

func SetProjectRouter(router *mux.Router) *mux.Router  {
	router.HandleFunc("/projects",controllers.CreateProjectController).Methods("POST")
	router.HandleFunc("/projects",controllers.GetProjectsController).Methods("GET")
	router.HandleFunc("/projects/{id}",controllers.GetProjectController).Methods("GET")
	router.HandleFunc("/projects/{id}", controllers.UpdateProjectController).Methods("PUT")
	return router
}