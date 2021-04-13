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
	Author          sql.NullString
	Pages           int
	PublicationDate sql.NullString
	AddedOn         string
}

type Book struct {
	BookSingle
}

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("mysql", "root:root@2021@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Print(err.Error())
	}

}
func main() {

	defer db.Close()
	//multiple row ///////////////////////////////////////////////////
	var n int
	fmt.Printf("Enter the book id: ")
	fmt.Scanf("%d", &n)
	var book Book

	//case-1
	//err = db.QueryRow("SELECT * FROM `books` where id = ? ", n).Scan(&book.ID, &book.Name, &book.Author, &book.Pages, &book.PublicationDate, &book.AddedOn)

	//case-2 [ select * == Scan(...), here 6 = 6 ]
	row := db.QueryRow("SELECT * FROM `books` where id = ? ", n)                                          //6
	err = row.Scan(&book.ID, &book.Name, &book.Author, &book.Pages, &book.PublicationDate, &book.AddedOn) //6

	if err != nil {
		log.Println(err)
	}

	fmt.Print(book.ID, "====> ")
	fmt.Println(book)
	fmt.Println(book.Author.Value())
	fmt.Println(book.Author.String, book.Author.Valid)

	Query()

}

func Query() {

	defer db.Close()
	//multiple row ///////////////////////////////////////////////////
	var version string
	var ctime string
	err = db.QueryRow("SELECT VERSION(), now()").Scan(&version, &ctime)
	fmt.Println(version, ctime)

}
