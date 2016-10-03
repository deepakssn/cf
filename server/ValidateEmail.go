package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// ValidateUser will check if the domain received in the email is valid and generate OTP
func ValidateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var email = strings.ToLower(r.Form["email"][0])
	if validateEmail(email) {
		var emailSplit = strings.Split(email, "@")
		domainFlag := checkAllowedDomain(emailSplit[1])
		if domainFlag == 1 {
			otp := random(1000, 9999)
			if insertAuthToDB(email, otp) {
				if sendOTP(otp, "deepakssn.aws@gmail.com") {
					response, err := GenerateSuccessResponse("AUTH", "OTP Emailed Successfully")
					if err != nil {
						panic(err)
					}
					fmt.Fprintf(w, string(response))
				}
			}
		} else if domainFlag == 2 {

		}
	} else {
		response, err := (json.MarshalIndent(jsonError{Result: "FAIL", ErrorCategory: "EMAAIL", ErrorCode: "VE001", ErrorMessage: "Invalid Email Address"}, "", " "))
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, string(response))
	}
}
func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

//CheckAllowedDomain Checks if the Domain is in allowed list from DB
/*
+-------+----------------+
| VALUE | STATE          |
+-------+----------------+
| 0     | BLOCKED        |
+-------+----------------+
| 1     | ALLOWED        |
+-------+----------------+
| 2     | NOT CONFIGURED |
+-------+----------------+
*/
func checkAllowedDomain(d string) int {
	db, err := sql.Open("mysql", "sql:deepu@/devtest")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.Ping()
	stmtOut, err := db.Prepare("SELECT count(DOMAIN) FROM DOMAIN WHERE DOMAIN = ? and ALLOW = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()
	var count int
	err = stmtOut.QueryRow(d, 1).Scan(&count)
	if err != nil {
		panic(err)
	}
	if count > 0 {
		return 1 // Allowed
	}
	err = stmtOut.QueryRow(d, 0).Scan(&count)
	if err != nil {
		panic(err)
	}
	if count > 0 {
		return 0 // blocked
	}
	return 2 // Default
}

// Get the IP Address of the http Request
func getIP(r *http.Request) string {
	fmt.Printf("r: %+v\n", r)
	ip := r.Referer()
	fmt.Println(ip)
	return ip
}
