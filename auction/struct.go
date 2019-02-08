package auction

import "github.com/guregu/null"

type Team struct{
	Id int `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	TeamName string `json:"name"`
	Logo string `json:"logo"`
	PurseAmount int `json:"purseAmount"`
	MaxBidAmount int `json:"maxBidAmount"`
	TotalPlayers int `json:"totalPlayers"`
	OwnersName string `json:"ownersName"`
	Icon1 string `json:"icon1"`
	Icon2 string `json:"icon2"`
	IsAdmin bool `json:"isAdmin"`
	Players []Player `json:"players" gorm:"-"`
}

type Player struct {
	Id int `json:"id"`
	Name string `json:"name"`
	NickName null.String `json:"nickName"`
	SkillArea string `json:"skillArea"`
	BattingHand null.String `json:"battingHand"`
	BowlingHand null.String `json:"bowlingHand"`
	MobileNumber int `json:"mobileNumber"`
	WhatsappNumber int `json:"whatsappNumber"`
	PreviouslyPlayedTeams string `json:"previouslyPlayedTeams"`
	Image string `json:"image"`
	BidAmount int `json:"bid_amount"`
	IsSold bool `json:"-"`
	TeamId int `json:"teamId"`
}