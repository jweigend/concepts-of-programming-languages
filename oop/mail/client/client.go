// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0

// Package client contains sample code for the mail components.
package client

import "github.com/jweigend/concepts-of-programming-languages/oop/mail"

func sendMail(s mail.Sender) {

	address := mail.Address{Address: "johannes.weigend@qaware.de"}
	message := "EMail from Go!"

	s.SendMail(address, message)
}

// EOF OMIT
