package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/shastri17/hplauction/auction"
	"github.com/shastri17/hplauction/models"
	"io/ioutil"
)

type LoginHandler struct {
}

func (l LoginHandler) Login(c echo.Context) error {
	var body models.LoginRequest
	b, _ := ioutil.ReadAll(c.Request().Body)
	json.Unmarshal(b, &body)
	resp, err := auction.Login(body)
	if err != nil {
		return c.JSON(400, models.Response{Code: 200, Message: err.Error()})
	}
	return c.JSON(200, models.Response{Code: 200, Message: "success", Data: resp})
}
