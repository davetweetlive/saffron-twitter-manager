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
