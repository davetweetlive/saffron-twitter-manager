package views

import (
	"fmt"
	"net/http"
	"text/template"
	"twitter-stat/models"

	"golang.org/x/crypto/bcrypt"
)

type LoginInfo struct {
	username      string
	authenticated bool
}

type TemplateInformation struct {
	Title     string
	PoweredBy string
}

var tmplt *template.Template

func init() {
	tmplt = template.Must(template.ParseGlob("templates/*.html"))
}

// type loggedin
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tempInfo := TemplateInformation{
		Title:     "Login to Twitter Handler",
		PoweredBy: "Saffron Coders",
	}
	if r.Method == "GET" {
		// http.ServeFile(w, r, "templates/index.html")

		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			fmt.Println(err)
		}

		t.Execute(w, tempInfo)
		// err := tmplt.Execute(w)
		// if err != nil {
		// 	panic(err)
		// }
	} else if r.Method == "POST" {
		fmt.Println("Write the login function!")
	} else {
		fmt.Println("We couldn't find any get or a post method!")
	}
	// 	// utility.CheckIfTableExist()
	// 	temlt, _ := template.ParseFiles("template/index.html")
	// 	temlt.Execute(w, nil)
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "templates/signup.html")
	case "POST":
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

		encpass, err := bcrypt.GenerateFromPassword([]byte(password), 5)
		if err != nil {
			fmt.Println("Couldn't encrypt")
		}
		fmt.Printf("%T\n", string(encpass[:]))
		db, err := models.EstablishDBConnection()
		defer db.Close()

		if err != nil {
			fmt.Println(err)
		}

		// _, err = db.Query(`INSERT INTO TABLE VALUES(`  `,` + emailID + `, ` + string(encpass[:]) + `)`)
		_, err = db.Query(`CREATE TABLE "User" (
			"id" INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
			"username" VARCHAR(50) NOT NULL,
			"email" VARCHAR(50) NOT NULL,
			"passwd" VARCHAR(50),
			reg_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
			)`)
		if err != nil {
			fmt.Println(err)
		}
		_, err = db.Query(`INSERT INTO MyGuests (username, email, passwd) VALUES (` + username + `, ` + emailID + `, ` + string(encpass[:]) + `)`)
		if err != nil {
			fmt.Println(err)
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "templates/home.html")
	} else if r.Method == "POST" {
		fmt.Println("Write the login function!")
	} else {
		fmt.Println("We couldn't find any get or a post method!")
	}
	// 	// utility.CheckIfTableExist()
	// 	temlt, _ := template.ParseFiles("template/index.html")
	// 	temlt.Execute(w, nil)
}

func TweetStreaming(w http.ResponseWriter, r *http.Request) {
	info := LoginInfo{
		username:      "na",
		authenticated: false,
	}
	err := tmplt.ExecuteTemplate(w, "home.html", info)
	if err != nil {
		fmt.Println(err)
	}
}
