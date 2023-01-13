package auction

import (
	"errors"
	"github.com/shastri17/hplauction/db"
	"github.com/shastri17/hplauction/models"
)

type PlayerHandler struct {
}

type PlayerResp struct {
	Players []models.Player `json:"players"`
}

func UpdatePlayer(body models.PlayerUpdatedRequest, isAdmin bool) (models.Player, error) {
	if !isAdmin {
		return models.Player{}, errors.New("not a admin")
	}
	var player models.Player
	db.DB.Table("player").Where("id=?", body.Id).First(&player)
	if player.IsSold {
		return models.Player{}, errors.New("player already sold")
	}
	player.IsSold = true
	player.TeamId = body.TeamId
	player.BidAmount = body.BiddingAmount

	var team models.Team
	db.DB.Table("team").Where("id=?", body.TeamId).First(&team)
	team.TotalPlayers = team.TotalPlayers + 1
	team.PurseAmount = team.PurseAmount - body.BiddingAmount
	team.MaxBidAmount = team.PurseAmount - ((11 - team.TotalPlayers) * 100)
	if team.PurseAmount <= team.MaxBidAmount {
		team.MaxBidAmount = team.PurseAmount
	}

	if team.PurseAmount <= 0 {
		team.PurseAmount = 0
		team.MaxBidAmount = 0
	}
	db.DB.Table("team").Save(&team)
	player.TeamName = team.TeamName
	db.DB.Table("player").Where("id=?", body.Id).Update(&player)
	return player, nil
}

func GetPlayers() []models.Player {
	var players []models.Player
	db.DB.Table("player").Find(&players)
	return players
}
