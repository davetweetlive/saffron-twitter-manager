package utility

import (
	"fmt"
	"log"
	"twitter-stat/models"
)

func CheckIfTableExist() {
	// Establish the database connection
	db, err := models.EstablishDBConnection()
	defer db.Close()
	if err != nil {
		log.Fatal("Couldn't connect to the database!")
	}
	fmt.Println("The database connection has been established!", db)

	// To check if the table exists or not if the table exists don't create a new table
	tableExists, err := db.Query("SELECT 1 FROM `user` LIMIT 1;")
	if err != nil {
		// log.Fatal("Couldn't prcess the query to check for the table existance!")
		selDB, err := db.Query("CREATE TABLE `user` (`id` int(6) unsigned NOT NULL AUTO_INCREMENT, `user` varchar(30) NOT NULL, `password` varchar(30) NOT NULL, `first_name` varchar(50), `last_name` varchar(50), `email` varchar(50), PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;")

		if err != nil {
			log.Fatal("Couldn't create a table!")
		}
		fmt.Println(selDB)
	} else {
		fmt.Println("The table exists and it's being used for info access!")
	}
	if tableExists != nil {
		fmt.Println("The table exists!")
	}

	// insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")

	// // if there is an error inserting, handle it
	// if err != nil {
	// 	panic(err.Error())
	// }
	// // be careful deferring Queries if you are using transactions
	// defer insert.Close()
}

func CreateUsertable() {
	db, err := models.EstablishDBConnection()
	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}

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
}
