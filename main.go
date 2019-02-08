package main

import (
	"github.com/hplauction/auction"
	"github.com/zopnow/z"
	"os"
)

var handlers = z.Handlers{
	"team":   auction.TeamHandler{},
	"purse":  auction.PurseHandler{},
	"player": auction.PlayerHandler{},
}

var config = z.Config{
	ServerPort: 8500,
	Database:   z.MysqlConfig{os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), "hpl_auction"},
}

func main() {
	z.Service{Handlers: handlers, Config: &config}.Run()
}
