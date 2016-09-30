package main

import (
	"fmt"
	"log"
	"net/smtp"
)

// SendOTP will send an OTP
func sendOTP(OTP string, to string) {
	from := "deepakssn.aws@gmail.com"
	pass := "AWSgeek1$"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: OTP using GO: " + OTP + "\n\n" + OTP

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("Email Sent!")
	fmt.Println("OTP : ", OTP)

}
