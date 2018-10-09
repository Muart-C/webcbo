package routes

import (
	"github.com/Muart-C/webcbo/controllers"
	"github.com/gorilla/mux"
)

func SetUserRouter(router *mux.Router) *mux.Router  {
	router.HandleFunc("/users",controllers.CreateUserController).Methods("POST")
	router.HandleFunc("/users",controllers.GetUsersController).Methods("GET")
	router.HandleFunc("/users/{id}",controllers.GetUserController).Methods("GET")
	router.HandleFunc("/users/{id}",controllers.UpdateUserController).Methods("PUT")
	router.HandleFunc("/users/{id}",controllers.DeleteUserController).Methods("DELETE")

	//Employee routes
	router.HandleFunc("/users/{id}/employees",controllers.CreateEmployeeController).Methods("POST")
	router.HandleFunc("/employees",controllers.GetEmployeesController).Methods("GET")
	router.HandleFunc("/employees/{eid}",controllers.GetEmployeeController).Methods("GET")
	router.HandleFunc("/employees/{eid}",controllers.UpdateEmployeeController).Methods("PUT")
	router.HandleFunc("/employees/{eid}",controllers.DeleteEmployeeController).Methods("DELETE")
	return router
}