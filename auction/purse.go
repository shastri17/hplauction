package auction

import (
	"github.com/hplauction/db"
	"net/http"
)

type PurseHandler struct {
}

func (p PurseHandler) Index(r *http.Request) interface{} {
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
		db.DB.Table("team").Where("id!=?", id).Find(&teams)
		for _, v := range teams {
			re = append(re, res{v.PurseAmount, v.MaxBidAmount, v.TeamName})
		}
		return Response{Code: 200, Data: re}
	} else {
		var team Team
		db.DB.Table("team").Where("id=?", id).First(&team)
		return Response{Code: 200, Data: res{team.PurseAmount, team.MaxBidAmount, team.TeamName}}
	}
	return nil
}
