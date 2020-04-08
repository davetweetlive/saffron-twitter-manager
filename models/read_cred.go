package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type DBStruct struct {
	Db Mysqldb `json:"mysqldb"`
}

type Mysqldb struct {
	Name      string `json:"name"`
	Connector string `json:"connector"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Database  string `json:"database"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func SqlConnectionString() string {

	jsonFile, err := os.Open("config/dbcred.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened dbcred.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our dbStruct struct
	var dbStruct DBStruct

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'dbStruct' which we defined above
	json.Unmarshal(byteValue, &dbStruct)

	dbConnString := fmt.Sprintf("%s:%s@/%s", dbStruct.Db.Username, dbStruct.Db.Password, dbStruct.Db.Database)

	return dbConnString
}

func CreateUserTable() {
	query := `CREATE TABLE "User" (
		"id" INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		"username" VARCHAR(50) NOT NULL,
		"email" VARCHAR(50) NOT NULL,
		"passwd" VARCHAR(50),
		reg_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		)`

	// selDB, err := db.Query("CREATE TABLE `user` (`id` int(6) unsigned NOT NULL AUTO_INCREMENT, `user` varchar(30) NOT NULL, `password` varchar(30) NOT NULL, `first_name` varchar(50), `last_name` varchar(50), `email` varchar(50), PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;")
	fmt.Println(query)
}
