package main

import (
    "fmt"
	"net/http"
	// "log"
)

func handlerFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1> Welcome to the twitter Statics </h1>")
}

func main() {
	
	//Go server resides here 
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8080", nil)
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}