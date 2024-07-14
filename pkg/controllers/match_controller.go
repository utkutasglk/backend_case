package controllers

import (
	"backend_case/pkg/config"
	"backend_case/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetMatches(w http.ResponseWriter, r *http.Request){
	var matches []models.Match
	config.DB.Preload("HomeTeam").Preload("AwayTeam").Find(&matches)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matches)

}

func GetMatchesByWeek(w http.ResponseWriter, r *http.Request) {
	var matches []models.Match
	vars := mux.Vars(r)
	week, _ := strconv.Atoi(vars["week"])

	config.DB.Where("week = ?", week).Preload("HomeTeam").Preload("AwayTeam").Find(&matches)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matches)
}