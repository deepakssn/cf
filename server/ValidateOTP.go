package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ValidateOTP will check if the passed email and OTP values are matching in DB
func ValidateOTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var email = strings.ToLower(r.Form["email"][0])
	otp, err := strconv.Atoi(strings.ToLower(r.Form["OTP"][0]))
	if err != nil {
		fmt.Println(err.Error())
	}
	dbOTP, dbExpiry := getOTPFromDB(email)
	now := time.Now().UTC()
	dbExpirytimeStamp, err := time.Parse("2006-01-02 15:04:05.999", dbExpiry)
	if err != nil {
		// do something with err...
	}
	if dbOTP == otp {
		if dbExpirytimeStamp.After(now) {
			fmt.Println("OTP VALID and MATCHES")
		} else {
			fmt.Println("OTP EXPIRED, but MATCHES")
		}
	} else {
		fmt.Println("OTP DOESN'T MATCH")
	}

}

// getOTPFromDB will return OTP
func getOTPFromDB(email string) (int, string) {
	db, err := sql.Open("mysql", "sql:deepu@/devtest")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.Ping()
	stmtOut, err := db.Prepare("SELECT OTP, EXPIRY FROM AUTH WHERE EMAIL = ? and ATTEMPTS < ? AND ACTIVE = ?")
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, ""
		}
		panic(err.Error())

	}
	defer stmtOut.Close()

	var dbOTP int
	var dbExpiry string
	err = stmtOut.QueryRow(email, 3, 1).Scan(&dbOTP, &dbExpiry)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, ""
		}
		panic(err)
	}
	return dbOTP, dbExpiry
}
