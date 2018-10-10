package controllers

import (
	"encoding/json"
	"github.com/Muart-C/webcbo/models"
	"github.com/Muart-C/webcbo/repository"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//CreateUserController controller definition
func CreateUserController(w http.ResponseWriter, r *http.Request) {
	var user models.User
	//Decode the incoming user json data
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		RespondWithError(w, 500, "Invalid user data on the body")
		return
	}
	//save a new user

	response, err := repository.CreateUser(user.UserName, user.Email, user.LastName, user.FirstName)

	if err != nil {
		RespondWithError(w, 500, "An unexpected error occured")
	}
	RespondWithJSON(w, http.StatusCreated, response)
}

//GetUsersController controller definition
func GetUsersController(w http.ResponseWriter, r *http.Request) {
	users, err := repository.FetchUsers()
	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
	}
	RespondWithJSON(w, http.StatusFound, users)
}

//GetUserController controller definition
func GetUserController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err == nil {
		user, err := repository.FetchUser(userID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "User not found")
			return
		}
		RespondWithJSON(w, http.StatusFound, user)
	}
}

//UpdateUserController controller definition
func UpdateUserController(w http.ResponseWriter, r *http.Request) {

	var user models.User
	//Decode the incoming user json data
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		RespondWithError(w, 500, "Invalid user data on the body")
		return
	}
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err == nil {
		update, err := repository.UpdateUser(user.UserName, user.Email, user.FirstName, user.LastName, userID)
		if err != nil {
			RespondWithError(w, http.StatusNotModified, "Error updating user")
			return
		}
		RespondWithJSON(w, http.StatusCreated, update)
	}
}

//DeleteUserController controller definition
func DeleteUserController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err == nil {
		err := repository.DeleteUser(userID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "User not found")
			return
		}
		RespondWithJSON(w, http.StatusFound, "User Deleted successfully")
	}
}

//CreateEmployeeController controller definition
func CreateEmployeeController(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err == nil {
		user, err := repository.FetchUser(userID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "User not found register him/her as a user first")
			return
		}
		er := json.NewDecoder(r.Body).Decode(&employee)
		if er != nil {
			RespondWithError(w, 500, "Invalid employee payload")
			return
		}
		//save a new user

		response, err := repository.CreateEmployee(*user, employee.EmployeeProfession)

		if err != nil {
			RespondWithError(w, 500, "An unexpected error occured")
		}
		RespondWithJSON(w, http.StatusCreated, response)
	}
}

//GetEmployeesController controller definition
func GetEmployeesController(w http.ResponseWriter, r *http.Request) {
	employees, err := repository.FetchEmployees()
	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
	}
	RespondWithJSON(w, http.StatusFound, employees)
}

//GetEmployeeController controller definition
func GetEmployeeController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	employeeID, err := strconv.Atoi(params["id"])
	if err == nil {
		employee, err := repository.FetchEmployee(employeeID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "User not found")
			return
		}
		RespondWithJSON(w, http.StatusFound, employee)
	}
}

//UpdateEmployeeController controller definition
func UpdateEmployeeController(w http.ResponseWriter, r *http.Request) {

	var employee models.Employee
	//Decode the incoming user json data
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		RespondWithError(w, 500, "Invalid user data on the body")
		return
	}
	params := mux.Vars(r)
	employeeID, err := strconv.Atoi(params["id"])
	if err == nil {
		update, err := repository.UpdateEmployee(employee.EmployeeProfession, employeeID)
		if err != nil {
			RespondWithError(w, http.StatusNotModified, "Error updating user")
			return
		}
		RespondWithJSON(w, http.StatusCreated, update)
	}
}

