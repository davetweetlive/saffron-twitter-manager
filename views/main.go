package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"
	"tweet_olivia/settings"

	"golang.org/x/crypto/bcrypt"
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
	defer db.Close()
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

func SignupHandler(res http.ResponseWriter, req *http.Request) {
	db, _ := settings.EstablishDBConnection()
	if req.Method != "POST" {
		http.ServeFile(res, req, "templates/signup.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	var user string

	err := db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)
	fmt.Println("Getting Error Information!")
	fmt.Println(err)
	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		}

		_, err = db.Exec("INSERT INTO users(username, password) VALUES(?, ?)", username, hashedPassword)
		if err != nil {
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		}

		res.Write([]byte("User created!"))
		return
	case err != nil:
		http.Error(res, "Server error, unable to create your account.", 500)
		return
	default:
		http.Redirect(res, req, "/", 301)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

}

func MainPageHandler(w http.ResponseWriter, r *http.Request) {

}
