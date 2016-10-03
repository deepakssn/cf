package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func insertAuthToDB(email string, OTP int) bool {
	db, err := sql.Open("mysql", "sql:deepu@/devtest")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	db.Ping()

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO AUTH (email, OTP, ATTEMPTS, ACTIVE, EXPIRY) VALUES( ?, ?, ?, ?, ? )") // ? = placeholder
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		return false
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
	//Time + 24 Hours in UTC
	otpExpiry := time.Now().UTC().Add(24 * time.Hour)
	_, err = stmtIns.Exec(email, OTP, 0, 1, otpExpiry) // Insert tuples (i, i^2)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return true
}
