package auction

import (
	"github.com/zopnow/z"
	"net/http"
)

type PurseHandler struct {
}

func (p PurseHandler) Index(r *http.Request) (interface{}, error) {
	id := r.Context().Value("id")
	isAdmin := r.Context().Value("isAdmin")
	type res struct {
		Purse        int    `json:"purse"`
		MaxBidAmount int    `json:"maxBidAmount"`
		TeamName     string `json:"teamName"`
	}
	if isAdmin.(bool) {
		var teams []Team
		var re []res
		z.DB.Table("team").Where("id!=?", id).Find(&teams)
		for _, v := range teams {
			re = append(re, res{v.PurseAmount, v.MaxBidAmount, v.TeamName})
		}
		return re, nil
	} else {
		var team Team
		z.DB.Table("team").Where("id=?", id).First(&team)
		return res{team.PurseAmount, team.MaxBidAmount, team.TeamName}, nil
	}
	return nil, nil
}
