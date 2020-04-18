package settings

import (
	"fmt"
	"twitter-stat/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	id       int
	username string
	email    string
	password string
}

func MYSQLConn() {
	db, err := gorm.Open("mysql", models.SqlConnectionString()+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// db.CreateTable(&User{})
	fmt.Println("The ORM db connection is working")
}
