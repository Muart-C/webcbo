package repository

import (
	"errors"
	"github.com/Muart-C/webcbo/models"
)

//CreateProject method definition
func CreateProject(projectName,projectDescription string,projectBudget,projectMaterialCost,projectLaborCost float64,projectPlannedStartDate,projectPlannedEndDate,projectActualStartDate,projectActualEndDate string,projectTotalHours float64) (*models.Project, error) {
	var project models.Project
	project.ProjectName = projectName
	project.ProjectDescription = projectDescription
	project.ProjectBudget = projectBudget
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
func UpdateProject(projectName,projectDescription string,projectBudget,projectLaborCost,projectMaterialCost float64,projectPlannedStartDate,projectPlannedEndDate,projectActualStartDate,projectActualEndDate string,projectTotalHours float64, id int) (*models.Project, error)  {
	var project models.Project
	models.DB.Where("id = ?",id).Find(&project)
	if project.ID == id {
		project.ProjectName = projectName
		project.ProjectDescription = projectDescription
		project.ProjectBudget = projectBudget
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

//CreateProjectStatus repo method definition
func CreateProjectStatus(projectStatus string,projectIsActive bool)(*models.ProjectStatus,error)  {
	var status models.ProjectStatus
	status.ProjectStatus = projectStatus
	status.ProjectIsActive = projectIsActive
	result := models.DB.Create(status)
	if result != nil {
		return &status,nil
	}
	return nil,errors.New("an error occurred while assigning a project status")
}

//UpdateProjectStatus repo method definition
func UpdateProjectStatus(projectStatus string, projectIsActive bool, id int) (*models.ProjectStatus,error) {
	var status models.ProjectStatus
	models.DB.Where("id=?",id).Find(&status)
	if status.ID == id {
		status.ProjectStatus = projectStatus
		status.ProjectIsActive = projectIsActive
		models.DB.Save(&status)
		return &status, nil
	}
	return nil, errors.New("an error occurred while updating the status of the project")
}

//CreateMilestone repo method definition
func CreateMilestone(milestoneName,milestoneDeliverables,milestoneDueDate,milestoneAchievedDate string,milestoneTotalHours float64) (*models.Milestone, error) {
	var milestone models.Milestone
	milestone.MilestoneName = milestoneName
	milestone.MilestoneDeliverables = milestoneDeliverables
	milestone.MilestoneDueDate = milestoneDueDate
	milestone.MilestoneAchievedDate = milestoneAchievedDate
	milestone.MilestoneTotalHours = milestoneTotalHours
	results := models.DB.Create(&milestone)
	if results != nil {
		return &milestone, nil
	}
	return nil, errors.New("an error occurred while assigning the milestone")
}

//UpdateMilestone repo method definition
func UpdateMilestone(milestoneName,milestoneDeliverables,milestoneDueDate,milestoneAchievedDate string,milestoneTotalHours float64, id int) (*models.Milestone, error) {
	var milestone models.Milestone
	models.DB.Where("id=?",id).Find(&milestone)
	if milestone.ID == id {
		milestone.MilestoneName= milestoneName
		milestone.MilestoneDeliverables= milestoneDeliverables
		milestone.MilestoneDueDate = milestoneDueDate
		milestone.MilestoneAchievedDate = milestoneAchievedDate
		milestone.MilestoneTotalHours = milestoneTotalHours
		models.DB.Save(&milestone)
		return &milestone, nil
	}
	return nil, errors.New("an error occurred while updating the milestone")
}

//CreateMilestoneStatus repo method definition
func CreateMilestoneStatus(milestoneStatus string)(*models.MilestoneStatus,error)  {
	var status models.MilestoneStatus
	status.MilestoneStatus= milestoneStatus
	result := models.DB.Create(status)
	if result != nil {
		return &status,nil
	}
	return nil,errors.New("an error occurred while assigning a milestone status")
}

//UpdateMilestoneStatus repo method definition
func UpdateMilestoneStatus(milestoneStatus string, id int) (*models.MilestoneStatus,error) {
	var status models.MilestoneStatus
	models.DB.Where("id=?",id).Find(&status)
	if status.ID == id {
		status.MilestoneStatus= milestoneStatus
		models.DB.Save(&status)
		return &status, nil
	}
	return nil, errors.New("an error occurred while updating the status of the milestone")
}