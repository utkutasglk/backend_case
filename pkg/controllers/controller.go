package controllers

import (
	"backend_case/pkg/models"
	"backend_case/pkg/utils"
	"encoding/json"
	"net/http"
)

func GetTeam(writer http.ResponseWriter, request *http.Request){
  newTeams := models.GetAllTeams()
	res, _ := json.Marshal(newTeams)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)

}

func CreateTeam(writer http.ResponseWriter, request *http.Request){

	CreateTeam := &models.Team{}
	utils.ParseBody(request, CreateTeam)
	team := CreateTeam.CreateTeam()
	res, _ := json.Marshal(team)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write(res)
	
}