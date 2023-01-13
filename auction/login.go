package auction

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/shastri17/hplauction/db"
	"github.com/shastri17/hplauction/models"
)

func Login(body models.LoginRequest) (models.LoginResponse, error) {
	var team models.Team
	db.DB.Table("team").Where("username=? and password=?", body.Username, body.Password).First(&team)
	if team.Id == 0 {
		return models.LoginResponse{}, errors.New("invalid username or password")
	}
	token := createToken(team)
	team.Token = token
	db.DB.Table("team").Where("id=?", team.Id).Update(&team)
	var data models.LoginResponse
	if team.IsAdmin {
		data.IsAdmin = true
		data.AccessToken = token
	} else {
		data.IsOwner = true
		data.AccessToken = token
		data.OwnerId = team.Id
	}
	return data, nil
}

func createToken(team models.Team) string {
	token := team.Username + team.Password
	h := sha1.New()
	h.Write([]byte(token))
	b := h.Sum(nil)
	token = fmt.Sprintf("%x", b)
	return token
}
