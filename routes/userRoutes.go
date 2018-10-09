package routes

import (
	"github.com/Muart-C/webcbo/controllers"
	"github.com/gorilla/mux"
)

func SetUserRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/users", controllers.CreateUserController).Methods("POST")
	router.HandleFunc("/users", controllers.GetUsersController).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUserController).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUserController).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUserController).Methods("DELETE")

	//Employee routes
	router.HandleFunc("/users/{id}/employees", controllers.CreateEmployeeController).Methods("POST")
	router.HandleFunc("/employees", controllers.GetEmployeesController).Methods("GET")
	router.HandleFunc("/employees/{id}", controllers.GetEmployeeController).Methods("GET")
	router.HandleFunc("/employees/{id}", controllers.UpdateEmployeeController).Methods("PUT")
	router.HandleFunc("/employees/{id}", controllers.DeleteEmployeeController).Methods("DELETE")

	//EmployeeHours routes
	router.HandleFunc("/employees/{id}/hours", controllers.CreateEmployeeHoursController).Methods("POST")
	router.HandleFunc("/hours", controllers.GetEmployeesHoursController).Methods("GET")
	router.HandleFunc("/hours/{id}", controllers.GetEmployeeHoursController).Methods("GET")
	router.HandleFunc("/hours/{id}", controllers.UpdateEmployeeHoursController).Methods("PUT")
	router.HandleFunc("/hours/{id}", controllers.DeleteEmployeeHoursController).Methods("DELETE")

	//Client routes
	router.HandleFunc("/clients", controllers.CreateClientController).Methods("POST")
	router.HandleFunc("/clients", controllers.GetClientsController).Methods("GET")
	router.HandleFunc("/clients/{id}", controllers.GetClientController).Methods("GET")
	router.HandleFunc("/clients/{id}", controllers.UpdateClientController).Methods("PUT")
	router.HandleFunc("/clients/{id}", controllers.DeleteClientController).Methods("DELETE")

	//Role routes
	router.HandleFunc("/roles", controllers.CreateRoleController).Methods("POST")
	router.HandleFunc("/roles", controllers.GetRolesController).Methods("GET")
	router.HandleFunc("/roles/{id}", controllers.GetRoleController).Methods("GET")
	router.HandleFunc("/roles/{id}", controllers.UpdateRoleController).Methods("PUT")
	router.HandleFunc("/roles/{id}", controllers.DeleteRoleController).Methods("DELETE")

	//Team routes
	router.HandleFunc("/teams", controllers.CreateTeamController).Methods("POST")
	router.HandleFunc("/teams", controllers.GetTeamsController).Methods("GET")
	router.HandleFunc("/teams/{id}", controllers.GetTeamController).Methods("GET")
	router.HandleFunc("/teams/{id}", controllers.UpdateTeamController).Methods("PUT")
	router.HandleFunc("/teams/{id}", controllers.DeleteTeamController).Methods("DELETE")
	return router
}
