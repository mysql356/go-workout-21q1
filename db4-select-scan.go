package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Bk struct {
	ID      int
	Name    string
	AddedOn string
}

/*
-Struct (ID, Name, AddedOn)

#Mandatory
-Query (seelct id,name from books)
-Scan (&ID, &Name)

*/

func main() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:root@2021@tcp(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	//multiple row ///////////////////////////////////////////////////
	rows, err := db.Query("SELECT id, name FROM `books` ")
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var book Bk
		err = rows.Scan(&book.ID, &book.Name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(book)
	}

}
