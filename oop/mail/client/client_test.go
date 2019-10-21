// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0
package client

import (
	"testing"

	"github.com/jweigend/concepts-of-programming-languages/oop/mail/smtp"
)

// configure Registry to know which mail implementation is used.
func init() {
	Registry.Register("mail.Sender", new(smtp.MailSenderImpl))
}

func TestMail(t *testing.T) {
	SendMail("johannes.weigend@qaware.de", "Hello from Go!")
}
