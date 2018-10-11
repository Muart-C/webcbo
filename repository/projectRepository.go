package repository

import (
	"errors"
	"github.com/Muart-C/webcbo/models"
	"time"
)

//CreateProject method definition
func CreateProject(projectName,projectDescription string,projectMaterialCost,projectLaborCost float64,projectPlannedStartDate,projectPlannedEndDate,projectActualStartDate,projectActualEndDate time.Time,projectTotalHours float64) (*models.Project, error) {
	var project models.Project
	project.ProjectName = projectName
	project.ProjectDescription = projectDescription
	project.ProjectMaterialCost = projectMaterialCost
	project.ProjectLaborCost = projectLaborCost
	project.ProjectTotalHours = projectTotalHours
	project.ProjectPlannedStartDate = projectPlannedStartDate
	project.ProjectPlannedEndDate = projectPlannedEndDate
	project.ProjectActualStartDate = projectActualStartDate
	project.ProjectActualEndDate = projectActualEndDate

	result := models.DB.Create(&project)
	if result != nil {
		return &project,nil
	}
	return nil,errors.New("an error occurred while adding a user")
}

//FetchProjects repository method definition
func FetchProjects()(*[]models.Project,error)  {
	var projects []models.Project
	result :=models.DB.Find(&projects)
	if  result != nil{
		return &projects,nil
	}
	return nil, errors.New("error returning projects")
}

//FetchProject repository method definition
func FetchProject(id int)(*models.Project,error)  {
	var project models.Project
	models.DB.Where("id = ?",id).Find(&project)
	if project.ID == id {
		return &project,nil
	}
	return nil,errors.New("error returning the specified project")
}

//UpdateProject repo method definition
func UpdateProject(projectName,projectDescription string,projectLaborCost,projectMaterialCost float64,projectPlannedStartDate,projectPlannedEndDate,projectActualStartDate,projectActualEndDate time.Time,projectTotalHours float64, id int) (*models.Project, error)  {
	var project models.Project
	models.DB.Where("id = ?",id).Find(&project)
	if project.ID == id {
		project.ProjectName = projectName
		project.ProjectDescription = projectDescription
		project.ProjectMaterialCost = projectMaterialCost
		project.ProjectLaborCost = projectLaborCost
		project.ProjectPlannedStartDate = projectPlannedStartDate
		project.ProjectPlannedEndDate = projectPlannedEndDate
		project.ProjectActualStartDate = projectActualStartDate
		project.ProjectActualEndDate = projectActualEndDate
		project.ProjectTotalHours = projectTotalHours
		models.DB.Save(&project)
		return &project,nil
	}
	return nil,errors.New("error occurred while updating the project")
}

