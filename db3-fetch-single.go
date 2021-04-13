package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type BookSingle struct {
	ID              int
	Name            string
	Author          string
	Pages           int
	PublicationDate string
	AddedOn         string
}

type Book struct {
	BookSingle
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
	var n int
	fmt.Printf("Enter the book id: ")
	fmt.Scanf("%d", &n)
	var book Book

	err = db.QueryRow("SELECT * FROM `books` where id = ? ", n).Scan(&book.ID, &book.Name, &book.Author, &book.Pages, &book.PublicationDate, &book.AddedOn)
	//err = db.QueryRow("SELECT * FROM `books` where id = ? ", n)
	//err.Scan(&book)

	if err != nil {
		log.Println(err)
	}

	fmt.Print(book.ID, "====> ")
	fmt.Println(book)

}
