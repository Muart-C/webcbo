package controllers

import (
	"encoding/json"
	"github.com/Muart-C/webcbo/models"
	"github.com/Muart-C/webcbo/repository"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//CreateTaskController controller definition
func CreateTaskController(w http.ResponseWriter, r *http.Request)  {
	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		RespondWithError(w,http.StatusBadRequest,"Could not decode the task json")
	}
	//save a task
	response, err := repository.CreateTask(task.TaskName,task.TaskInstructions,task.TaskTotalHours,task.TaskPlannedBudget,task.TaskActualBudget,task.TaskPlannedStartDate,task.TaskActualStartDate,task.TaskPlannedEndDate,task.TaskActualEndDate)
	if err != nil {
		RespondWithError(w,http.StatusNotImplemented,"could not create a task")
	}
	RespondWithJSON(w,http.StatusCreated,response)
}

//GetTasksController controller definition
func GetTasksController(w http.ResponseWriter, r *http.Request){
	tasks, err := repository.FetchTasks()
	if err != nil {
		RespondWithError(w,http.StatusNotFound,"error retrieving the tasks")
	}
	RespondWithJSON(w,http.StatusFound,tasks)
}

//GetTaskController controller definition
func GetTaskController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])
	if err != nil {
		task, err := repository.FetchTask(taskID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "error retrieving the task")
			return
		}
		RespondWithJSON(w, http.StatusFound, task)
	}
}

//UpdateTaskController controller definition
func UpdateTaskController(w http.ResponseWriter, r *http.Request)  {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		RespondWithError(w,http.StatusInternalServerError,"error decoding the project struct")
	}
	params := mux.Vars(r)
	taskID,err := strconv.Atoi(params["id"])
	if err == nil {
		update,err := repository.UpdateTask(task.TaskName,task.TaskInstructions,task.TaskTotalHours,task.TaskPlannedBudget,task.TaskActualBudget,task.TaskPlannedStartDate,task.TaskActualStartDate,task.TaskPlannedEndDate,task.TaskActualEndDate, taskID)
		if err != nil {
			RespondWithError(w,http.StatusNotModified,"an error occurred while updating the task")
			return
		}
		RespondWithJSON(w,http.StatusCreated,update)
	}
}

//CreateTaskStatusController controller definition
func CreateTaskStatusController(w http.ResponseWriter, r *http.Request)  {
	var status  models.TaskStatus

	err := json.NewDecoder(r.Body).Decode(&status)
	if err != nil {
		RespondWithError(w,http.StatusBadRequest,"Could not decode the status json")
	}
	//save a status
	response, err := repository.CreateTaskStatus(status.TaskStatus,status.TaskPriority)
	if err != nil {
		RespondWithError(w,http.StatusNotImplemented,"could not create a status")
	}
	RespondWithJSON(w,http.StatusCreated,response)
}


//UpdateTaskStatusController controller definition
func UpdateTaskStatusController(w http.ResponseWriter, r *http.Request)  {
	var status models.TaskStatus
	err := json.NewDecoder(r.Body).Decode(&status)
	if err != nil {
		RespondWithError(w,http.StatusInternalServerError,"error decoding the status struct")
	}
	params := mux.Vars(r)
	statusID,err := strconv.Atoi(params["id"])
	if err == nil {
		update,err := repository.UpdateTaskStatus(status.TaskStatus,status.TaskPriority, statusID)
		if err != nil {
			RespondWithError(w,http.StatusNotModified,"an error occurred while updating the task status")
			return
		}
		RespondWithJSON(w,http.StatusCreated,update)
	}
}

//CreateActivityStatusController controller definition
func CreateActivityStatusController(w http.ResponseWriter, r *http.Request)  {
	var status  models.ActivityStatus

	err := json.NewDecoder(r.Body).Decode(&status)
	if err != nil {
		RespondWithError(w,http.StatusBadRequest,"Could not decode the status json")
	}
	//save a status
	response, err := repository.CreateActivityStatus(status.ActivityStatus)
	if err != nil {
		RespondWithError(w,http.StatusNotImplemented,"could not create a status")
	}
	RespondWithJSON(w,http.StatusCreated,response)
}


//UpdateActivityStatusController controller definition
func UpdateActivityStatusController(w http.ResponseWriter, r *http.Request)  {
	var status models.ActivityStatus
	err := json.NewDecoder(r.Body).Decode(&status)
	if err != nil {
		RespondWithError(w,http.StatusInternalServerError,"error decoding the status struct")
	}
	params := mux.Vars(r)
	statusID,err := strconv.Atoi(params["id"])
	if err == nil {
		update,err := repository.UpdateActivityStatus(status.ActivityStatus,statusID)
		if err != nil {
			RespondWithError(w,http.StatusNotModified,"an error occurred while updating the activity status")
			return
		}
		RespondWithJSON(w,http.StatusCreated,update)
	}
}

//CreateActivityController controller definition
func CreateActivityController(w http.ResponseWriter, r *http.Request)  {
	var activity models.Activity

	err := json.NewDecoder(r.Body).Decode(&activity)
	if err != nil {
		RespondWithError(w,http.StatusBadRequest,"Could not decode the activity  json")
	}
	//save a activity
	response, err := repository.CreateActivity(activity.ActivityName,activity.ActivityPlannedBudget,activity.ActivityActualBudget,activity.ActivityPlannedStartDate,activity.ActivityPlannedEndDate,activity.ActivityActualStartDate,activity.ActivityActualEndDate)
	if err != nil {
		RespondWithError(w,http.StatusNotImplemented,"could not create an activity")
	}
	RespondWithJSON(w,http.StatusCreated,response)
}

////GetProjectsController controller definition
//func GetProjectsController(w http.ResponseWriter, r *http.Request){
//	projects, err := repository.FetchProjects()
//	if err != nil {
//		RespondWithError(w,http.StatusNotFound,"error retrieving the projects")
//	}
//	RespondWithJSON(w,http.StatusFound,projects)
//}
//
////GetProjectController controller definition
//func GetProjectController(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	projectID, err := strconv.Atoi(params["id"])
//	if err != nil {
//		project, err := repository.FetchProject(projectID)
//		if err != nil {
//			RespondWithError(w, http.StatusNotFound, "error retrieving the project")
//			return
//		}
//		RespondWithJSON(w, http.StatusFound, project)
//	}
//}

//UpdateActivityController controller definition
func UpdateActivityController(w http.ResponseWriter, r *http.Request)  {
	var activity models.Activity
	err := json.NewDecoder(r.Body).Decode(&activity)
	if err != nil {
		RespondWithError(w,http.StatusInternalServerError,"error decoding the activity struct")
	}
	params := mux.Vars(r)
	activityID,err := strconv.Atoi(params["id"])
	if err == nil {
		update,err := repository.UpdateActivity(activity.ActivityName,activity.ActivityPlannedBudget,activity.ActivityActualBudget,activity.ActivityPlannedStartDate,activity.ActivityPlannedEndDate,activity.ActivityActualStartDate,activity.ActivityActualEndDate, activityID)
		if err != nil {
			RespondWithError(w,http.StatusNotModified,"an error occurred while updating the activity")
			return
		}
		RespondWithJSON(w,http.StatusCreated,update)
	}
}