package controllers

import (
	"encoding/json"
	"github.com/Muart-C/webcbo/models"
	"net/http"
)

//CreateUserController controller definition
func CreateUserController(w http.ResponseWriter, r *http.Request)  {
	var user models.User

	//Decode the incoming user json data
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		RespondWithError(w,500,"Invalid user data on the body")
		return
	}
	//save a new user

	response,err:=models.CreateUser(user.UserName,user.Email,user.LastName,user.FirstName)

	if err != nil {
		RespondWithError(w,500,"An unexpected error occured")
	}
	resp, err := json.Marshal(response)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
