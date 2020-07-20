package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DBStruct struct {
	DB MysqlCred `json:"mysqlCred"`
}

type MysqlCred struct {
	Database string `json:"database"`
	Host     string `json:"host"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Port     int64  `json:"port"`
	Username string `json:"username"`
}

func ConnectionStr() string {
	// Open the json credential file and defer close
	jsonFile, err := os.Open("config/dbcred.json")
	if err != nil {
		log.Fatal("ERROR: couldn't open json file", err)
	}
	defer jsonFile.Close()

	// Read the file and get the byte values
	jsonByteVal, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("ERROR: couldn't read json file", err)
	}

	// Declare DBStruct type and unmarshal
	var credential DBStruct
	json.Unmarshal(jsonByteVal, &credential)

	// name := "mysql"
	host := credential.DB.Host
	port := credential.DB.Port
	database := credential.DB.Database
	username := credential.DB.Username
	password := credential.DB.Password

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)
	// fmt.Println(connStr)
	return connStr
}

func ConDB() {
	db, err := sql.Open("mysql", ConnectionStr())

	if err != nil {
		fmt.Println("ERROR:")
	}

	stm, err := db.Prepare(`CREATE Table employee(id int NOT NULL AUTO_INCREMENT, first_name varchar(50), last_name varchar(30), PRIMARY KEY (id));`)
	if err != nil {
		fmt.Println("ERRPR!")
	}
	fmt.Println(stm)

}
