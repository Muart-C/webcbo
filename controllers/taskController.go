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
	params := mux.Vars(r)
	projectID,err := strconv.Atoi(params["id"])
	if err == nil {
		project,err := repository.FetchProject(projectID)
		if err != nil {
			RespondWithError(w,http.StatusNotFound,"such project does not exist please create project first then add tasks to it")
			return
		}
		er := json.NewDecoder(r.Body).Decode(&task)
		if er != nil {
			RespondWithError(w,http.StatusBadRequest,"Could not decode the task json")
		}
		//save a task
		response, err := repository.CreateTask(*project,task.TaskName,task.TaskStatus,task.TaskInstructions,task.TaskPriority,task.TaskTotalHours,task.TaskPlannedBudget,task.TaskActualBudget,task.TaskPlannedStartDate,task.TaskActualStartDate,task.TaskPlannedEndDate,task.TaskActualEndDate)
		if err != nil {
			RespondWithError(w,http.StatusNotImplemented,"could not create a task")
		}
		RespondWithJSON(w,http.StatusCreated,response)
	}
}

//GetTasksInProjectController controller definition
func GetTasksInProjectController(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	projectID, err := strconv.Atoi(params["id"])
	if err == nil {

		activities, err := repository.FetchTasksInProject(projectID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "error retrieving the activities of the particular task")
		}
		RespondWithJSON(w, http.StatusFound, activities)
	}

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
		update,err := repository.UpdateTask(task.TaskName,task.TaskStatus,task.TaskInstructions,task.TaskPriority,task.TaskTotalHours,task.TaskPlannedBudget,task.TaskActualBudget,task.TaskPlannedStartDate,task.TaskActualStartDate,task.TaskPlannedEndDate,task.TaskActualEndDate, taskID)
		if err != nil {
			RespondWithError(w,http.StatusNotModified,"an error occurred while updating the task")
			return
		}
		RespondWithJSON(w,http.StatusCreated,update)
	}
}


//CreateActivityController controller definition
func CreateActivityController(w http.ResponseWriter, r *http.Request) {
	var activity models.Activity

	params := mux.Vars(r)
	taskID,err := strconv.Atoi(params["id"])
	if err == nil {
		task, err := repository.FetchTask(taskID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "such task does not exist please create task first then add activities to it")
			return
		}
		er := json.NewDecoder(r.Body).Decode(&activity)
		if er != nil {
			RespondWithError(w, http.StatusBadRequest, "Could not decode the activity json")
		}
		//save a activity
		response, err := repository.CreateActivity(*task, activity.ActivityName, activity.ActivityStatus, activity.ActivityPlannedBudget, activity.ActivityActualBudget, activity.ActivityPlannedStartDate, activity.ActivityPlannedEndDate, activity.ActivityActualStartDate, activity.ActivityActualEndDate)
		if err != nil {
			RespondWithError(w, http.StatusNotImplemented, "could not create an activity")
		}
		RespondWithJSON(w, http.StatusCreated, response)
	}
}

//GetActivitiesInTaskController controller definition
func GetActivitiesInTaskController(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])
	if err == nil {

		activities, err := repository.FetchActivitiesInTask(taskID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "error retrieving the activities of the particular task")
		}
		RespondWithJSON(w, http.StatusFound, activities)
	}

}
//GetActivityController controller definition
func GetActivityController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	milestoneID, err := strconv.Atoi(params["id"])
	if err != nil {
		milestone, err := repository.FetchMilestone(milestoneID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "error retrieving the milestone")
			return
		}
		RespondWithJSON(w, http.StatusFound, milestone)
	}
}
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
		update,err := repository.UpdateActivity(activity.ActivityName,activity.ActivityStatus,activity.ActivityPlannedBudget,activity.ActivityActualBudget,activity.ActivityPlannedStartDate,activity.ActivityPlannedEndDate,activity.ActivityActualStartDate,activity.ActivityActualEndDate, activityID)
		if err != nil {
			RespondWithError(w,http.StatusNotModified,"an error occurred while updating the activity")
			return
		}
		RespondWithJSON(w,http.StatusCreated,update)
	}
}