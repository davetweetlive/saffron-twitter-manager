package main

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"html/template"
)

var templates *template.Template

func main() {
	connectToTheDatabase()

	templates = template.Must(template.ParseGlob("templates/*.html"))

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}


// Mysql database connection db open db close, earror handle and success notification on terminal
func connectToTheDatabase()  {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/database")

	if err != nil{
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO Persons (PersonID, FirstName, LastName) VALUES(2, 'Pappu', 'Kumar')")

	if err != nil{
		panic(err.Error())
		fmt.Println("There is a problem in Insert Query!")
	}
	defer insert.Close()

	fmt.Println("Successful")

}


// Index Page which handles all the method and route performed on the index page
func indexHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Welcome to index page")
	templates.ExecuteTemplate(w, "index.html", nil)
}








// func handlerFunc(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprint(w, "<h1> The Application is under construction! </h1>")
// 	// if r.URL.Path == "/" {
// 	//
// 	// } else {
// 	// 	w.WriteHeader(http.StatusNotFound)
// 	// }
// }
