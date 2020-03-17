package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	"twitter-stat/database"
)

type templateType struct {
	title string
	time  string
}

func main() {

	http.HandleFunc("/login", loginHandler)
	http.ListenAndServe(":8000", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := database.MysqlConnect()
	if err != nil {
		log.Fatal("Couldn't connect to the database!")
	}
	fmt.Println("The database connection has been established!", conn)

	dt := time.Now()
	p := templateType{title: "Twitter Stat", time: dt.String()}
	t, _ := template.ParseFiles("template/index.html")
	t.Execute(w, p)
}
