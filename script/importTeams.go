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
	f, err := os.Open("team.csv")
	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for ", err)
	}
	query := "INSERT INTO team(username,password,team_name,owners_name, icon1,icon2)VALUES(?,?,?,?,?,?)"
	for _, v := range records {
		_, err := dbo.Exec(query, v[0], v[1], v[2], v[3], v[4], v[5])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("done")
}
