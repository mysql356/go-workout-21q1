package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type BookAll struct {
	ID              int
	Name            string
	Author          string
	Pages           int
	PublicationDate string
	AddedOn         string
}

type Book struct {
	BookAll
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
		err = rows.Scan(&book.ID, &book.Name, &book.Author, &book.Pages, &book.PublicationDate, &book.AddedOn)
		if err != nil {
			panic(err.Error())
		}

		fmt.Print(book.ID, "====> ")
		fmt.Println(book)

	}

}
