package repository

import (
	"errors"
	"github.com/Muart-C/webcbo/models"
)

//CreateProject method definition
func CreateProject(projectName, projectDescription string, projectIsActive bool, projectBudget, projectMaterialCost, projectLaborCost float64, projectPlannedStartDate, projectPlannedEndDate, projectActualStartDate, projectActualEndDate, projectStatus string, projectTotalHours float64) (*models.Project, error) {
	var project models.Project
	project.ProjectName = projectName
	project.ProjectDescription = projectDescription
	project.ProjectStatus = projectStatus
	project.ProjectIsActive = projectIsActive
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
		return &project, nil
	}
	return nil, errors.New("an error occurred while adding a user")
}

//FetchProjects repository method definition
func FetchProjects() (*[]models.Project, error) {
	var projects []models.Project
	result := models.DB.Find(&projects)
	if result != nil {
		return &projects, nil
	}
	return nil, errors.New("error returning projects")
}

//FetchProject repository method definition
func FetchProject(id int) (*models.Project, error) {
	var project models.Project
	models.DB.Where("id = ?", id).Find(&project)
	if project.ID == id {
		return &project, nil
	}
	return nil, errors.New("error returning the specified project")
}

//UpdateProject repo method definition
func UpdateProject(projectName, projectDescription string, projectIsActive bool, projectBudget, projectLaborCost, projectMaterialCost float64, projectPlannedStartDate, projectPlannedEndDate, projectActualStartDate, projectActualEndDate, projectStatus string, projectTotalHours float64, id int) (*models.Project, error) {
	var project models.Project
	models.DB.Where("id = ?", id).Find(&project)
	if project.ID == id {
		project.ProjectName = projectName
		project.ProjectDescription = projectDescription
		project.ProjectStatus = projectStatus
		project.ProjectIsActive = projectIsActive
		project.ProjectBudget = projectBudget
		project.ProjectMaterialCost = projectMaterialCost
		project.ProjectLaborCost = projectLaborCost
		project.ProjectPlannedStartDate = projectPlannedStartDate
		project.ProjectPlannedEndDate = projectPlannedEndDate
		project.ProjectActualStartDate = projectActualStartDate
		project.ProjectActualEndDate = projectActualEndDate
		project.ProjectTotalHours = projectTotalHours
		models.DB.Save(&project)
		return &project, nil
	}
	return nil, errors.New("error occurred while updating the project")
}

//CreateMilestone repo method definition
func CreateMilestone(project models.Project, milestoneName, milestoneStatus, milestoneDeliverables, milestoneDueDate, milestoneAchievedDate string, milestoneTotalHours float64) (*models.Milestone, error) {
	var milestone models.Milestone
	milestone.MilestoneName = milestoneName
	milestone.MilestoneStatus = milestoneStatus
	milestone.MilestoneDeliverables = milestoneDeliverables
	milestone.MilestoneDueDate = milestoneDueDate
	milestone.MilestoneAchievedDate = milestoneAchievedDate
	milestone.MilestoneTotalHours = milestoneTotalHours
	milestone.MilestoneProjectID = project.ID
	results := models.DB.Create(&milestone)
	if results != nil {
		return &milestone, nil
	}
	return nil, errors.New("an error occurred while assigning the milestone")
}

//FetchMilestones repository method definition
func FetchMilestones() (*[]models.Milestone, error) {
	var milestones []models.Milestone
	result := models.DB.Find(&milestones)
	if result != nil {
		return &milestones, nil
	}
	return nil, errors.New("error returning milestones")
}

//FetchMilestonesInProject repository method definition
func FetchMilestonesInProject(id int) (*[]models.Milestone, error) {
	var milestones []models.Milestone
	//result :=models.DB.Find(&milestones)
	result := models.DB.Where("id = ?", id).Find(&milestones)
	if result != nil {
		return &milestones, nil
	}
	return nil, errors.New("error returning milestones for that particular project")
}

//FetchMilestone repository method definition
func FetchMilestone(id int) (*models.Milestone, error) {
	var milestone models.Milestone
	models.DB.Where("id = ?", id).Find(&milestone)
	if milestone.ID == id {
		return &milestone, nil
	}
	return nil, errors.New("error returning the specified project")
}

//UpdateMilestone repo method definition
func UpdateMilestone(milestoneName, milestoneStatus, milestoneDeliverables, milestoneDueDate, milestoneAchievedDate string, milestoneTotalHours float64, id int) (*models.Milestone, error) {
	var milestone models.Milestone
	models.DB.Where("id=?", id).Find(&milestone)
	if milestone.ID == id {
		milestone.MilestoneName = milestoneName
		milestone.MilestoneStatus = milestoneStatus
		milestone.MilestoneDeliverables = milestoneDeliverables
		milestone.MilestoneDueDate = milestoneDueDate
		milestone.MilestoneAchievedDate = milestoneAchievedDate
		milestone.MilestoneTotalHours = milestoneTotalHours
		models.DB.Save(&milestone)
		return &milestone, nil
	}
	return nil, errors.New("an error occurred while updating the milestone")
}