//CreateEmployeeHoursController controller definition
func CreateEmployeeHoursController(w http.ResponseWriter, r *http.Request) {
	var hours models.Hours
	params := mux.Vars(r)
	employeeID, err := strconv.Atoi(params["id"])
	if err == nil {
		employee, err := repository.FetchEmployee(employeeID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "employee not found ")
			return
		}
		er := json.NewDecoder(r.Body).Decode(&employee)
		if er != nil {
			RespondWithError(w, 500, "Invalid employee payload")
			return
		}
		//save a new hours

		response, err := repository.CreateEmployeeHours(*employee, hours.AssignedOn, hours.AssignedAt, hours.HoursCompleted, hours.WorkCompleted)

		if err != nil {
			RespondWithError(w, 500, "An unexpected error occured")
		}
		RespondWithJSON(w, http.StatusCreated, response)
	}
}

//GetEmployeesHoursController controller definition
func GetEmployeesHoursController(w http.ResponseWriter, r *http.Request) {
	hours, err := repository.FetchEmployeesHours()
	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
	}
	RespondWithJSON(w, http.StatusFound, hours)
}

//GetEmployeeHoursController controller definition
func GetEmployeeHoursController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	employeeID, err := strconv.Atoi(params["id"])
	if err == nil {
		employee, err := repository.FetchEmployeeHours(employeeID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "hours not found")
			return
		}
		RespondWithJSON(w, http.StatusFound, employee)
	}
}

//UpdateEmployeeHoursController controller definition
func UpdateEmployeeHoursController(w http.ResponseWriter, r *http.Request) {

	var hours models.Hours
	//Decode the incoming hours json data
	err := json.NewDecoder(r.Body).Decode(&hours)
	if err != nil {
		RespondWithError(w, 500, "Invalid hours data on the body")
		return
	}
	params := mux.Vars(r)
	employeeID, err := strconv.Atoi(params["id"])
	if err == nil {
		update, err := repository.UpdateEmployeeHours(hours.AssignedOn, hours.AssignedAt, hours.HoursCompleted, hours.WorkCompleted, employeeID)
		if err != nil {
			RespondWithError(w, http.StatusNotModified, "Error updating employee hours")
			return
		}
		RespondWithJSON(w, http.StatusCreated, update)
	}
}

//CreateRoleController controller definition
func CreateRoleController(w http.ResponseWriter, r *http.Request) {
	var role models.Role
	//Decode the incoming role json data
	err := json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		RespondWithError(w, 500, "Invalid role data on the body")
		return
	}
	//save a new role

	response, err := repository.CreateRole(role.RoleName)

	if err != nil {
		RespondWithError(w, 500, "An unexpected error occured")
	}
	RespondWithJSON(w, http.StatusCreated, response)
}

//GetRolesController controller definition
func GetRolesController(w http.ResponseWriter, r *http.Request) {
	roles, err := repository.FetchRoles()
	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
	}
	RespondWithJSON(w, http.StatusFound, roles)
}

//GetRoleController controller definition
func GetRoleController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	roleID, err := strconv.Atoi(params["id"])
	if err == nil {
		role, err := repository.FetchRole(roleID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "role not found")
			return
		}
		RespondWithJSON(w, http.StatusFound, role)
	}
}

//UpdateRoleController controller definition
func UpdateRoleController(w http.ResponseWriter, r *http.Request) {

	var role models.Role
	//Decode the incoming role json data
	err := json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		RespondWithError(w, 500, "Invalid role data on the body")
		return
	}
	params := mux.Vars(r)
	roleID, err := strconv.Atoi(params["id"])
	if err == nil {
		update, err := repository.UpdateRole(role.RoleName, roleID)
		if err != nil {
			RespondWithError(w, http.StatusNotModified, "Error updating role")
			return
		}
		RespondWithJSON(w, http.StatusCreated, update)
	}
}

//DeleteRoleController controller definition
func DeleteRoleController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	roleID, err := strconv.Atoi(params["id"])
	if err == nil {
		err := repository.DeleteRole(roleID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "role not found")
			return
		}
		RespondWithJSON(w, http.StatusFound, "role Deleted successfully")
	}
}

//CreateTeamController controller definition
func CreateTeamController(w http.ResponseWriter, r *http.Request) {
	var team models.Team
	//Decode the incoming Team json data
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		RespondWithError(w, 500, "Invalid Team data on the body")
		return
	}
	//save a new Team

	response, err := repository.CreateTeam(team.TeamName)

	if err != nil {
		RespondWithError(w, 500, "An unexpected error occured")
	}
	RespondWithJSON(w, http.StatusCreated, response)
}

