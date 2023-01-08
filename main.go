package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shastri17/hplauction/auction"
	"github.com/shastri17/hplauction/db"
	"github.com/shastri17/hplauction/handler"
	"github.com/shastri17/hplauction/middle"
	"log"
	"net/http"
	"strings"
)

func init() {
	db.DB, _ = db.GetMysqlDb()
}

func main() {
	log.Println("Listening on 8000")
	log.Fatal(http.ListenAndServe(":8000", Handler()))
}

func Handler() http.Handler {
	handlerFunc := http.HandlerFunc(hplHandler)
	h := middle.Authenticator(handlerFunc)
	return h
}
func hplHandler(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")
	api := urlParts[1]
	var res interface{}
	switch r.Method {
	case "GET":
		get := from(api)
		res = get.(auction.RestGet).Index(r)
	case "POST":
		get := from(api)
		res = get.(auction.RestPost).Create(r)
	case "PUT":
		get := from(api)
		res = get.(auction.RestPut).Update(r)
	}
	json.NewEncoder(w).Encode(res)
}

func from(api string) interface{} {
	var m = map[string]interface{}{"players": auction.PlayerHandler{}, "teams": auction.TeamHandler{}, "purse": auction.PurseHandler{}}
	if hand, ok := m[api]; ok {
		return hand
	}
	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/players", handler.PlayerHandler{}.GetPlayer)
	e.PUT("/players", handler.PlayerHandler{}.UpdatePlayer)
}
