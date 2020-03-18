package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	"twitter-stat/database"
)

type templateType struct {
	title string
	time  string
}

func main() {

	http.HandleFunc("/login", loginHandler)
	http.ListenAndServe(":8000", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	checkIfTableExist()
	dt := time.Now()
	p := templateType{title: "Twitter Stat", time: dt.String()}
	t, _ := template.ParseFiles("template/index.html")
	t.Execute(w, p)
}

func checkIfTableExist() {
	// Establish the database connection
	conn, err := database.MysqlConnect()
	if err != nil {
		log.Fatal("Couldn't connect to the database!")
	}
	fmt.Println("The database connection has been established!", conn)

	// To check if the table exists or not if the table exists don't create a new table
	tableExists, err := conn.Query("show tables like `account`;")
	if err != nil {
		log.Fatal("Couldn't prcess the query to check for the table existance!")
	}
	if tableExists != nil {
		selDB, err := conn.Query("CREATE TABLE `account` (`id` int(6) unsigned NOT NULL AUTO_INCREMENT, `username` varchar(30) NOT NULL, `password` varchar(30) NOT NULL, PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;")

		if err != nil {
			log.Fatal("Couldn't create a table!")
		}
		fmt.Println(selDB)
	} else {
		fmt.Println("The table exists and it's being used for info access!")
	}

}
