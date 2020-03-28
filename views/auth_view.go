package views

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// utility.CheckIfTableExist()
	temlt, _ := template.ParseFiles("template/index.html")
	temlt.Execute(w, nil)
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Printf(r.Method)
		tmplt, err := template.ParseFiles("template/signup.html")
		if err != nil {
			log.Fatal(err)
		}
		tmplt.Execute(w, nil)

	} else {
		uname := r.PostFormValue("username")
		// logic part of log in
		fmt.Println(uname)
		fmt.Println("It's a Post method!")
		fmt.Println(r.Method)
	}

}
