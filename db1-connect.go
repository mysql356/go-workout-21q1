package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:root@2021@tcp(127.0.0.1:3306)/test")

	err = db.Ping()

	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Println("Db connected")

	defer db.Close()
}
