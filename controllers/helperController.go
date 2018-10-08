package controllers

import (
	"encoding/json"
	"net/http"
)

//RespondWithJSON definition
func RespondWithJSON(w http.ResponseWriter, statusCode int, payLoad interface{}){
	response, _ := json.Marshal(payLoad)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

//RespondWithError definition
func RespondWithError(w http.ResponseWriter, statusCode int, message string)  {
	RespondWithJSON(w, statusCode, map[string]string{"An error ocurred":message})
}