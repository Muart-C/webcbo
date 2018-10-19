package routes

import (
	"github.com/Muart-C/webcbo/controllers"
	"github.com/gorilla/mux"
)


//SetProjectRouter method definition
func SetProjectRouter(router *mux.Router) *mux.Router  {
	router.HandleFunc("/projects",controllers.AuthenticationMiddleware(controllers.CreateProjectController)).Methods("POST")
	router.HandleFunc("/projects",controllers.AuthenticationMiddleware(controllers.GetProjectsController)).Methods("GET")
	router.HandleFunc("/projects/{id}",controllers.AuthenticationMiddleware(controllers.GetProjectController)).Methods("GET")
	router.HandleFunc("/projects/{id}", controllers.AuthenticationMiddleware(controllers.UpdateProjectController)).Methods("PUT")
	router.HandleFunc("/projects{id}/milestones",controllers.AuthenticationMiddleware(controllers.CreateMilestoneController)).Methods("POST")
	router.HandleFunc("/milestones",controllers.AuthenticationMiddleware(controllers.GetMilestonesController)).Methods("GET")
	router.HandleFunc("/projects{id}/milestones",controllers.AuthenticationMiddleware(controllers.GetMilestonesInProjectController)).Methods("GET")
	router.HandleFunc("/milestones/{id}",controllers.AuthenticationMiddleware(controllers.GetMilestoneController)).Methods("GET")
	router.HandleFunc("/milestones/{id}", controllers.AuthenticationMiddleware(controllers.UpdateMilestoneController)).Methods("PUT")

	return router
}