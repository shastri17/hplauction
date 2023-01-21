package main

import (
	"encoding/csv"
	"fmt"
	"github.com/shastri17/hplauction/db"
	"log"
	"os"
)

func main() {
	dbo := db.GetMySQLObject()
	f, err := os.Open("hpl2023players.csv")
	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for ", err)
	}
	query := "INSERT INTO player (name, nick_name, skill_area, batting_hand, bowling_hand, mobile_number, whatsapp_number, previously_played_teams, image) VALUES (?,?,?,?,?,?,?,?,?)"
	for i, v := range records {
		if i == 0 {
			continue
		}

		_, err := dbo.Exec(query, v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("done")
}
