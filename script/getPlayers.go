package main

import (
	"encoding/csv"
	"fmt"
	"github.com/shastri17/hplauction/db"
	"github.com/shastri17/hplauction/models"
	"os"
)

func main() {
	dbo, err := db.GetMysqlDb()
	f, err := os.Create("hpl2023TeamWisePlayers.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	w := csv.NewWriter(f)
	var teams []models.Team
	dbo.Table("team").Where("is_admin=?", 0).Find(&teams)
	var players []models.Player
	for i := range teams {
		w.Write([]string{teams[i].TeamName})
		dbo.Table("player").Where("team_id=?", teams[i].Id).Find(&players)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, p := range players {
			w.Write([]string{p.Name, fmt.Sprint(p.WhatsappNumber)})
		}
		w.Write([]string{""})
	}
	w.Flush()

}
