package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shastri17/hplauction/db"
	"github.com/shastri17/hplauction/handler"
	"github.com/shastri17/hplauction/middle"
)

func init() {
	db.DB, _ = db.GetMysqlDb()
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middle.ServerHeader)
	e.POST("/login", handler.LoginHandler{}.Login)
	e.GET("/players", handler.PlayerHandler{}.GetPlayer, middle.Auth)
	e.PUT("/players", handler.PlayerHandler{}.UpdatePlayer, middle.Auth)
	e.GET("/teams", handler.TeamHandler{}.GetTeams, middle.Auth)
	e.Logger.Fatal(e.Start(":8000"))
}
