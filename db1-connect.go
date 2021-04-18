package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	//tmpDB, err := sql.Open("postgres", "user=postgres password=postgres dbname=dev sslmode=disable")
	tmpDB, err := sql.Open("mysql", "root:root@2021@tcp(localhost:3306)/test")
	logOnErr(err)
	db = tmpDB
	//	defer db.Close()
}

func logOnErr(err error) {
	if err != nil {
		//log.Println(err.Error())
		panic(err)
	}
}

func main() {

	err := db.Ping()
	logOnErr(err)

	fmt.Println("Db connected")

	defer db.Close()
}
