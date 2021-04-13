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
}

type Book struct {
	BookSingle
}

var db *sql.DB
var err error

func init() {

}

func main0() {

	var n int
	fmt.Printf("Enter the book id: ")
	fmt.Scanf("%d", &n)

	str := "SELECT name  FROM `books`  where id = ? "
	stmt, err := db.Prepare(str)

	if err != nil {
		log.Println(err)
	}

	var name string
	err = db.QueryRow(str, n).Scan(&name)

	defer stmt.Close()

	fmt.Println(name)
}

func main1() {

	var n int
	fmt.Printf("Enter the book id: ")
	fmt.Scanf("%d", &n)

	str := "SELECT name  FROM `books`  where id = ? "
	stmt, err := db.Prepare(str)

	if err != nil {
		log.Println(err)
	}

	var name string
	err = stmt.QueryRow(n).Scan(&name)

	defer stmt.Close()

	fmt.Println(name)
}

func main() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3307)/test")

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

	str := "SELECT name  FROM `books`  where id = ? "
	stmt, err := db.Prepare(str)

	if err != nil {
		log.Println(err)
	}

	//err = stmt.QueryRow(n).Scan(&book.ID, &book.Name, &book.Author, &book.Pages, &book.PublicationDate)
	err = stmt.QueryRow(n).Scan(&book.Name)

	defer stmt.Close()

	fmt.Println(book)
}
