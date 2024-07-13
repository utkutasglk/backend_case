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

				updateTeamStats(&teams[i], &teams[j], homeGoals, awayGoals)
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
func updateTeamStats(homeTeam *models.Team, awayTeam *models.Team, homeGoals int, awayGoals int) {
	if homeGoals > awayGoals {
		homeTeam.Points += 3
		homeTeam.Wins += 1
		awayTeam.Losses += 1
	} else if homeGoals < awayGoals {
		awayTeam.Points += 3
		awayTeam.Wins += 1
		homeTeam.Losses += 1
	} else {
		homeTeam.Points += 1
		awayTeam.Points += 1
		homeTeam.Draws += 1
		awayTeam.Draws += 1
	}

	homeTeam.GoalFor += homeGoals
	homeTeam.GoalAgainst += awayGoals
	homeTeam.GoalDifference = homeTeam.GoalFor - homeTeam.GoalAgainst

	awayTeam.GoalFor += awayGoals
	awayTeam.GoalAgainst += homeGoals
	awayTeam.GoalDifference = awayTeam.GoalFor - awayTeam.GoalAgainst

	config.DB.Save(homeTeam)
	config.DB.Save(awayTeam)
}
