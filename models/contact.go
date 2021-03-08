package models

import (
	"regexp"
	"strings"
	"time"

	"github.com/go-mail/mail"
)

type Talk struct {
	Name    string
	Email   string
	Phone   string
	Subject string
	Message string
	Errors  map[string]string
}

func (tlk *Talk) Deliver() error {
	email := mail.NewMessage()
	email.SetHeader("To", "deliri.ase@gmail.com")
	email.SetHeader("From", "deliri.ase@gmail.com")
	email.SetHeader("Reply-To", tlk.Email)
	email.SetHeader("Subject", "subject test")
	email.SetDateHeader("Date", time.Now())
	email.SetBody("text/plain", tlk.Message+" \nClient phone number: "+tlk.Phone)

	username := "afghancoders@gmail.com"
	password := "'y8MW1%r(72yFVIKHa60.ZKlGvA*1RR/~x33#bSVO[hrlhe!4d"

	return mail.NewDialer("smtp.gmail.com", 587, username, password).DialAndSend(email)
}

var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (tlk *Talk) Validate() bool {
	tlk.Errors = make(map[string]string)

	if strings.TrimSpace(tlk.Name) == "" {
		tlk.Errors["Name"] = "Please enter your name"
	}

	match := rxEmail.Match([]byte(tlk.Email))
	if match == false {
		tlk.Errors["Email"] = "Please enter a valid email address."
	}

	if len(strings.TrimSpace(tlk.Phone)) > 13 {
		tlk.Errors["Phone"] = "Please enter a valid phone number no longer than 12 digits"
	}

	if len(strings.TrimSpace(tlk.Subject)) > 25 {
		tlk.Errors["Subject"] = "Please enter a valid subject."
	}

	if len(strings.TrimSpace(tlk.Subject)) == 0 {
		tlk.Errors["Subject"] = "The Subject cannot be blank."
	}

	if len(strings.TrimSpace(tlk.Message)) > 1000 {
		tlk.Errors["Subject"] = "Please enter a valid Message no longer than 1000 characters."
	}

	if len(strings.TrimSpace(tlk.Message)) == 0 {
		tlk.Errors["Subject"] = "Please enter a message."
	}

	return len(tlk.Errors) == 0
}
