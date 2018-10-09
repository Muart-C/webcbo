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
func CreateUserController(w http.ResponseWriter, r *http.Request)  {
	var user models.User
	//Decode the incoming user json data
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		RespondWithError(w,500,"Invalid user data on the body")
		return
	}
	//save a new user

	response,err:=repository.CreateUser(user.UserName,user.Email,user.LastName,user.FirstName)

	if err != nil {
		RespondWithError(w,500,"An unexpected error occured")
	}
	RespondWithJSON(w,http.StatusCreated,response)
}


//GetUsersController controller definition
func GetUsersController(w http.ResponseWriter, r *http.Request)  {
	users,err := repository.FetchUsers()
	if err != nil {
		RespondWithError(w,http.StatusNotFound,err.Error())
	}
	RespondWithJSON(w,http.StatusFound,users)
}

//GetUserController controller definition
func GetUserController(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err == nil {
		user, err:= repository.FetchUser(userID)
		if err != nil {
			RespondWithError(w,http.StatusNotFound,"User not found")
			return
		}
	RespondWithJSON(w,http.StatusFound,user)
	}
}

//UpdateUserController controller definition
func UpdateUserController(w http.ResponseWriter, r *http.Request)  {

	var user models.User
	//Decode the incoming user json data
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		RespondWithError(w,500,"Invalid user data on the body")
		return
	}
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err == nil {
		update, err := repository.UpdateUser(user.UserName,user.Email,user.FirstName,user.LastName,userID)
		if err != nil {
			RespondWithError(w,http.StatusNotModified,"Error updating user")
			return
		}
		RespondWithJSON(w,http.StatusCreated,update)
	}
}

//DeleteUserController controller definition
func DeleteUserController(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err == nil {
		 err:= repository.DeleteUser(userID)
		if err != nil {
			RespondWithError(w,http.StatusNotFound,"User not found")
			return
		}
		RespondWithJSON(w,http.StatusFound,"User Deleted successfully")
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
func GetEmployeesController(w http.ResponseWriter, r *http.Request)  {
	employees,err := repository.FetchEmployees()
	if err != nil {
		RespondWithError(w,http.StatusNotFound,err.Error())
	}
	RespondWithJSON(w,http.StatusFound,employees)
}

//GetEmployeeController controller definition
func GetEmployeeController(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	employeeID, err := strconv.Atoi(params["eid"])
	if err == nil {
		employee, err:= repository.FetchEmployee(employeeID)
		if err != nil {
			RespondWithError(w,http.StatusNotFound,"User not found")
			return
		}
		RespondWithJSON(w,http.StatusFound,employee)
	}
}

//UpdateEmployeeController controller definition
func UpdateEmployeeController(w http.ResponseWriter, r *http.Request)  {

	var employee models.Employee
	//Decode the incoming user json data
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		RespondWithError(w,500,"Invalid user data on the body")
		return
	}
	params := mux.Vars(r)
	employeeID, err := strconv.Atoi(params["eid"])
	if err == nil {
		update, err := repository.UpdateEmployee(employee.EmployeeProfession,employeeID)
		if err != nil {
			RespondWithError(w,http.StatusNotModified,"Error updating user")
			return
		}
		RespondWithJSON(w,http.StatusCreated,update)
	}
}

//DeleteEmployeeController controller definition
func DeleteEmployeeController(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	employeeID, err := strconv.Atoi(params["eid"])
	if err == nil {
		err:= repository.DeleteEmployee(employeeID)
		if err != nil {
			RespondWithError(w,http.StatusNotFound,"User not found")
			return
		}
		RespondWithJSON(w,http.StatusFound,"Employee Deleted successfully")
	}
}