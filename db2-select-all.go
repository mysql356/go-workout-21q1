package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*

CREATE TABLE `books` (
  `id` int NOT NULL,
  `name` varchar(50) DEFAULT NULL,
  `author` varchar(50) DEFAULT NULL,
  `pages` int DEFAULT NULL,
  `publication_date` date DEFAULT NULL,
  `addedOn` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=MyISAM DEFAULT CHARSET=latin1;


INSERT INTO `books` (`id`, `name`, `author`, `pages`, `publication_date`, `addedOn`) VALUES
(1, 'Go', 'A Singh', 400, '2020-12-02', '2021-04-03 04:10:18'),
(2, 'Java', 'B K SINGH', 230, '2021-03-24', '2021-04-03 04:10:18'),
(3, 'C', 'C K sing', 22, '2021-03-24', '2021-04-03 04:10:18'),
(4, 'C++', 'D k sing', 33, '2021-03-24', '2021-04-03 04:10:18');

ALTER TABLE `books` ADD PRIMARY KEY (`id`);
ALTER TABLE `books` MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
*/

type BookAll struct {
	ID              int
	Name            string
	Author          sql.NullString
	Pages           int
	PublicationDate sql.NullString
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
			log.Println(err)
		}

		fmt.Print(book.ID, "====> ")
		fmt.Println(book)

	}

}
