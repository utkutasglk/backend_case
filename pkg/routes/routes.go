package routes

import (
	"backend_case/pkg/controllers"

	"github.com/gorilla/mux"
)


var RegisterRoutes = func (router *mux.Router)  {
	
	router.HandleFunc("/team",controllers.CreateTeam).Methods("POST")
	router.HandleFunc("/teams",controllers.GetTeam).Methods("GET")
	router.HandleFunc("/play-league",controllers.PlayLeague).Methods("POST")

	
}