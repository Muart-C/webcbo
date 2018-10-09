package controllers

import (
	"encoding/json"
	"github.com/Muart-C/webcbo/models"
	"github.com/Muart-C/webcbo/repositories"
	"net/http"
)
//CreateUserController controller definition
func CreateUserController(w http.ResponseWriter, r *http.Request)  {
	var userData models.User
	//Decode the incoming user json data
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		RespondWithError(w,500,"Invalid user data")
		return
	}
	//save a new user
	repositories.CreateUserRepository(&userData)

	resp, err := json.Marshal(userData)
	if err != nil {
		RespondWithError(w,500,"An unexpected error occured")
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}//Test for working
