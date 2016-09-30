package main

import (
	"fmt"
	"net/http"
	"regexp"
)

// ValidateEmail will check if the domain received in the email is valid and generate OTP
func ValidateEmail(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var email = r.Form["email"]
	fmt.Println(validateEmail(email[0]))
}
func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}
