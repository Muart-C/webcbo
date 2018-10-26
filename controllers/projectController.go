package controllers

import (
	"encoding/json"
	"github.com/Muart-C/webcbo/models"
	"github.com/Muart-C/webcbo/repository"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//CreateProjectController controller definition
func CreateProjectController(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	er := json.NewDecoder(r.Body).Decode(&project)
	if er != nil {
		RespondWithError(w, 500, "Invalid project payload")
		return
	}
	//save a project
	response, err := repository.CreateProject(project.ProjectName, project.ProjectDescription, project.ProjectIsActive, project.ProjectBudget, project.ProjectMaterialCost, project.ProjectLaborCost, project.ProjectPlannedStartDate, project.ProjectPlannedEndDate, project.ProjectActualStartDate, project.ProjectActualEndDate, project.ProjectStatus, project.ProjectTotalHours)
	if err != nil {
		RespondWithError(w, http.StatusNotImplemented, "could not create a project")
	}
	RespondWithJSON(w, http.StatusCreated, response)
}

//GetProjectsController controller definition
func GetProjectsController(w http.ResponseWriter, r *http.Request) {
	projects, err := repository.FetchProjects()
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "error retrieving the projects")
	}
	RespondWithJSON(w, http.StatusFound, projects)
}

//GetProjectController controller definition
func GetProjectController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectID, err := strconv.Atoi(params["id"])
	if err != nil {
		project, err := repository.FetchProject(projectID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "error retrieving the project")
			return
		}
		RespondWithJSON(w, http.StatusFound, project)
	}
}

//UpdateProjectController controller definition
func UpdateProjectController(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "error decoding the project struct")
	}
	params := mux.Vars(r)
	projectID, err := strconv.Atoi(params["id"])
	if err == nil {
		update, err := repository.UpdateProject(project.ProjectName, project.ProjectDescription, project.ProjectIsActive, project.ProjectBudget, project.ProjectMaterialCost, project.ProjectLaborCost, project.ProjectPlannedStartDate, project.ProjectPlannedEndDate, project.ProjectActualStartDate, project.ProjectActualEndDate, project.ProjectStatus, project.ProjectTotalHours, projectID)
		if err != nil {
			RespondWithError(w, http.StatusNotModified, "an error occurred while updating the project ")
			return
		}
		RespondWithJSON(w, http.StatusCreated, update)
	}
}

//CreateMilestoneController controller definition
func CreateMilestoneController(w http.ResponseWriter, r *http.Request) {
	var milestone models.Milestone

	params := mux.Vars(r)
	projectID, err := strconv.Atoi(params["id"])
	if err == nil {
		project, err := repository.FetchProject(projectID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "project not found create a project first then assign milestones to it")
			return
		}
		er := json.NewDecoder(r.Body).Decode(&milestone)
		if er != nil {
			RespondWithError(w, 500, "Invalid employee payload")
			return
		}
		//save a milestone
		response, err := repository.CreateMilestone(*project, milestone.MilestoneName, milestone.MilestoneStatus, milestone.MilestoneDeliverables, milestone.MilestoneDueDate, milestone.MilestoneAchievedDate, milestone.MilestoneTotalHours)
		if err != nil {
			RespondWithError(w, http.StatusNotImplemented, "could not create a milestone")
		}
		RespondWithJSON(w, http.StatusCreated, response)
	}
}

//GetMilestonesController controller definition
func GetMilestonesController(w http.ResponseWriter, r *http.Request) {
	milestones, err := repository.FetchMilestones()
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "error retrieving the milestones")
	}
	RespondWithJSON(w, http.StatusFound, milestones)
}

//GetMilestonesInProjectController controller definition
func GetMilestonesInProjectController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectID, er := strconv.Atoi(params["id"])
	if er == nil {
		milestones, err := repository.FetchMilestonesInProject(projectID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "error retrieving the milestones of that particular project")
		}
		RespondWithJSON(w, http.StatusFound, milestones)
	}
}

//GetMilestoneController controller definition
func GetMilestoneController(w http.ResponseWriter, r *http.Request) {
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

//UpdateMilestoneController controller definition
func UpdateMilestoneController(w http.ResponseWriter, r *http.Request) {
	var milestone models.Milestone
	err := json.NewDecoder(r.Body).Decode(&milestone)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "error decoding the milestone struct")
	}
	params := mux.Vars(r)
	milestoneID, err := strconv.Atoi(params["id"])
	if err == nil {
		update, err := repository.UpdateMilestone(milestone.MilestoneName, milestone.MilestoneStatus, milestone.MilestoneDeliverables, milestone.MilestoneDueDate, milestone.MilestoneAchievedDate, milestone.MilestoneTotalHours, milestoneID)
		if err != nil {
			RespondWithError(w, http.StatusNotModified, "an error occurred while updating the milestone")
			return
		}
		RespondWithJSON(w, http.StatusCreated, update)
	}
}
