package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/shastri17/hplauction/auction"
)

type PlayerHandler struct {
}

type PlayerResp struct {
	Players []auction.Player `json:"players"`
}

func (h PlayerHandler) GetPlayer(e echo.Context) error {
	players := auction.GetPlayers()
	return e.JSON(200, auction.Response{
		Message: "SUCCESS", Data: PlayerResp{players},
	})
}

func (h PlayerHandler) UpdatePlayer(e echo.Context) error {

}
