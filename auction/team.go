package auction

import (
	"encoding/json"
	"github.com/zopnow/z"
	"io/ioutil"
	"net/http"
	"strconv"
)

type TeamHandler struct {
}

type TeamResponse struct {
	Teams interface{} `json:"team"`
}

func (t TeamHandler) Index(r *http.Request) (interface{}, error) {
	q := r.URL.Query()
	id := q.Get("id")
	if id == "" {
		return nil, z.Error{Code: http.StatusBadRequest, Message: "please send id"}
	}
	ret := GetDetails(id)
	var re TeamResponse
	re.Teams=ret
	return re, nil
}

func GetDetails(id string) interface{} {
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

func (t TeamHandler) Create(r *http.Request) (interface{}, error) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &body)
	var team Team
	z.DB.Table("team").Where("username=? and password=?", body.Username, body.Password).First(&team)
	if team.Id==0{
		return nil,z.Error{http.StatusBadRequest,"invalid username or password"}
	}
	ret := GetDetails(strconv.Itoa(team.Id))
	var re TeamResponse
	re.Teams=ret
	return re, nil
}
