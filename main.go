package main

import (
	"net/http"
	"twitter-stat/views"
)

type templateType struct {
	title string
	time  string
}

func main() {

	http.HandleFunc("/login", views.LoginHandler)
	http.HandleFunc("/signup", views.SignUpHandler)
	http.ListenAndServe(":8000", nil)
}
