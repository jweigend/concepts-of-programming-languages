// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0

// Package mail contains the Mail API interfaces and datatypes for sending Emails.
package mail

// Address is the address of the mail receiver.
type Address struct {
	Address string
}

// Sender is a interface to send mails.
type Sender interface {

	// Send an email to a given address with a  message.
	SendMail(address Address, message string)
}

// END OMIT
