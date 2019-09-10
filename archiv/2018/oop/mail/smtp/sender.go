// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0

// Package smtp sends mails over the smtp protocol.
package smtp

import (
	"log"

	"github.com/jweigend/concepts-of-programming-languages/oop/mail"
)

// MailSenderImpl is a sender object.
type MailSenderImpl struct {
}

// SendMail sends a mail to a receiver.
func (m *MailSenderImpl) SendMail(address mail.Address, message string) {
	log.Println("Sending message with SMTP to " + address.Address + " message: " + message)
	return
}

//END OMIT
