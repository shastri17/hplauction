package auction

import (
	"github.com/shastri17/hplauction/db"
	"github.com/shastri17/hplauction/models"
)

func Authorise(token string) models.AuthResponse {
	var team models.Team
	if token == "" {
		return models.AuthResponse{Code: 400}
	}
	db.DB.Table("team").Where("token=?", token).First(&team)
	if team.Id > 0 {
		return models.AuthResponse{Code: 200, IsAdmin: team.IsAdmin, UserId: team.Id}
	}
	return models.AuthResponse{Code: 400}
}
