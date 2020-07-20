package main

import (
	"log"
	"net/http"
	"twitter-stat/db"
)

func main() {
	// To check for db connection establishment
	db.ConnectionStr()
	db.ConDB()

	// Server setup
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error: There is trouble running the server!", err)
	}
}
