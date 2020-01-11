package main

import (
	// "io"
	// "net"
	"log"
	"fmt"
	"net/http"
)

// var templates *template.Template

func main() {
	// handle index page handler function
	http.HandleFunc("/", indexPageHandler)

	// Establishing a server and connection over http
	err := http.ListenAndServe(":8080", nil)
	if err!= nil{
		log.Panic(err)
	}


}

func indexPageHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "A template need to be placed here!")
}
