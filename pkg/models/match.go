package models

import (
	"backend_case/pkg/config"

	"gorm.io/gorm"
)

var DB *gorm.DB

type Match struct {
	Id          int `json:"id"`
	HomeTeamID  int `json:"home_team_id"`
	AwayTeamID  int `json:"away_team_id"`
	HomeGoals   int `json:"home_goals"`
	AwayGoals   int `json:"away_goals"`
	Week        int `json:"week"`
	HomeTeam    Team `gorm:"foreignKey:HomeTeamID"`
	AwayTeam    Team `gorm:"foreignKey:AwayTeamID"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Match{})
}