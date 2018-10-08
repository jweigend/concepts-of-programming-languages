// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0
package client

import (
	"testing"

	"github.com/qaware/programmieren-mit-go/01_object-oriented-programming/mail/smtp"
)

func TestMail(t *testing.T) {
	// Create an implementation for the mail.Sender interface.
	sender := new(smtp.MailSenderImpl)

	// We can use different mail implementations for this method.
	sendMail(sender)
}
