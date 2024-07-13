package controllers

import (
	"backend_case/pkg/config"
	"backend_case/pkg/models"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)


func PlayLeague(writer http.ResponseWriter, request *http.Request) {

	var teams []models.Team
	config.DB.Find(&teams)

	var matches []models.Match
	weeks := len(teams) - 1
	for i := 0; i < weeks; i++ {
		for j := 0; j < len(teams); j++ {
			if j != i {
				match := models.Match{
					HomeTeamID: teams[i].Id,
					AwayTeamID: teams[j].Id,
					Week:       i + 1,
				}
				homeGoals, awayGoals := simulateMatch(teams[i], teams[j])
				match.HomeGoals = homeGoals
				match.AwayGoals = awayGoals

				config.DB.Create(&match)
				matches = append(matches, match)
			}
		}
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(matches)
	
}

func simulateMatch(homeTeam models.Team, awayTeam models.Team) (int, int) {
	rand.Seed(time.Now().UnixNano())
	homeGoals := rand.Intn(homeTeam.Strength+1)
	awayGoals := rand.Intn(awayTeam.Strength+1)
	return homeGoals, awayGoals
}

