package auction

import "github.com/zopnow/z"

type AuthResponse struct {
	Code    int  `json:"code"`
	IsAdmin bool `json:"isOwner"`
	UserId  int  `json:"userId"`
}

func Authorise(params map[string]interface{}) AuthResponse {
	var team Team
	token := ""
	if accessToken, ok := params["accessToken"].(string); ok {
		token = accessToken
	}
	if token == "" {
		return AuthResponse{Code: 400}
	}
	z.DB.Table("team").Where("token=?", token).First(&team)
	if team.Id > 0 {
		return AuthResponse{Code: 200, IsAdmin: team.IsAdmin, UserId: team.Id}
	}
	return AuthResponse{Code: 400}
}
