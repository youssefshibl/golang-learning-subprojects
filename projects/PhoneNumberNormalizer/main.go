package main

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"

	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var phone_numbers = []string{
	"1234567890",
	"123 456 7891",
	"(123) 456 7892",
	"(123) 456-7893",
	"123-456-7894",
	"123-456-7890",
	"1234567892",
	"(123)456-7892",
}

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	checkErr(err)
	defer db.Close()
	// create table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS phone_numbers (number CHAR(255))")
	checkErr(err)
	// truncate table
	_, err = db.Exec("DELETE FROM phone_numbers")
	checkErr(err)
	// insert data
	_, err = db.Exec("INSERT INTO phone_numbers (number) VALUES ('123-456-7890'), ('987-654-3210')")
	checkErr(err)

	insertStmt, err := db.Prepare("INSERT INTO phone_numbers (number) VALUES (?)")
	checkErr(err)
	for _, number := range phone_numbers {
		_, err = insertStmt.Exec(number)
		checkErr(err)
	}
	printAllPhoneNumbers(db)
	fmt.Println("---------------------")
	updateAllPhoneNumbers(db)
	fmt.Println("---------------------")
	printAllPhoneNumbers(db)

}

func printAllPhoneNumbers(db *sql.DB) {
	rows, err := db.Query("select * from phone_numbers")
	checkErr(err)
	for rows.Next() {
		var phone_number string
		rows.Scan(&phone_number)
		fmt.Println(phone_number)
	}
	rows.Close()
}

func updateAllPhoneNumbers(db *sql.DB) {
	rows, err := db.Query("select * from phone_numbers")
	checkErr(err)

	var phoneNumbers []string
	for rows.Next() {
		var phone_number string
		rows.Scan(&phone_number)
		phoneNumbers = append(phoneNumbers, phone_number)
	}
	rows.Close()

	updateStmt, err := db.Prepare("UPDATE phone_numbers SET number = ? WHERE number = ?")
	checkErr(err)
	checkExist, err := db.Prepare(
		`SELECT 1 FROM phone_numbers WHERE number = ? LIMIT 1`)
	checkErr(err)
	deletePhone, err := db.Prepare(
		`DELETE FROM phone_numbers WHERE number = ?`)
	checkErr(err)
	defer updateStmt.Close()
	defer checkExist.Close()
	defer deletePhone.Close()

	for _, phone_number := range phoneNumbers {
		new_phone_number := remainDigitOnly(phone_number)
		res, err := checkExist.Query(new_phone_number)
		checkErr(err)
		duplicate := res.Next()
		res.Close()
		if duplicate {
			fmt.Printf("Phone number %s already exists, deleting it.\n", new_phone_number)
			_, err = deletePhone.Exec(phone_number)
			checkErr(err)
			continue
		}
		_, err = updateStmt.Exec(new_phone_number, phone_number)
		checkErr(err)
	}
}

func remainDigitOnly(phone_number string) string {
	r, _ := regexp.Compile(`[^\d]`)
	return r.ReplaceAllString(phone_number, "")

}
