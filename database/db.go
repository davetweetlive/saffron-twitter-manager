package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DBCredential struct {
	DB       string
	Database string
	Host     string
	Port     int
	Username string
	Password string
}

var dbCred = DBCredential{
	DB:       "mysql",
	Database: "twitter-stat",
	Host:     "localhost",
	Port:     3606,
	Username: "root",
	Password: "Megamind@1",
}

func MysqlConnect() (*sql.DB, error) {
	dbConnString := fmt.Sprintf("%s:%s@/%s", dbCred.Username, dbCred.Password, dbCred.Database)
	fmt.Println(dbConnString)

	db, err := sql.Open("mysql", dbConnString)
	//  "root:Megamind@1@/twitter-stat"
	if err != nil {
		log.Fatal("Couldn't connect to the database!")
	}
	fmt.Printf("%T\n", db)
	return db, err
}
