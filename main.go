package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/bye", byeHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}


func helloHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Hello World!")
}
func byeHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Good Bye!")
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
