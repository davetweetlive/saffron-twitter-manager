package settings

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
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
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("Successfully Opened dbcred.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our dbStruct struct
	var dbStruct DBStruct

	json.Unmarshal(byteValue, &dbStruct)

	dbConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbStruct.Db.Username, dbStruct.Db.Password, dbStruct.Db.Host, strconv.Itoa(dbStruct.Db.Port), dbStruct.Db.Database)

	return dbConnString
}

func EstablishDBConnection() (*sql.DB, error) {

	db, err := sql.Open("mysql", SqlConnectionString())

	if err != nil {
		log.Fatal("Couldn't connect to the database!")
	}

	return db, err
}