//GetTeamsController controller definition
func GetTeamsController(w http.ResponseWriter, r *http.Request) {
	teams, err := repository.FetchTeams()
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "No teams found")
	}
	RespondWithJSON(w, http.StatusFound, teams)
}

//GetTeamController controller definition
func GetTeamController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	teamID, err := strconv.Atoi(params["id"])
	if err == nil {
		team, err := repository.FetchTeam(teamID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "Team not found")
			return
		}
		RespondWithJSON(w, http.StatusFound, team)
	}
}

//UpdateTeamController controller definition
func UpdateTeamController(w http.ResponseWriter, r *http.Request) {

	var team models.Team
	//Decode the incoming Team json data
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		RespondWithError(w, 500, "Invalid Team data on the body")
		return
	}
	params := mux.Vars(r)
	teamID, err := strconv.Atoi(params["id"])
	if err == nil {
		update, err := repository.UpdateTeam(team.TeamName, teamID)
		if err != nil {
			RespondWithError(w, http.StatusNotModified, "Error updating Team")
			return
		}
		RespondWithJSON(w, http.StatusCreated, update)
	}
}

//DeleteTeamController controller definition
func DeleteTeamController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	teamID, err := strconv.Atoi(params["id"])
	if err == nil {
		err := repository.DeleteTeam(teamID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "Team not found")
			return
		}
		RespondWithJSON(w, http.StatusFound, "Team Deleted successfully")
	}
}

//CreateClientController controller definition
func CreateClientController(w http.ResponseWriter, r *http.Request) {
	var client models.Client
	//Decode the incoming Client json data
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		RespondWithError(w, 500, "Invalid Client data on the body")
		return
	}
	//save a new Client

	response, err := repository.CreateClient(client.ClientName, client.ClientDescription, client.ClientCounty, client.ClientIndustrySector, client.ClientCity, client.ClientPhone, client.ClientZip)

	if err != nil {
		RespondWithError(w, 500, "An unexpected error occured")
	}
	RespondWithJSON(w, http.StatusCreated, response)
}

//GetClientsController controller definition
func GetClientsController(w http.ResponseWriter, r *http.Request) {
	clients, err := repository.FetchClients()
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "An Error occurred no clients were found")
	}
	RespondWithJSON(w, http.StatusFound, clients)
}

//GetClientController controller definition
func GetClientController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	clientID, err := strconv.Atoi(params["id"])
	if err == nil {
		client, err := repository.FetchClient(clientID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "Client not found")
			return
		}
		RespondWithJSON(w, http.StatusFound, client)
	}
}

//UpdateClientController controller definition
func UpdateClientController(w http.ResponseWriter, r *http.Request) {

	var client models.Client
	//Decode the incoming Client json data
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		RespondWithError(w, 500, "Invalid Client data on the body")
		return
	}
	params := mux.Vars(r)
	clientID, err := strconv.Atoi(params["id"])
	if err == nil {
		update, err := repository.UpdateClient(client.ClientName, client.ClientDescription, client.ClientCounty, client.ClientIndustrySector, client.ClientCity, client.ClientPhone, client.ClientZip, clientID)
		if err != nil {
			RespondWithError(w, http.StatusNotModified, "Error updating Client")
			return
		}
		RespondWithJSON(w, http.StatusCreated, update)
	}
}

//DeleteClientController controller definition
func DeleteClientController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	clientID, err := strconv.Atoi(params["id"])
	if err == nil {
		err := repository.DeleteClient(clientID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "Client not found")
			return
		}
		RespondWithJSON(w, http.StatusFound, "Client Deleted successfully")
	}
}

