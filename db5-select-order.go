package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type BookOrder struct {
	A int
	B string
	C string
	D int
	E string
	F string
}

type Book struct {
	BookOrder
}

func main() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:root@2021@tcp(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	//multiple row ///////////////////////////////////////////////////
	rows, err := db.Query("SELECT * FROM `books` ")
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var book Book
		err = rows.Scan(&book.A, &book.B, &book.C, &book.D, &book.E, &book.F)
		if err != nil {
			panic(err.Error())
		}

		fmt.Print(book.A, "====> ")
		fmt.Println(book)

	}

}
