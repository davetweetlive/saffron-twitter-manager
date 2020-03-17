package main

import (
	"html/template"
	"net/http"
	"time"
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
	dt := time.Now()
	p := templateType{title: "Twitter Stat", time: dt.String()}
	t, _ := template.ParseFiles("template/index.html")
	t.Execute(w, p)
}