//CreateProjectManagerController controller definition
func CreateProjectManagerController(w http.ResponseWriter, r *http.Request) {
	var manager models.ProjectManager
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err == nil {
		user, err := repository.FetchUser(userID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "User not found register him/her as a user first")
			return
		}
		er := json.NewDecoder(r.Body).Decode(&manager)
		if er != nil {
			RespondWithError(w, 500, "Invalid manager payload")
			return
		}
		//save a new manager
		response, err := repository.CreateProjectManager(*user, manager.ProjectManagerProjectID, manager.ProjectManagerClientID)
		if err != nil {
			RespondWithError(w, 500, "Manager not created")
		}
		RespondWithJSON(w, http.StatusCreated, response)
	}
}

//GetProjectManagersController controller definition
func GetProjectManagersController(w http.ResponseWriter, r *http.Request) {
	managers, err := repository.FetchProjectManagers()
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "No Managers")
	}
	RespondWithJSON(w, http.StatusFound, managers)
}

//GetProjectManagerController controller definition
func GetProjectManagerController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	managerID, err := strconv.Atoi(params["id"])
	if err == nil {
		manager, err := repository.FetchProjectManager(managerID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "manager not found")
			return
		}
		RespondWithJSON(w, http.StatusFound, manager)
	}
}


//	DEBUG
////UpdateProjectManagerController controller definition
//func UpdateProjectManagerController(w http.ResponseWriter, r *http.Request) {
//	var manager models.ProjectManager
//	//Decode the incoming manager json data
//	err := json.NewDecoder(r.Body).Decode(&manager)
//	if err != nil {
//		RespondWithError(w, 500, "Invalid manager data on the body")
//		return
//	}
//	params := mux.Vars(r)
//	managerID, err := strconv.Atoi(params["id"])
//	if err == nil {
//		update, err := repository.UpdateProjectManager(manager.ProjectManagerProjectID, manager.ProjectManagerClientID, managerID)
//		if err != nil {
//			RespondWithError(w, http.StatusNotModified, "Error updating manager")
//			return
//		}
//		RespondWithJSON(w, http.StatusCreated, update)
//	}
//}
//	DEBUG


//CreateAssignmentController controller definition
func CreateAssignmentController(w http.ResponseWriter, r *http.Request) {
	var assign models.Assigned
	params := mux.Vars(r)
	employeeID, err := strconv.Atoi(params["id"])
	if err == nil {
		employee, err := repository.FetchEmployee(employeeID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "employee not found register him/her as an employee first before assignment")
			return
		}
		er := json.NewDecoder(r.Body).Decode(&assign)
		if er != nil {
			RespondWithError(w, 500, "Invalid assign payload")
			return
		}
		//save a new assign
		response, err := repository.CreateAssignment(*employee, assign.AssignedActivityID, assign.AssignedRoleID)
		if err != nil {
			RespondWithError(w, 500, "assign not created")
		}
		RespondWithJSON(w, http.StatusCreated, response)
	}
}

//GetAssignmentsController controller definition
func GetAssignmentsController(w http.ResponseWriter, r *http.Request) {
	assigns, err := repository.FetchAssignments()
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "No assigns")
	}
	RespondWithJSON(w, http.StatusFound, assigns)
}

//GetAssignmentController controller definition
func GetAssignmentController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	assignID, err := strconv.Atoi(params["id"])
	if err == nil {
		assign, err := repository.FetchAssignment(assignID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "assign not found")
			return
		}
		RespondWithJSON(w, http.StatusFound, assign)
	}
}

//UpdateAssignmentController controller definition
func UpdateAssignmentController(w http.ResponseWriter, r *http.Request) {
	var assign models.Assigned
	//Decode the incoming assign json data
	err := json.NewDecoder(r.Body).Decode(&assign)
	if err != nil {
		RespondWithError(w, 500, "Invalid assign data on the body")
		return
	}
	params := mux.Vars(r)
	employeeID, err := strconv.Atoi(params["id"])
	if err == nil {
		update, err := repository.UpdateAssignment(assign.AssignedActivityID, assign.AssignedRoleID, employeeID)
		if err != nil {
			RespondWithError(w, http.StatusNotModified, "Error updating assign")
			return
		}
		RespondWithJSON(w, http.StatusCreated, update)
	}
}
