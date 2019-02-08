package auction

import (
	"github.com/zopnow/z"
	"net/http"
)

type PurseHandler struct {

}

func (p PurseHandler)Index(r *http.Request)(interface{},error){
	q:=r.URL.Query()
	teamId:=q.Get("teamId")
	var team Team
	z.DB.Table("team").Where("id=?",teamId).First(&team)
	var res struct{
		Purse int `json:"purse"`
		MaxBidAmount int `json:"maxBidAmount"`
	}
	res.MaxBidAmount=team.MaxBidAmount
	res.Purse=team.PurseAmount
	return res,nil
}