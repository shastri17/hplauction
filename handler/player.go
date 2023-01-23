package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/shastri17/hplauction/auction"
	"github.com/shastri17/hplauction/models"
	"io/ioutil"
)

type PlayerHandler struct {
}

type PlayerResp struct {
	Players []models.Player `json:"players"`
}

func (h PlayerHandler) GetPlayer(e echo.Context) error {
	players := auction.GetPlayers()
	return e.JSON(200, models.Response{
		Message: "SUCCESS", Data: PlayerResp{players},
	})
}

func (h PlayerHandler) UpdatePlayer(e echo.Context) error {
	var resp models.Response
	var body models.PlayerUpdatedRequest
	b, _ := ioutil.ReadAll(e.Request().Body)
	json.Unmarshal(b, &body)
	isAdmin := e.Get("isAdmin").(bool)
	response, err := auction.UpdatePlayer(body, isAdmin)
	if err != nil {
		resp.Code = 200
		resp.Data = nil
		resp.Message = fmt.Sprintln(err)
		return e.JSON(resp.Code, resp)
	}
	resp.Code = 200
	resp.Data = response
	resp.Message = "Success"
	return e.JSON(resp.Code, resp)
}
