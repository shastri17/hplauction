package auction

import (
	"encoding/json"
	"github.com/zopnow/z"
	"io/ioutil"
	"net/http"
)

type PlayerHandler struct {
}


type PlayerResp struct {
	Players []Player `json:"players"`
}
func (p PlayerHandler) Index(r *http.Request) (interface{}, error) {
	players := getPlayers()
	return PlayerResp{players}, nil
}

func (p PlayerHandler) Create(r *http.Request) (interface{}, error) {
	var body struct {
		Id            int `json:"id"`
		BiddingAmount int `json:"bidAmount"`
		TeamId        int `json:"teamId"`
	}
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &body)
	z.DB.Exec("update player set bid_amount=?,team_id=?,is_sold=true where id=?", body.BiddingAmount, body.TeamId, body.Id)
	var team Team
	z.DB.Table("team").Where("id=?", body.TeamId).First(&team)
	team.PurseAmount = team.PurseAmount - body.BiddingAmount
	team.TotalPlayers = team.TotalPlayers + 1
	team.MaxBidAmount = team.PurseAmount - ((11 - team.TotalPlayers) * 100)
	z.DB.Table("team").Where("id=?", body.TeamId).Update(&team)
	players := getPlayers()
	return PlayerResp{players}, nil
}

func getPlayers() []Player {
	var players []Player
	z.DB.Table("player").Where("is_sold=false").Find(&players)
	return players
}
