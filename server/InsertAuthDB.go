package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func insertAuthToDB(email string, otp int) bool {
	db, err := sql.Open("mysql", "sql:deepu@/devtest")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	db.Ping()
	// Check if email address exist
	stmtOut, err := db.Prepare("SELECT count(*) FROM AUTH WHERE EMAIL = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()
	var count int
	err = stmtOut.QueryRow(email).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			count = 0
		} else {
			panic(err) // proper error handling instead of panic in your app
			// return false
		}
	}
	if count < 1 {
		// Prepare statement for inserting data
		stmtIns, err := db.Prepare("INSERT INTO AUTH (email, OTP, ATTEMPTS, ACTIVE, EXPIRY) VALUES( ?, ?, ?, ?, ? )") // ? = placeholder
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
			//return false
		}
		defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
		//Time + 24 Hours in UTC
		otpExpiry := time.Now().UTC().Add(24 * time.Hour)
		_, err = stmtIns.Exec(email, otp, 0, 1, otpExpiry) // Insert tuples (i, i^2)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	} else {
		stmtUpd, err := db.Prepare("UPDATE AUTH SET OTP = ?, ATTEMPTS = ?, ACTIVE = ?, EXPIRY = ? WHERE EMAIL = ?") // ? = placeholder
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
			//return false
		}
		defer stmtUpd.Close() // Close the statement when we leave main() / the program terminates
		//Time + 24 Hours in UTC
		otpExpiry := time.Now().UTC().Add(24 * time.Hour)
		_, err = stmtUpd.Exec(otp, 0, 1, otpExpiry, email) // Insert tuples (i, i^2)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}
	return true
}
