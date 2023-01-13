package middle

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/shastri17/hplauction/auction"
	"strings"
)

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Custom-Header", "blah!!!")
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, x-requested-with, origin, X-API-VERSION")
		c.Response().Header().Set("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
		c.Response().Header().Set("Content-Type", "application/json")
		return next(c)
	}
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		s := strings.Split(token, " ")
		if len(s) == 2 {
			resp := auction.Authorise(s[1])
			if resp.Code != 200 {
				return errors.New("not authorised")
			}
			c.Set("isAdmin", resp.IsAdmin)
			c.Set("id", resp.UserId)
		}
		return next(c)
	}

}
