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
func CreateProjectController(w http.ResponseWriter, r *http.Request)  {
	var project  models.Project

	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		RespondWithError(w,http.StatusBadRequest,"Could not decode the project json")
	}
	//save a project
	response, err := repository.CreateProject(project.ProjectName,project.ProjectDescription,project.ProjectBudget,project.ProjectMaterialCost,project.ProjectLaborCost,project.ProjectPlannedStartDate,project.ProjectPlannedEndDate,project.ProjectActualStartDate,project.ProjectActualEndDate,project.ProjectTotalHours)
	if err != nil {
		RespondWithError(w,http.StatusNotImplemented,"could not create a user")
	}
	RespondWithJSON(w,http.StatusCreated,response)
}

//GetProjectsController controller definition
func GetProjectsController(w http.ResponseWriter, r *http.Request){
	projects, err := repository.FetchProjects()
	if err != nil {
		RespondWithError(w,http.StatusNotFound,"error retrieving the projects")
	}
	RespondWithJSON(w,http.StatusFound,projects)
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
func UpdateProjectController(w http.ResponseWriter, r *http.Request)  {
	var project models.Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		RespondWithError(w,http.StatusInternalServerError,"error decoding the project struct")
	}
	params := mux.Vars(r)
	projectID,err := strconv.Atoi(params["id"])
	if err == nil {
		update,err := repository.UpdateProject(project.ProjectName,project.ProjectDescription,project.ProjectBudget,project.ProjectMaterialCost,project.ProjectLaborCost,project.ProjectPlannedStartDate,project.ProjectPlannedEndDate,project.ProjectActualStartDate,project.ProjectActualEndDate,project.ProjectTotalHours, projectID)
		if err != nil {
			RespondWithError(w,http.StatusNotModified,"an error occurred while updating the project ")
			return
		}
		RespondWithJSON(w,http.StatusCreated,update)
	}
}

//CreateProjectStatusController controller definition
func CreateProjectStatusController(w http.ResponseWriter, r *http.Request)  {
	var status  models.ProjectStatus

	err := json.NewDecoder(r.Body).Decode(&status)
	if err != nil {
		RespondWithError(w,http.StatusBadRequest,"Could not decode the status json")
	}
	//save a status
	response, err := repository.CreateProjectStatus(status.ProjectStatus,status.ProjectIsActive)
	if err != nil {
		RespondWithError(w,http.StatusNotImplemented,"could not create a status")
	}
	RespondWithJSON(w,http.StatusCreated,response)
}


//UpdateProjectStatusController controller definition
func UpdateProjectStatusController(w http.ResponseWriter, r *http.Request)  {
	var status models.ProjectStatus
	err := json.NewDecoder(r.Body).Decode(&status)
	if err != nil {
		RespondWithError(w,http.StatusInternalServerError,"error decoding the status struct")
	}
	params := mux.Vars(r)
	statusID,err := strconv.Atoi(params["id"])
	if err == nil {
		update,err := repository.UpdateProjectStatus(status.ProjectStatus,status.ProjectIsActive, statusID)
		if err != nil {
			RespondWithError(w,http.StatusNotModified,"an error occurred while updating the project status")
			return
		}
		RespondWithJSON(w,http.StatusCreated,update)
	}
}

//CreateMilestoneStatusController controller definition
func CreateMilestoneStatusController(w http.ResponseWriter, r *http.Request)  {
	var status  models.MilestoneStatus

	err := json.NewDecoder(r.Body).Decode(&status)
	if err != nil {
		RespondWithError(w,http.StatusBadRequest,"Could not decode the status json")
	}
	//save a status
	response, err := repository.CreateMilestoneStatus(status.MilestoneStatus)
	if err != nil {
		RespondWithError(w,http.StatusNotImplemented,"could not create a status")
	}
	RespondWithJSON(w,http.StatusCreated,response)
}


//UpdateMilestoneStatusController controller definition
func UpdateMilestoneStatusController(w http.ResponseWriter, r *http.Request)  {
	var status models.MilestoneStatus
	err := json.NewDecoder(r.Body).Decode(&status)
	if err != nil {
		RespondWithError(w,http.StatusInternalServerError,"error decoding the status struct")
	}
	params := mux.Vars(r)
	statusID,err := strconv.Atoi(params["id"])
	if err == nil {
		update,err := repository.UpdateMilestoneStatus(status.MilestoneStatus,statusID)
		if err != nil {
			RespondWithError(w,http.StatusNotModified,"an error occurred while updating the milestone status")
			return
		}
		RespondWithJSON(w,http.StatusCreated,update)
	}
}

//CreateMilestoneController controller definition
func CreateMilestoneController(w http.ResponseWriter, r *http.Request)  {
	var milestone models.Milestone

	err := json.NewDecoder(r.Body).Decode(&milestone)
	if err != nil {
		RespondWithError(w,http.StatusBadRequest,"Could not decode the milestone json")
	}
	//save a milestone
	response, err := repository.CreateMilestone(milestone.MilestoneName,milestone.MilestoneDeliverables,milestone.MilestoneDueDate,milestone.MilestoneAchievedDate,milestone.MilestoneTotalHours)
	if err != nil {
		RespondWithError(w,http.StatusNotImplemented,"could not create a milestone")
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

//UpdateMilestoneController controller definition
func UpdateMilestoneController(w http.ResponseWriter, r *http.Request)  {
	var milestone models.Milestone
	err := json.NewDecoder(r.Body).Decode(&milestone)
	if err != nil {
		RespondWithError(w,http.StatusInternalServerError,"error decoding the milestone struct")
	}
	params := mux.Vars(r)
	milestoneID,err := strconv.Atoi(params["id"])
	if err == nil {
		update,err := repository.UpdateMilestone(milestone.MilestoneName,milestone.MilestoneDeliverables,milestone.MilestoneDueDate,milestone.MilestoneAchievedDate,milestone.MilestoneTotalHours, milestoneID)
		if err != nil {
			RespondWithError(w,http.StatusNotModified,"an error occurred while updating the milestone")
			return
		}
		RespondWithJSON(w,http.StatusCreated,update)
	}
}