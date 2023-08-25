package initializers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	dbname = "go-phone-number-normalizer"
)

// InitConnection - initiate connection to the database and return it
func InitConnection() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("PASSWORD")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db, err
}
