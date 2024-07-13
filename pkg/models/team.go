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
	GoalFor int `json:"goal_for"`
	GoalAgainst int `json:"goal_against"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Team{})
}

func(b *Team) CreateTeam() *Team{
	db.Create(&b)
	return b
}

func GetAllTeams() []Team {
	var Teams []Team
	db.Find(&Teams)
	return Teams
}