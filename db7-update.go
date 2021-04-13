package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

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

	sql := fmt.Sprintf("Update books set author ='aa', publication_date = '2021-03-24'  WHERE id  = '%d' ", n)

	res, err := db.Exec(sql)
	if err != nil {
		panic(err.Error())
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The statement affected %d rows\n", affectedRows)
}
