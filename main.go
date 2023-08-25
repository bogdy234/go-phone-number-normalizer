package main

import (
	"database/sql"
	"fmt"
	"log"
	"phone-number-normalizer/initializers"
	"strings"

	"github.com/joho/godotenv"
)

type PhoneNumberEntry struct {
	id          int
	phoneNumber string
}

// getPhoneNumbers - get the phone numbers from the database and return it as an array
func getPhoneNumbers(db *sql.DB) []PhoneNumberEntry {
	rows, _ := db.Query("SELECT id, phone_number FROM phone_numbers")
	defer rows.Close()

	var phoneNumbers []PhoneNumberEntry
	for rows.Next() {
		var p PhoneNumberEntry
		err := rows.Scan(&p.id, &p.phoneNumber)
		if err != nil {
			panic(err)
		}

		phoneNumbers = append(phoneNumbers, PhoneNumberEntry{
			id:          p.id,
			phoneNumber: p.phoneNumber,
		})
	}

	return phoneNumbers
}

// formatPhoneNumber - take phone numbers, format them. If it's a duplicate remove it from the database, else update it
func formatPhoneNumber(p *PhoneNumberEntry, db *sql.DB) {
	replacer := strings.NewReplacer(" ", "", "(", "", ")", "", "-", "")
	p.phoneNumber = replacer.Replace(p.phoneNumber)

	query := fmt.Sprintf("SELECT id FROM phone_numbers WHERE phone_number='%s'", p.phoneNumber)
	fmt.Println(query)

	rows, _ := db.Query(query)
	defer rows.Close()

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			panic(err)
		}

		fmt.Println(id, p.id)
		if id != p.id {
			db.Query(fmt.Sprintf("DELETE FROM phone_numbers WHERE id=%d", p.id))
		}
	}

	db.Query(fmt.Sprintf("UPDATE phone_numbers SET phone_number='%s' WHERE id=%d", p.phoneNumber, p.id))
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	db, err := initializers.InitConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	phoneNumbers := getPhoneNumbers(db)
	fmt.Println(phoneNumbers)

	for i := range phoneNumbers {
		formatPhoneNumber(&phoneNumbers[i], db)
	}
	fmt.Println(phoneNumbers)
}
