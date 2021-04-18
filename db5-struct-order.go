package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type BookOrder struct {
	A int
	B string
	C sql.NullString
	D int
	E sql.NullString
	F string
}

type Book struct {
	BookOrder
}

var db *sql.DB
var err error

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

	rows, err := db.Query("SELECT * FROM `books` ")
	logOnErr(err)

	for rows.Next() {
		var book Book
		err = rows.Scan(&book.A, &book.B, &book.C, &book.D, &book.E, &book.F)
		logOnErr(err)

		fmt.Print(book.A, "====> ")
		fmt.Println(book)

	}

}
