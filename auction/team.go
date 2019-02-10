package auction

import (
	"github.com/zopnow/z"
	"net/http"
)

type TeamHandler struct {
}

type TeamResponse struct {
	Teams interface{} `json:"team"`
}

func (t TeamHandler) Index(r *http.Request) (interface{}, error) {
	id := r.Context().Value("id").(int)
	ret := GetDetails(id)
	var re TeamResponse
	re.Teams = ret
	return re, nil
}

func GetDetails(id int) interface{} {
	var team Team
	var players []Player
	z.DB.Table("team").Where("id=?", id).First(&team)
	if team.IsAdmin {
		var teams []Team
		z.DB.Table("team").Where("is_admin!=true").Find(&teams)
		for i := range teams {
			var players []Player
			z.DB.Table("player").Where("team_id=?", teams[i].Id).Find(&players)
			teams[i].Players = players
		}
		return teams
	}
	z.DB.Table("player").Where("team_id=?", id).Find(&players)
	team.Players = players
	return team
}
