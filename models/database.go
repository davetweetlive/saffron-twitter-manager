package models

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

func EstablishDBConnection() (*sql.DB, error) {

	//Connect to mySql database SqlConnectionString() func sends connection string
	db, err := sql.Open("mysql", SqlConnectionString())
	if err != nil {
		log.Fatal("Couldn't connect to the database!")
	}
	fmt.Println("DB has been setup!")
	// Return type *sql.DB and err (if any)
	return db, err
}
