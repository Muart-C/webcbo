package routes

import (
	"github.com/Muart-C/webcbo/controllers"
	"github.com/gorilla/mux"
)

//SetUserRouter method definition
func SetUserRouter(router *mux.Router) *mux.Router {

	//Authentication get token
	router.HandleFunc("/auth", controllers.GetAuthToken).Methods("POST")

	router.HandleFunc("/users", controllers.CreateUserController).Methods("POST")
	router.HandleFunc("/users", controllers.AuthenticationMiddleware(controllers.GetUsersController)).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.AuthenticationMiddleware(controllers.GetUserController)).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.AuthenticationMiddleware(controllers.UpdateUserController)).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.AuthenticationMiddleware(controllers.DeleteUserController)).Methods("DELETE")


	//Project Manager routes
	router.HandleFunc("/users/{id}/managers", controllers.AuthenticationMiddleware(controllers.CreateProjectManagerController)).Methods("POST")
	router.HandleFunc("/managers", controllers.AuthenticationMiddleware(controllers.GetProjectManagersController)).Methods("GET")
	router.HandleFunc("/managers/{id}", controllers.AuthenticationMiddleware(controllers.GetProjectManagerController)).Methods("GET")
	//router.HandleFunc("/managers/{id}", controllers.UpdateProjectManagerController).Methods("PUT")//	DEBUG

	//Employee routes
	router.HandleFunc("/users/{id}/employees", controllers.AuthenticationMiddleware(controllers.CreateEmployeeController)).Methods("POST")
	router.HandleFunc("/employees", controllers.AuthenticationMiddleware(controllers.GetEmployeesController)).Methods("GET")
	router.HandleFunc("/employees/{id}", controllers.AuthenticationMiddleware(controllers.GetEmployeeController)).Methods("GET")
	router.HandleFunc("/employees/{id}", controllers.AuthenticationMiddleware(controllers.UpdateEmployeeController)).Methods("PUT")


	//EmployeeHours routes
	router.HandleFunc("/employees/{id}/hours", controllers.AuthenticationMiddleware(controllers.CreateEmployeeHoursController)).Methods("POST")
	router.HandleFunc("/hours", controllers.AuthenticationMiddleware(controllers.GetEmployeesHoursController)).Methods("GET")
	router.HandleFunc("/hours/{id}", controllers.AuthenticationMiddleware(controllers.GetEmployeeHoursController)).Methods("GET")
	router.HandleFunc("/hours/{id}", controllers.AuthenticationMiddleware(controllers.UpdateEmployeeHoursController)).Methods("PUT")


	//Assign routes
	router.HandleFunc("/employees/{id}/assignments", controllers.AuthenticationMiddleware(controllers.CreateAssignmentController)).Methods("POST")
	router.HandleFunc("/assignments", controllers.AuthenticationMiddleware(controllers.GetAssignmentsController)).Methods("GET")
	router.HandleFunc("/assignments/{id}", controllers.AuthenticationMiddleware(controllers.GetAssignmentController)).Methods("GET")
	router.HandleFunc("/assignments/{id}", controllers.AuthenticationMiddleware(controllers.UpdateAssignmentController)).Methods("PUT")


	//Client routes
	router.HandleFunc("/clients", controllers.AuthenticationMiddleware(controllers.CreateClientController)).Methods("POST")
	router.HandleFunc("/clients", controllers.AuthenticationMiddleware(controllers.GetClientsController)).Methods("GET")
	router.HandleFunc("/clients/{id}", controllers.AuthenticationMiddleware(controllers.GetClientController)).Methods("GET")
	router.HandleFunc("/clients/{id}", controllers.AuthenticationMiddleware(controllers.UpdateClientController)).Methods("PUT")
	router.HandleFunc("/clients/{id}", controllers.AuthenticationMiddleware(controllers.DeleteClientController)).Methods("DELETE")

	//Role routes
	router.HandleFunc("/roles", controllers.AuthenticationMiddleware(controllers.CreateRoleController)).Methods("POST")
	router.HandleFunc("/roles", controllers.AuthenticationMiddleware(controllers.GetRolesController)).Methods("GET")
	router.HandleFunc("/roles/{id}", controllers.AuthenticationMiddleware(controllers.GetRoleController)).Methods("GET")
	router.HandleFunc("/roles/{id}", controllers.AuthenticationMiddleware(controllers.UpdateRoleController)).Methods("PUT")
	router.HandleFunc("/roles/{id}", controllers.AuthenticationMiddleware(controllers.DeleteRoleController)).Methods("DELETE")

	//Team routes
	router.HandleFunc("/teams", controllers.AuthenticationMiddleware(controllers.CreateTeamController)).Methods("POST")
	router.HandleFunc("/teams", controllers.AuthenticationMiddleware(controllers.GetTeamsController)).Methods("GET")
	router.HandleFunc("/teams/{id}", controllers.AuthenticationMiddleware(controllers.GetTeamController)).Methods("GET")
	router.HandleFunc("/teams/{id}", controllers.AuthenticationMiddleware(controllers.UpdateTeamController)).Methods("PUT")
	router.HandleFunc("/teams/{id}", controllers.AuthenticationMiddleware(controllers.DeleteTeamController)).Methods("DELETE")
	return router
}

