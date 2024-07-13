package routes

import (
	"backend_case/pkg/controllers"

	"github.com/gorilla/mux"
)


var CreateGetTeamRouters = func (router *mux.Router)  {
	
	router.HandleFunc("/team",controllers.CreateTeam).Methods("POST")
	router.HandleFunc("/teams",controllers.GetTeam).Methods("GET")
	
}