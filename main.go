package main

import (
	"html/template"
	"log"
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
	http.HandleFunc("/signup", signUpHandler)
	http.ListenAndServe(":8000", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	utility.CheckIfTableExist()
	dt := time.Now()
	p := templateType{title: "Twitter Stat", time: dt.String()}
	temlt, _ := template.ParseFiles("template/index.html")
	temlt.Execute(w, p)
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	tmplt, err := template.ParseFiles("template/signup.html")
	if err != nil {
		log.Fatal(err)
	}
	tmplt.Execute(w, nil)
}
