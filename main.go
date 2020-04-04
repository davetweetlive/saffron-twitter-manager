package main

import (
	"fmt"
	"net/http"
	"twitter-stat/models"
	"twitter-stat/views"
)

type templateType struct {
	title string
	time  string
}

func main() {
	fmt.Println(models.SqlConnectionString())

	http.HandleFunc("/login", views.LoginHandler)
	http.HandleFunc("/signup", views.SignUpHandler)
	http.ListenAndServe(":8000", nil)
}
