package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:Password@/database_name")

	if err != nil {
		log.Fatal("Couldn't connect to the database!")
	}
	fmt.Printf("%T\n", db)
	return db, err
}
