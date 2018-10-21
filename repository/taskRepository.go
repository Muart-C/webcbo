package repository

import (
	"errors"
	"github.com/Muart-C/webcbo/models"
)

//CreateTask method definition
func CreateTask(project models.Project, taskName, taskStatus, taskInstructions, taskPriority string, taskTotalHours, taskPlannedBudget, taskActualBudget float64, taskPlannedStartDate, taskPlannedEndDate, taskActualStartDate, taskActualEndDate string) (*models.Task, error) {
	var task models.Task
	task.TaskName = taskName
	task.TaskStatus = taskStatus
	task.TaskInstructions = taskInstructions
	task.TaskPriority = taskPriority
	task.TaskTotalHours = taskTotalHours
	task.TaskPlannedBudget = taskPlannedBudget
	task.TaskActualBudget = taskActualBudget
	task.TaskPlannedStartDate = taskPlannedStartDate
	task.TaskActualStartDate = taskActualStartDate
	task.TaskPlannedEndDate = taskPlannedEndDate
	task.TaskActualEndDate = taskActualEndDate
	task.TaskProjectID = project.ID
	result := models.DB.Create(&task)
	if result != nil {
		return &task, nil
	}
	return nil, errors.New("an error occurred while adding a task")
}

//FetchTasksInProject repository method definition
func FetchTasksInProject(id int) (*[]models.Task, error) {
	var tasks []models.Task
	//result :=models.DB.Find(&milestones)
	result := models.DB.Where("id = ?", id).Find(&tasks)
	if result != nil {
		return &tasks, nil
	}
	return nil, errors.New("error returning tasks belonging to that particular project")
}

//FetchTasks repository method definition
func FetchTasks() (*[]models.Task, error) {
	var tasks []models.Task
	result := models.DB.Find(&tasks)
	if result != nil {
		return &tasks, nil
	}
	return nil, errors.New("error returning tasks")
}

//FetchTask repository method definition
func FetchTask(id int) (*models.Task, error) {
	var task models.Task
	models.DB.Where("id = ?", id).Find(&task)
	if task.ID == id {
		return &task, nil
	}
	return nil, errors.New("error returning the specified task")
}

//UpdateTask repo method definition
func UpdateTask(taskName, taskStatus, taskInstructions, taskPriority string, taskTotalHours, taskPlannedBudget, taskActualBudget float64, taskPlannedStartDate, taskPlannedEndDate, taskActualStartDate, taskActualEndDate string, id int) (*models.Task, error) {
	var task models.Task
	models.DB.Where("id = ?", id).Find(&task)
	if task.ID == id {
		task.TaskName = taskName
		task.TaskStatus = taskStatus
		task.TaskInstructions = taskInstructions
		task.TaskPriority = taskPriority
		task.TaskTotalHours = taskTotalHours
		task.TaskPlannedBudget = taskPlannedBudget
		task.TaskActualBudget = taskActualBudget
		task.TaskPlannedStartDate = taskPlannedStartDate
		task.TaskActualStartDate = taskActualStartDate
		task.TaskPlannedEndDate = taskPlannedEndDate
		task.TaskActualEndDate = taskActualEndDate
		models.DB.Save(&task)
		return &task, nil
	}
	return nil, errors.New("error occurred while updating the task")
}

//CreateActivity repo method definition
func CreateActivity(task models.Task, activityName, activityStatus string, activityPlannedBudget, activityActualBudget float64, activityPlannedStartDate, activityPlannedEndDate, activityActualStartDate, activityActualEndDate string) (*models.Activity, error) {
	var activity models.Activity
	activity.ActivityName = activityName
	activity.ActivityStatus = activityStatus
	activity.ActivityPlannedBudget = activityPlannedBudget
	activity.ActivityActualBudget = activityActualBudget
	activity.ActivityPlannedStartDate = activityPlannedStartDate
	activity.ActivityPlannedEndDate = activityPlannedEndDate
	activity.ActivityActualStartDate = activityActualStartDate
	activity.ActivityActualEndDate = activityActualEndDate
	activity.TaskActivityID = task.ID
	results := models.DB.Create(&activity)
	if results != nil {
		return &activity, nil
	}
	return nil, errors.New("an error occurred while assigning the activity")
}

//FetchActivitiesInTask repository method definition
func FetchActivitiesInTask(id int) (*[]models.Activity, error) {
	var activities []models.Activity
	//result :=models.DB.Find(&milestones)
	result := models.DB.Where("id = ?", id).Find(&activities)
	if result != nil {
		return &activities, nil
	}
	return nil, errors.New("error returning milestones")
}

//UpdateActivity repo method definition
func UpdateActivity(activityName, activityStatus string, activityPlannedBudget, activityActualBudget float64, activityPlannedStartDate, activityPlannedEndDate, activityActualStartDate, activityActualEndDate string, id int) (*models.Activity, error) {
	var activity models.Activity
	models.DB.Where("id=?", id).Find(&activity)
	if activity.ID == id {
		activity.ActivityName = activityName
		activity.ActivityStatus = activityStatus
		activity.ActivityPlannedBudget = activityPlannedBudget
		activity.ActivityActualBudget = activityActualBudget
		activity.ActivityPlannedStartDate = activityPlannedStartDate
		activity.ActivityPlannedEndDate = activityPlannedEndDate
		activity.ActivityActualStartDate = activityActualStartDate
		activity.ActivityActualEndDate = activityActualEndDate
		models.DB.Save(&activity)
		return &activity, nil
	}
	return nil, errors.New("an error occurred while updating the activity")
}
