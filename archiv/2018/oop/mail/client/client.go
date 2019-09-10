// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0

// Package client contains sample code for the mail components.
package client

import (
	"github.com/jweigend/concepts-of-programming-languages/oop/mail"
	"github.com/jweigend/concepts-of-programming-languages/oop/mail/util"
)

// Registry is the central configration for the service locator
var Registry = util.NewRegistry()

// SendMail sends a mail to a receiver.
func SendMail(address, message string) {

	// Create an implementation for the mail.Sender interface.
	var sender = Registry.Get("mail.Sender").(mail.Sender)

	mailaddrs := mail.Address{Address: address}
	sender.SendMail(mailaddrs, message)
}

// EOF OMIT
