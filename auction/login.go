package auction

import (
	"encoding/json"
	"github.com/zopnow/z"
	"io/ioutil"
	"net/http"
)

type LoginHandler struct {
}

func (l LoginHandler) Create(r *http.Request) (interface{}, error) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &body)
	var team Team
	z.DB.Table("team").Where("username=? and password=?", body.Username, body.Password).First(&team)
	if team.Id == 0 {
		return nil, z.Error{Code: http.StatusBadRequest, Message: "invalid username or password"}
	}
	token := createToken(team)
	team.Token = token
	z.DB.Table("team").Where("id=?", team.Id).Update(&team)
	var loginResp struct {
		IsAdmin     bool   `json:"isAdmin"`
		AccessToken string `json:"accessToken"`
	}
	loginResp.IsAdmin = team.IsAdmin
	loginResp.AccessToken = token
	return loginResp, nil
}

func createToken(team Team) string {
	token := team.Username + team.Password
	return token
}
