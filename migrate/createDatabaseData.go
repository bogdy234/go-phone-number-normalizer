package main

import (
	"database/sql"
	"fmt"
	"phone-number-normalizer/initializers"
)

func createTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS phone_numbers(
		id SERIAL,
		phone_number varchar(100)
	 );`)

	if err != nil {
		fmt.Println(err)
	}
}

func populateTable(db *sql.DB) {
	_, err := db.Exec(`INSERT INTO phone_numbers (phone_number)
	VALUES('1234567890'),
		('123 456 7891'),
		('(123) 456 7892'),
		('(123) 456-7893'),
		('123-456-7894'),
		('123-456-7890'),
		('1234567892'),
		('(123)456-7892')`)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	db, err := initializers.InitConnection()
	if err != nil {
		panic(err)
	}

	// createTable(db)
	populateTable(db)
	defer db.Close()
}
