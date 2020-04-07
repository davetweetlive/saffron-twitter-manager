package views

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "templates/index.html")
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

		encpass, err := bcrypt.GenerateFromPassword([]byte(password), 5)
		if err != nil {
			fmt.Println("Couldn't encrypt")
		}
		fmt.Println(username)
		fmt.Println(emailID)
		fmt.Println(string(encpass))
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


