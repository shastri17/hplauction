package auction

import (
	"encoding/json"
	"fmt"
	"github.com/hplauction/db"
	"io/ioutil"
	"net/http"
)

type PlayerHandler struct {
}

type PlayerResp struct {
	Players []Player `json:"players"`
}

func (p PlayerHandler) Index(r *http.Request) interface{} {
	players := getPlayers()
	return Response{Code: 200, Message: "SUCCESS", Data: PlayerResp{players}}
}

func (p PlayerHandler) Update(r *http.Request) interface{} {
	var body struct {
		Id            int `json:"id"`
		BiddingAmount int `json:"bidAmount"`
		TeamId        int `json:"teamId"`
	}
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &body)
	isAdmin := r.Context().Value("isAdmin").(bool)
	if !isAdmin {
		return Response{Code: 400, Message: "Permission denied"}
	}
	var player Player
	db.DB.Table("player").Where("id=?", body.Id).First(&player)
	if player.IsSold {
		return Response{Code: 400, Message: "Player is already sold to " + fmt.Sprint(player.TeamName)}
	}
	player.IsSold = true
	player.TeamId = body.TeamId
	player.BidAmount = body.BiddingAmount

	var team Team
	db.DB.Table("team").Where("id=?", body.TeamId).First(&team)
	team.PurseAmount = team.PurseAmount - body.BiddingAmount
	team.TotalPlayers = team.TotalPlayers + 1
	team.MaxBidAmount = team.PurseAmount - ((11 - team.TotalPlayers) * 100)
	db.DB.Table("team").Where("id=?", body.TeamId).Update(&team)
	player.TeamName = team.TeamName
	db.DB.Table("player").Where("id=?", body.Id).Update(&player)
	return Response{Code: 200, Message: "SUCCESS", Data: player}
}

func getPlayers() []Player {
	var players []Player
	db.DB.Table("player").Find(&players)
	return players
}
