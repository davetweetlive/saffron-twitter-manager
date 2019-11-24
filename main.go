package main

import (
	// "fmt"
	"github.com/gorilla/mux"
	"net/http"
	"html/template"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}


func indexHandler(w http.ResponseWriter, r *http.Request)  {
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
