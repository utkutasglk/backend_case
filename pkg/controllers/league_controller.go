package controllers

import (
	"backend_case/pkg/config"
	"backend_case/pkg/models"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

// tum ligi oyna
func PlayLeague(writer http.ResponseWriter, request *http.Request) {

	// olusturulan takimlari arrayda topla
	var teams []models.Team
	config.DB.Find(&teams)

	// Mac fiksturunu olusturmak icin yeni bir metoda yonlendir
	schedule := createMatchFixture(len(teams))

	
	var matches []models.Match

	for week, games := range schedule {
		for _, game := range games {
			homeTeam := teams[game.home]
			awayTeam := teams[game.away]

			match := models.Match{
				HomeTeamID: homeTeam.Id,
				AwayTeamID: awayTeam.Id,
				Week:       week + 1,
			}

			homeGoals, awayGoals := simulateMatch(homeTeam, awayTeam)
			match.HomeGoals = homeGoals
			match.AwayGoals = awayGoals

			updateData(&teams[game.home], &teams[game.away], homeGoals, awayGoals)
			config.DB.Create(&match)
			matches = append(matches, match)
		}
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(matches)
}

type game struct {
	home int
	away int
}

func createMatchFixture(numTeams int) [][]game {

	// takim sayisi tek ise ++1 arttir
	if numTeams%2 != 0 {
		numTeams++
	}

	// icerde ve disarda oynanacak maclar icin hafta numarasini 2 ile carp
	numWeeks := (numTeams - 1) * 2
	gamesPerWeek := numTeams / 2

	// olusturdugumuz game struct sayesinde takimlari esle
	schedule := make([][]game, numWeeks)
	for week := 0; week < numWeeks; week++ {
		schedule[week] = make([]game, gamesPerWeek)
		for match := 0; match < gamesPerWeek; match++ {
			home := (week + match) % (numTeams - 1)
			away := (numTeams - 1 - match + week) % (numTeams - 1)
			if match == 0 {
				away = numTeams - 1
			}
			schedule[week][match] = game{home: home, away: away}
		}
	}

	return schedule
}

// mac skorunu takim guclerine ve random sayilara gore belirle
func simulateMatch(homeTeam models.Team, awayTeam models.Team) (int, int) {
	rand.Seed(time.Now().UnixNano())
	homeGoals := rand.Intn(homeTeam.Power/100 + 5)
	awayGoals := rand.Intn(awayTeam.Power/100 + 5)
	return homeGoals, awayGoals
}

// database guncelle
func updateData(homeTeam *models.Team, awayTeam *models.Team, homeGoals int, awayGoals int) {

	// skora gore puan guncellemesi yap
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
	
	// atilan ve yenilen golleri ekle
	homeTeam.GoalScore += homeGoals
	homeTeam.GoalConcede += awayGoals
	homeTeam.GoalDiff = homeTeam.GoalScore - homeTeam.GoalConcede

	awayTeam.GoalScore += awayGoals
	awayTeam.GoalConcede += homeGoals
	awayTeam.GoalDiff = awayTeam.GoalScore - awayTeam.GoalConcede

	// database kaydet
	config.DB.Save(homeTeam)
	config.DB.Save(awayTeam)
}