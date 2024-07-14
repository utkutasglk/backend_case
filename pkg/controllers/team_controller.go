package controllers

import (
	"backend_case/pkg/config"
	"backend_case/pkg/models"
	"encoding/json"
	"net/http"
)

func CreateTeam(writer http.ResponseWriter, request *http.Request){
	
	var team models.Team
	json.NewDecoder(request.Body).Decode(&team)
	config.DB.Create(&team)

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(team)
	
}

func GetTeams(writer http.ResponseWriter, request *http.Request){

	var teams []models.Team
	config.DB.Find(&teams)

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(teams)
	

}