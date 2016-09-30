package main

import (
	"log"
	"net/smtp"
)

// SendOTP will send an OTP
func SendOTP(OTP string, to string) {
	from := "deepakssn.aws@gmail.com"
	pass := "******"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: OTP using GO: " + OTP + "\n\n" + "Mike, Please acknowledge the email in slack/whatsapp :)"

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("Email Sent!")
}
