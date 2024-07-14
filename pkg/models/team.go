package models

import (
	"backend_case/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Team struct{
	Id int      `json:"id"`
	Name string `json:"name"`
	Power int `json:"power"`
	Points int  `json:"points"`
	Wins    int    `json:"wins"`
	Draws   int    `json:"draws"`
	Losses  int    `json:"losses"`
	GoalScore int `json:"goal_score"`
	GoalConcede int `json:"goal_concede"`
	GoalDiff int `json:"goal_diff"`

}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Team{})
}
