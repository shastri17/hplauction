package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/shastri17/hplauction/auction"
	"github.com/shastri17/hplauction/models"
)

type TeamHandler struct {
}

func (t TeamHandler) GetTeams(e echo.Context) error {
	id := e.Get("id").(int)
	ret := auction.GetDetails(id)
	resp := models.Response{Code: 200, Message: "success", Data: ret}
	return e.JSON(resp.Code, resp)
}
