package main

import (
	"html/template"
	"net/http"
	"time"
	"twitter-stat/utility"
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
	utility.CheckIfTableExist()
	dt := time.Now()
	p := templateType{title: "Twitter Stat", time: dt.String()}
	t, _ := template.ParseFiles("template/index.html")
	t.Execute(w, p)
}
