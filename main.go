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

var ownerId = 1
var m map[int]bool
var sendMessage = `{"ownerId":1,"number":23,"isEnabled":false}`

func init() {
	m = make(map[int]bool)
	m[23] = true
	m[0] = true
	m[1] = true
}

func main() {
	e := echo.New()
	hub := newHub()
	go hub.run()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middle.ServerHeader)
	e.POST("/login", handler.LoginHandler{}.Login)
	e.GET("/players", handler.PlayerHandler{}.GetPlayer, middle.Auth)
	e.PUT("/players", handler.PlayerHandler{}.UpdatePlayer, middle.Auth)
	e.GET("/teams", handler.TeamHandler{}.GetTeams, middle.Auth)
	e.GET("/ws", func(c echo.Context) error {
		serveWs(hub, c.Response().Writer, c.Request())
		return nil
	})
	e.Logger.Fatal(e.Start(":8000"))
}
