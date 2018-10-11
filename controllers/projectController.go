package controllers

import (
	"encoding/json"
	"github.com/Muart-C/webcbo/models"
	"github.com/Muart-C/webcbo/repository"
	"net/http"
)

//CreateProjectController controller definition
func CreateProjectController(w http.ResponseWriter, r *http.Request)  {
	var project  models.Project

	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		RespondWithError(w,http.StatusBadRequest,"Could not decode the project json")
	}
	//save a project
	response, err := repository.CreateProject(project.ProjectName,project.ProjectDescription,project.ProjectMaterialCost,project.ProjectLaborCost,project.ProjectPlannedStartDate,project.ProjectPlannedEndDate,project.ProjectActualStartDate,project.ProjectActualEndDate,project.ProjectTotalHours)
	if err != nil {
		RespondWithError(w,http.StatusNotImplemented,"could not create a user")
	}
	RespondWithJSON(w,http.StatusCreated,response)
}

