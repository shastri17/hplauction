package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func GetMysqlDb() (*gorm.DB, error) {
	mysqlConn := "root:Root@123@tcp(localhost:3306)/hpl_auction?charset=utf8&parseTime=True&loc=Local"
	log.Printf("Connecting MySQL: %v", mysqlConn)
	var err error
	DB, err := gorm.Open("mysql", mysqlConn)
	if err != nil {
		log.Printf("Error in MySQL Connection: %v", err.Error())
		return nil, err
	}
	// Make a test query
	var response struct {
		Version string
	}
	DB.Raw("select @@version as version").Scan(&response)
	log.Printf("MySQL Version: %v", response.Version)
	return DB, err
}
