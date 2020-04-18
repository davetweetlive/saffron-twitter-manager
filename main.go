package main

import (
	"net/http"
	"twitter-stat/settings"
	"twitter-stat/views"
)

func main() {
	settings.MYSQLConn()

	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	// db, err := models.EstablishDBConnection()
	// if err != nil {
	// 	fmt.Println("DB connection denied!")
	// }
	// fmt.Println(db)
	http.HandleFunc("/login", views.LoginHandler)
	http.HandleFunc("/signup", views.SignUpHandler)
	http.HandleFunc("/", views.IndexHandler)
	http.ListenAndServe(":8000", nil)
}
