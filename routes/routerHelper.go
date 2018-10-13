package routes

import "github.com/gorilla/mux"

func InitializeRoutes() *mux.Router  {
	router := mux.NewRouter().StrictSlash(false)

	//router for User
	router = SetUserRouter(router)
	//router for Project
	router = SetProjectRouter(router)
	//router for Task
	router = SetTaskRouter(router)

	return router
}