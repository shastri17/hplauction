package auction

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/hplauction/db"
	"io/ioutil"
	"net/http"
)

type LoginHandler struct {
}

type LoginResponse struct {
	IsAdmin     bool   `json:"isAdmin,omitempty"`
	AccessToken string `json:"accessToken"`
	IsOwner     bool   `json:"isOwner,omitempty"`
	OwnerId     int    `json:"ownerId,omitempty"`
}

func (l LoginHandler) Create(r *http.Request) interface{} {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &body)
	var team Team
	db.DB.Table("team").Where("username=? and password=?", body.Username, body.Password).First(&team)
	if team.Id == 0 {
		return Response{Code: 400, Message: "invalid username or password"}
	}
	token := createToken(team)
	team.Token = token
	db.DB.Table("team").Where("id=?", team.Id).Update(&team)
	var data LoginResponse
	if team.IsAdmin {
		data.IsAdmin = true
		data.AccessToken = token
	} else {
		data.IsOwner = true
		data.AccessToken = token
		data.OwnerId = team.Id
	}
	return Response{Code: 200, Message: "SUCCESS", Data: data}
}

func createToken(team Team) string {
	token := team.Username + team.Password
	h := sha1.New()
	h.Write([]byte(token))
	b := h.Sum(nil)
	token = fmt.Sprintf("%x", b)
	return token
}
