package main

import (
	"log"
	"net/smtp"
	"strconv"
)

// SendOTP will send an OTP
func sendOTP(otp int, to string) bool {
	from := ""
	pass := ""

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: OTP using GO: " + strconv.Itoa(otp) + "\n\n" + strconv.Itoa(otp)

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return false
	}

	log.Print("Email Sent!")
	log.Print("OTP : ", otp)
	return true

}
