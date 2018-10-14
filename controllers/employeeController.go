package controllers

import (
	"encoding/json"
	"github.com/Muart-C/webcbo/models"
	"github.com/Muart-C/webcbo/repository"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

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
		//save a new employee

		response, err := repository.CreateEmployee(*user, employee.EmployeeProfession)

		if err != nil {
			RespondWithError(w, 500, "An unexpected error occurred")
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
			RespondWithError(w, 500, "An unexpected error occurred")
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
