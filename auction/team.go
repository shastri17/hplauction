package auction

import (
	"github.com/shastri17/hplauction/db"
	"github.com/shastri17/hplauction/models"
)

type TeamHandler struct {
}

type TeamResponse struct {
	Teams interface{} `json:"team"`
}

func GetDetails(id int) interface{} {
	var team models.Team
	var players []models.Player
	db.DB.Table("team").Where("id=?", id).First(&team)
	if team.IsAdmin {
		var teams []models.Team
		db.DB.Table("team").Where("is_admin!=true").Find(&teams)
		for i := range teams {
			var players []models.Player
			db.DB.Table("player").Where("team_id=?", teams[i].Id).Find(&players)
			teams[i].Players = players
		}
		return TeamResponse{teams}
	}
	db.DB.Table("player").Where("team_id=?", id).Find(&players)
	team.Players = players
	return TeamResponse{team}
}
