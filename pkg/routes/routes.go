package routes

import (
	"backend_case/pkg/controllers"

	"github.com/gorilla/mux"
)


var RegisterRoutes = func (router *mux.Router)  {
	
	router.HandleFunc("/team",controllers.CreateTeam).Methods("POST")
	router.HandleFunc("/teams",controllers.GetTeam).Methods("GET")
	router.HandleFunc("/play-league",controllers.PlayLeague).Methods("POST")
	router.HandleFunc("/matches",controllers.GetMatches).Methods("GET")
	router.HandleFunc("/matches/{week}", controllers.GetMatchesByWeek).Methods("GET")



	
}