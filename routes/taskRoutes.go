package routes

import (
	"github.com/Muart-C/webcbo/controllers"
	"github.com/gorilla/mux"
)

func SetTaskRouter(router *mux.Router) *mux.Router  {

	router.HandleFunc("/projects/{id}/tasks",controllers.AuthenticationMiddleware(controllers.CreateTaskController)).Methods("POST")
	router.HandleFunc("/tasks",controllers.AuthenticationMiddleware(controllers.GetTasksController)).Methods("GET")
	router.HandleFunc("/tasks/{id}",controllers.AuthenticationMiddleware(controllers.GetTaskController)).Methods("GET")
	router.HandleFunc("/tasks/{id}", controllers.AuthenticationMiddleware(controllers.UpdateTaskController	)).Methods("PUT")
	router.HandleFunc("/tasks/{id}/activities",controllers.AuthenticationMiddleware(controllers.CreateActivityController)).Methods("POST")
	router.HandleFunc("/activities/{id}", controllers.AuthenticationMiddleware(controllers.UpdateActivityController)).Methods("PUT")
	router.HandleFunc("/taskStatus",controllers.AuthenticationMiddleware(controllers.CreateTaskStatusController)).Methods("POST")
	router.HandleFunc("/taskStatus/{id}", controllers.AuthenticationMiddleware(controllers.UpdateTaskStatusController)).Methods("PUT")
	router.HandleFunc("/activityStatus",controllers.AuthenticationMiddleware(controllers.CreateActivityStatusController)).Methods("POST")
	router.HandleFunc("/activityStatus/{id}", controllers.AuthenticationMiddleware(controllers.UpdateActivityStatusController)).Methods("PUT")
	return router
}