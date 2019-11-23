package main

import (
    "fmt"
	"net/http"
	"log"
)

func handlerFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1> Welcome to the Twitter Statics </h1>")
}

func main() {
	
	//Go server resides here 
	http.HandleFunc("/", handlerFunc)
	if err := http.ListenAndServe(":8080", nil); 
	
	err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}