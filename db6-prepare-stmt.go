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

var db *sql.DB

func init() {
	//tmpDB, err := sql.Open("postgres", "user=postgres password=postgres dbname=dev sslmode=disable")
	tmpDB, err := sql.Open("mysql", "root:root@2021@tcp(localhost:3306)/test")
	if err != nil {
		log.Println(err)
		log.Println(err.Error())
	}
	db = tmpDB
	//	defer db.Close()
}

//simple
func main1() {

	var n int
	fmt.Printf("Enter the book id: ")
	fmt.Scanf("%d", &n)

	str := "SELECT name  FROM `books`  where id = ? "

	var name string
	err := db.QueryRow(str, n).Scan(&name)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(name)
	defer db.Close()
}

//prepare + single value
func main2() {

	var n int
	fmt.Printf("Enter the book id: ")
	fmt.Scanf("%d", &n)

	str := "SELECT name  FROM `books`  where id = ? "
	stmt, err := db.Prepare(str)

	if err != nil {
		log.Println(err)
	}

	var name string
	fmt.Println(name)
	err = stmt.QueryRow(n).Scan(&name)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(name)

	defer stmt.Close()
	defer db.Close()

}

//prepare + struct
func main() {

	//multiple row ///////////////////////////////////////////////////
	var n int
	fmt.Printf("Enter the book id: ")
	fmt.Scanf("%d", &n)

	var book Book
	fmt.Println(book)

	str := "SELECT *  FROM `books`  where id = ? "
	stmt, err := db.Prepare(str)

	if err != nil {
		log.Println(err)
	}

	err = stmt.QueryRow(n).Scan(&book.ID, &book.Name, &book.Author, &book.Pages, &book.PublicationDate, &book.AddedOn)
	//err = stmt.QueryRow(n).Scan(&book.Name)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(book)

	defer stmt.Close()
	defer db.Close()

}
