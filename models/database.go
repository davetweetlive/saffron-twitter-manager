package models

import (
	"database/sql"
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

// var dbCred = DBCredential{
// 	DB:       "mysql",
// 	Database: "twitter-stat",
// 	Host:     "localhost",
// 	Port:     3306,
// 	Username: "root",
// 	Password: "Megamind@1",
// }

func EstablishDBConnection() (*sql.DB, error) {
	// Mysql satabase connection string formatted
	// dbConnString := fmt.Sprintf("%s:%s@/%s", dbCred.Username, dbCred.Password, dbCred.Database)

	//Connect to mySql database
	db, err := sql.Open("mysql", SqlConnectionString())
	if err != nil {
		log.Fatal("Couldn't connect to the database!")
	}

	// Return type *sql.DB and err (if any)
	return db, err
}
