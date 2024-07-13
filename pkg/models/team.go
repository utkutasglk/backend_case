package models

import (
	"backend_case/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Team struct{
	Id int      `json:"id"`
	Name string `json:"name"`
	Strength int `json:"strength"`
	Points int  `json:"points"`
	Wins    int    `json:"wins"`
	Draws   int    `json:"draws"`
	Losses  int    `json:"losses"`
	GoalFor int `json:"goal_for"`
	GoalAgainst int `json:"goal_against"`
	GoalDifference int `json:"goal_difference"`

}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Team{})
}
