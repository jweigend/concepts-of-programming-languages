// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0

// Package client contains sample code for the mail components.
package client

import (
	"github.com/jweigend/concepts-of-programming-languages/oop/mail"
	"github.com/jweigend/concepts-of-programming-languages/oop/mail/util"
)

// Registry is the central configration for the service locator
var Registry = new(util.Registry)

// SendMail sends a mail to a receiver.
func SendMail(address, message string) {

	// Create an implementation for the mail.Sender interface.
	var sender mail.Sender
	Registry.Get(&sender)

	mailaddrs := mail.Address{Address: address}
	sender.SendMail(mailaddrs, message)
}

// EOF OMIT
