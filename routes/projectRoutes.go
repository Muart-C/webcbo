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
	router.HandleFunc("/projectStatus",controllers.AuthenticationMiddleware(controllers.CreateProjectStatusController)).Methods("POST")
	router.HandleFunc("/projectStatus/{id}", controllers.AuthenticationMiddleware(controllers.UpdateProjectStatusController)).Methods("PUT")
	router.HandleFunc("/milestones",controllers.AuthenticationMiddleware(controllers.CreateMilestoneController)).Methods("POST")
	router.HandleFunc("/milestones/{id}", controllers.AuthenticationMiddleware(controllers.UpdateMilestoneController)).Methods("PUT")
	router.HandleFunc("/milestoneStatus",controllers.AuthenticationMiddleware(controllers.CreateMilestoneStatusController)).Methods("POST")
	router.HandleFunc("/milestoneStatus/{id}", controllers.AuthenticationMiddleware(controllers.UpdateMilestoneStatusController)).Methods("PUT")

	return router
}