package models

import "github.com/guregu/null"

type PlayerUpdatedRequest struct {
	Id            int `json:"id"`
	BiddingAmount int `json:"bidAmount"`
	TeamId        int `json:"teamId"`
}
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Team struct {
	Id           int      `json:"id"`
	Username     string   `json:"-"`
	Password     string   `json:"-"`
	TeamName     string   `json:"name"`
	Logo         string   `json:"logo"`
	PurseAmount  int      `json:"purseAmount"`
	MaxBidAmount int      `json:"maxBidAmount"`
	TotalPlayers int      `json:"totalPlayers"`
	OwnersName   string   `json:"ownersName"`
	Icon1        string   `json:"icon1"`
	Icon2        string   `json:"icon2"`
	IsAdmin      bool     `json:"-"`
	Token        string   `json:"-"`
	Players      []Player `json:"players" gorm:"-"`
}

type Player struct {
	Id                    int         `json:"id"`
	Name                  string      `json:"name"`
	NickName              null.String `json:"nickName"`
	SkillArea             string      `json:"skillArea"`
	BattingHand           null.String `json:"battingHand"`
	BowlingHand           null.String `json:"bowlingHand"`
	MobileNumber          int         `json:"mobileNumber"`
	WhatsappNumber        int         `json:"whatsappNumber"`
	PreviouslyPlayedTeams string      `json:"previouslyPlayedTeams"`
	Image                 string      `json:"image"`
	BidAmount             int         `json:"soldAmount"`
	IsSold                bool        `json:"isSold"`
	TeamId                int         `json:"-"`
	TeamName              string      `json:"teamName,omitempty"`
}

type AuthResponse struct {
	Code    int  `json:"code"`
	IsAdmin bool `json:"isOwner"`
	UserId  int  `json:"userId"`
}
type LoginResponse struct {
	IsAdmin     bool   `json:"isAdmin,omitempty"`
	AccessToken string `json:"accessToken"`
	IsOwner     bool   `json:"isOwner,omitempty"`
	OwnerId     int    `json:"ownerId,omitempty"`
}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
