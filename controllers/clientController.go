package controllers

import (
	"encoding/json"
	"github.com/Muart-C/webcbo/models"
	"github.com/Muart-C/webcbo/repository"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//CreateClientController controller definition
func CreateClientController(w http.ResponseWriter, r *http.Request) {
	var client models.Client
	//Decode the incoming Client json data
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		RespondWithError(w, 500, "Invalid Client data on the body")
		return
	}
	//save a new Client

	response, err := repository.CreateClient(client.ClientName, client.ClientDescription, client.ClientCounty, client.ClientIndustrySector, client.ClientCity, client.ClientPhone, client.ClientZip)

	if err != nil {
		RespondWithError(w, 500, "An unexpected error occured")
	}
	RespondWithJSON(w, http.StatusCreated, response)
}

//GetClientsController controller definition
func GetClientsController(w http.ResponseWriter, r *http.Request) {
	clients, err := repository.FetchClients()
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "An Error occurred no clients were found")
	}
	RespondWithJSON(w, http.StatusFound, clients)
}

//GetClientController controller definition
func GetClientController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	clientID, err := strconv.Atoi(params["id"])
	if err == nil {
		client, err := repository.FetchClient(clientID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "Client not found")
			return
		}
		RespondWithJSON(w, http.StatusFound, client)
	}
}

//UpdateClientController controller definition
func UpdateClientController(w http.ResponseWriter, r *http.Request) {

	var client models.Client
	//Decode the incoming Client json data
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		RespondWithError(w, 500, "Invalid Client data on the body")
		return
	}
	params := mux.Vars(r)
	clientID, err := strconv.Atoi(params["id"])
	if err == nil {
		update, err := repository.UpdateClient(client.ClientName, client.ClientDescription, client.ClientCounty, client.ClientIndustrySector, client.ClientCity, client.ClientPhone, client.ClientZip, clientID)
		if err != nil {
			RespondWithError(w, http.StatusNotModified, "Error updating Client")
			return
		}
		RespondWithJSON(w, http.StatusCreated, update)
	}
}

//DeleteClientController controller definition
func DeleteClientController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	clientID, err := strconv.Atoi(params["id"])
	if err == nil {
		err := repository.DeleteClient(clientID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "Client not found")
			return
		}
		RespondWithJSON(w, http.StatusFound, "Client Deleted successfully")
	}
}
