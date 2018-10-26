package controllers
//
//import (
//	"encoding/json"
//	"github.com/Muart-C/webcbo/models"
//	"github.com/Muart-C/webcbo/repository"
//	"github.com/gorilla/mux"
//	"net/http"
//	"strconv"
//)
//
////CreateRoleController controller definition
//func CreateRoleController(w http.ResponseWriter, r *http.Request) {
//	var role models.Role
//	//Decode the incoming role json data
//	err := json.NewDecoder(r.Body).Decode(&role)
//	if err != nil {
//		RespondWithError(w, 500, "Invalid role data on the body")
//		return
//	}
//	//save a new role
//
//	response, err := repository.CreateRole(role.RoleName)
//
//	if err != nil {
//		RespondWithError(w, 500, "An unexpected error occured")
//	}
//	RespondWithJSON(w, http.StatusCreated, response)
//}
//
////GetRolesController controller definition
//func GetRolesController(w http.ResponseWriter, r *http.Request) {
//	roles, err := repository.FetchRoles()
//	if err != nil {
//		RespondWithError(w, http.StatusNotFound, err.Error())
//	}
//	RespondWithJSON(w, http.StatusFound, roles)
//}
//
////GetRoleController controller definition
//func GetRoleController(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	roleID, err := strconv.Atoi(params["id"])
//	if err == nil {
//		role, err := repository.FetchRole(roleID)
//		if err != nil {
//			RespondWithError(w, http.StatusNotFound, "role not found")
//			return
//		}
//		RespondWithJSON(w, http.StatusFound, role)
//	}
//}
//
////UpdateRoleController controller definition
//func UpdateRoleController(w http.ResponseWriter, r *http.Request) {
//
//	var role models.Role
//	//Decode the incoming role json data
//	err := json.NewDecoder(r.Body).Decode(&role)
//	if err != nil {
//		RespondWithError(w, 500, "Invalid role data on the body")
//		return
//	}
//	params := mux.Vars(r)
//	roleID, err := strconv.Atoi(params["id"])
//	if err == nil {
//		update, err := repository.UpdateRole(role.RoleName, roleID)
//		if err != nil {
//			RespondWithError(w, http.StatusNotModified, "Error updating role")
//			return
//		}
//		RespondWithJSON(w, http.StatusCreated, update)
//	}
//}
//
////DeleteRoleController controller definition
//func DeleteRoleController(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	roleID, err := strconv.Atoi(params["id"])
//	if err == nil {
//		err := repository.DeleteRole(roleID)
//		if err != nil {
//			RespondWithError(w, http.StatusNotFound, "role not found")
//			return
//		}
//		RespondWithJSON(w, http.StatusFound, "role Deleted successfully")
//	}
//}
//
////CreateTeamController controller definition
//func CreateTeamController(w http.ResponseWriter, r *http.Request) {
//	var team models.Team
//	//Decode the incoming Team json data
//	err := json.NewDecoder(r.Body).Decode(&team)
//	if err != nil {
//		RespondWithError(w, 500, "Invalid Team data on the body")
//		return
//	}
//	//save a new Team
//
//	response, err := repository.CreateTeam(team.TeamName)
//
//	if err != nil {
//		RespondWithError(w, 500, "An unexpected error occured")
//	}
//	RespondWithJSON(w, http.StatusCreated, response)
//}
//
////GetTeamsController controller definition
//func GetTeamsController(w http.ResponseWriter, r *http.Request) {
//	teams, err := repository.FetchTeams()
//	if err != nil {
//		RespondWithError(w, http.StatusNotFound, "No teams found")
//	}
//	RespondWithJSON(w, http.StatusFound, teams)
//}
//
////GetTeamController controller definition
//func GetTeamController(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	teamID, err := strconv.Atoi(params["id"])
//	if err == nil {
//		team, err := repository.FetchTeam(teamID)
//		if err != nil {
//			RespondWithError(w, http.StatusNotFound, "Team not found")
//			return
//		}
//		RespondWithJSON(w, http.StatusFound, team)
//	}
//}
//
////UpdateTeamController controller definition
//func UpdateTeamController(w http.ResponseWriter, r *http.Request) {
//
//	var team models.Team
//	//Decode the incoming Team json data
//	err := json.NewDecoder(r.Body).Decode(&team)
//	if err != nil {
//		RespondWithError(w, 500, "Invalid Team data on the body")
//		return
//	}
//	params := mux.Vars(r)
//	teamID, err := strconv.Atoi(params["id"])
//	if err == nil {
//		update, err := repository.UpdateTeam(team.TeamName, teamID)
//		if err != nil {
//			RespondWithError(w, http.StatusNotModified, "Error updating Team")
//			return
//		}
//		RespondWithJSON(w, http.StatusCreated, update)
//	}
//}
//
////DeleteTeamController controller definition
//func DeleteTeamController(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	teamID, err := strconv.Atoi(params["id"])
//	if err == nil {
//		err := repository.DeleteTeam(teamID)
//		if err != nil {
//			RespondWithError(w, http.StatusNotFound, "Team not found")
//			return
//		}
//		RespondWithJSON(w, http.StatusFound, "Team Deleted successfully")
//	}
//}
