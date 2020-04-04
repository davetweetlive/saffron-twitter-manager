package main

import (
	"fmt"
	"net/http"
	"twitter-stat/models"
	"twitter-stat/views"
)

func main() {
	db, err := models.EstablishDBConnection()
	if err != nil {
		fmt.Println("DB connection denied!")
	}
	fmt.Println(db)
	http.HandleFunc("/login", views.LoginHandler)
	http.HandleFunc("/signup", views.SignUpHandler)
	http.ListenAndServe(":8000", nil)
}
