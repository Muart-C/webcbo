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
		RespondWithError(w,http.StatusInternalServerError, "An unexpected error occurred")
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
