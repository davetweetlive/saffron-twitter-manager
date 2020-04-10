package main

import (
	"fmt"
	"net/http"
	"text/template"
	"tweet_olivia/settings"
)

type LoggedInUser struct {
	username      string
	authenticated bool
	loginTime     string
}

var tmplt *template.Template

func init() {
	tmplt = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	// Some random function test
	db, err := settings.EstablishDBConnection()
	if err != nil {
		panic(err)
	}
	fmt.Println("DB is connected!", db)
	// Static file server
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Default Mux
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/signup", SignupHandler)
	http.ListenAndServe(":8080", nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmplt.ExecuteTemplate(w, "index.html", nil)
	} else if r.Method == "POST" {
		fmt.Fprintf(w, "Login function needs to be written")
	}

}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmplt.ExecuteTemplate(w, "signup.html", nil)
	} else if r.Method == "POST" {
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		// fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		username := r.FormValue("username")
		emailID := r.FormValue("email_id")
		password := r.FormValue("password")

		fmt.Printf("%T\n", username)
		fmt.Printf("%T\n", emailID)
		fmt.Printf("%T\n", password)

	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

}

func MainPageHandler(w http.ResponseWriter, r *http.Request) {

}
