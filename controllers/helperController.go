package controllers

import (
	"encoding/json"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, statusCode int, payLoad interface{}){
	response, _ := json.Marshal(payLoad)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, statusCode int, message string)  {
	RespondWithJson(w, statusCode, map[string]string{"An error ocurred":message})
}