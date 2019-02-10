package main

import (
	"encoding/json"
	"github.com/hplauction/auction"
	"github.com/hplauction/middle"
	"github.com/zopnow/z"
	"log"
	"net/http"
	"os"
	"strings"
)

//import (
//	"github.com/hplauction/auction"
//	"github.com/zopnow/z"
//	"os"
//)
//
//var handlers = z.Handlers{
//	"team":   auction.TeamHandler{},
//	"purse":  auction.PurseHandler{},
//	"player": auction.PlayerHandler{},
//	"login":  auction.LoginHandler{},
//}
func init() {
	conf := z.MysqlConfig{HostName: os.Getenv("DB_HOST"), Username: os.Getenv("DB_USER"), Password: os.Getenv("DB_PASSWORD"), Database: "hpl_auction"}
	z.ConnectMySQL(conf)
}

//
//func main() {
//	z.Service{Handlers: handlers, Config: &config}.Run()
//}
func main() {
	log.Println("Listening on 8000...")
	log.Fatal(http.ListenAndServe("localhost:8000", HandlerWithMiddleWares()))
}

func HandlerWithMiddleWares() http.Handler {
	handlerFunc := http.HandlerFunc(handler)
	// Add HTTP middlewares
	h := middle.Authenticator(handlerFunc) // Responds 404 for all invalid services
	return h
}
func handler(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")
	api := urlParts[1]
	var res interface{}
	var err error
	switch r.Method {
	case "GET":
		get := from(api)
		res, err = get.(auction.RestGet).Index(r)
	case "POST":
		get := from(api)
		res, err = get.(auction.RestPost).Create(r)
	case "PUT":
		get := from(api)
		res, err = get.(auction.RestPut).Update(r)
	}
	var resp middle.Response
	if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp.Code = 200
	resp.Status = "SUCCESS"
	resp.Data = res
	json.NewEncoder(w).Encode(resp)
}

func from(api string) interface{} {
	var m = map[string]interface{}{"player": auction.PlayerHandler{}, "team": auction.TeamHandler{}, "purse": auction.PurseHandler{}}
	if hand, ok := m[api]; ok {
		return hand
	}
	return nil
}
