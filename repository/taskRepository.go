package repository

import (
	"errors"
	"github.com/Muart-C/webcbo/models"
)

//CreateTask method definition
func CreateTask(taskName,taskInstructions string,taskTotalHours,taskPlannedBudget,taskActualBudget float64,taskPlannedStartDate,taskPlannedEndDate,taskActualStartDate,taskActualEndDate string) (*models.Task, error) {
	var task models.Task
	task.TaskName= taskName
	task.TaskInstructions= taskInstructions
	task.TaskTotalHours= taskTotalHours
	task.TaskPlannedBudget= taskPlannedBudget
	task.TaskActualBudget= taskActualBudget
	task.TaskPlannedStartDate= taskPlannedStartDate
	task.TaskActualStartDate= taskActualStartDate
	task.TaskPlannedEndDate= taskPlannedEndDate
	task.TaskActualEndDate= taskActualEndDate
	result := models.DB.Create(&task)
	if result != nil {
		return &task,nil
	}
	return nil,errors.New("an error occurred while adding a task")
}

//FetchTasks repository method definition
func FetchTasks()(*[]models.Task,error)  {
	var tasks []models.Task
	result :=models.DB.Find(&tasks)
	if  result != nil{
		return &tasks,nil
	}
	return nil, errors.New("error returning tasks")
}

//FetchTask repository method definition
func FetchTask(id int)(*models.Task ,error)  {
	var task models.Task
	models.DB.Where("id = ?",id).Find(&task)
	if task.ID == id {
		return &task,nil
	}
	return nil,errors.New("error returning the specified task")
}

//UpdateTaskrepo method definition
func UpdateTask(taskName,taskInstructions string,taskTotalHours,taskPlannedBudget,taskActualBudget float64,taskPlannedStartDate,taskPlannedEndDate,taskActualStartDate,taskActualEndDate string,id int) (*models.Task, error) {
	var task models.Task
	models.DB.Where("id = ?",id).Find(&task)
	if task.ID == id {
		task.TaskName= taskName
		task.TaskInstructions= taskInstructions
		task.TaskTotalHours= taskTotalHours
		task.TaskPlannedBudget= taskPlannedBudget
		task.TaskActualBudget= taskActualBudget
		task.TaskPlannedStartDate= taskPlannedStartDate
		task.TaskActualStartDate= taskActualStartDate
		task.TaskPlannedEndDate= taskPlannedEndDate
		task.TaskActualEndDate= taskActualEndDate
		models.DB.Save(&task)
		return &task,nil
	}
	return nil,errors.New("error occurred while updating the task")
}

//CreateTaskStatus repo method definition
func CreateTaskStatus(taskStatus,taskPriority string)(*models.TaskStatus,error)  {
	var status models.TaskStatus
	status.TaskStatus= taskStatus
	status.TaskPriority= taskPriority
	result := models.DB.Create(status)
	if result != nil {
		return &status,nil
	}
	return nil,errors.New("an error occurred while assigning a task status")
}

//UpdateTaskStatus repo method definition
func UpdateTaskStatus(taskStatus,taskPriority string, id int) (*models.TaskStatus,error) {
	var status models.TaskStatus
	models.DB.Where("id=?",id).Find(&status)
	if status.ID == id {
		status.TaskStatus= taskStatus
		status.TaskPriority= taskPriority
		models.DB.Save(&status)
		return &status, nil
	}
	return nil, errors.New("an error occurred while updating the status of the task")
}

//CreateActivity repo method definition
func CreateActivity(activityName string,activityPlannedBudget,activityActualBudget float64,activityPlannedStartDate,activityPlannedEndDate,activityActualStartDate,activityActualEndDate string) (*models.Activity, error) {
	var activity models.Activity
	activity.ActivityName= activityName
	activity.ActivityPlannedBudget= activityPlannedBudget
	activity.ActivityActualBudget= activityActualBudget
	activity.ActivityPlannedStartDate= activityPlannedStartDate
	activity.ActivityPlannedEndDate= activityPlannedEndDate
	activity.ActivityActualStartDate= activityActualStartDate
	activity.ActivityActualEndDate= activityActualEndDate
	results := models.DB.Create(&activity)
	if results != nil {
		return &activity, nil
	}
	return nil, errors.New("an error occurred while assigning the activity")
}

//UpdateMilestone repo method definition
func UpdateActivity(activityName string,activityPlannedBudget,activityActualBudget float64,activityPlannedStartDate,activityPlannedEndDate,activityActualStartDate,activityActualEndDate string, id int) (*models.Activity, error){
	var activity models.Activity
	models.DB.Where("id=?",id).Find(&activity)
	if activity.ID == id {
		activity.ActivityName= activityName
		activity.ActivityPlannedBudget= activityPlannedBudget
		activity.ActivityActualBudget= activityActualBudget
		activity.ActivityPlannedStartDate= activityPlannedStartDate
		activity.ActivityPlannedEndDate= activityPlannedEndDate
		activity.ActivityActualStartDate= activityActualStartDate
		activity.ActivityActualEndDate= activityActualEndDate
		models.DB.Save(&activity)
		return &activity, nil
	}
	return nil, errors.New("an error occurred while updating the activity")
}

//CreateActivityStatus repo method definition
func CreateActivityStatus(activityStatus string)(*models.ActivityStatus,error)  {
	var status models.ActivityStatus
	status.ActivityStatus= activityStatus
	result := models.DB.Create(status)
	if result != nil {
		return &status,nil
	}
	return nil,errors.New("an error occurred while assigning a activity status")
}

//UpdateActivityStatus repo method definition
func UpdateActivityStatus(activityStatus string, id int) (*models.ActivityStatus,error) {
	var status models.ActivityStatus
	models.DB.Where("id=?",id).Find(&status)
	if status.ID == id {
		status.ActivityStatus= activityStatus
		models.DB.Save(&status)
		return &status, nil
	}
	return nil, errors.New("an error occurred while updating the status of the activity")
}