// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package proxy contains the client side proxy.
package proxy

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jweigend/concepts-of-programming-languages/dp/idserv/remote/idserv"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "client"
)

// Proxy is a client side proxy which encapsulates the RPC logic. It implements the IDService interface.
type Proxy struct {
	connection *grpc.ClientConn
}

// NewProxy creates a Proxy and starts the server connection
func NewProxy() *Proxy {
	p := new(Proxy)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("did not connect: %v", err))
	}
	p.connection = conn
	return p
}

// NewUUID implements the IDService interface.
func (p *Proxy) NewUUID(clientID string) string {
	c := idserv.NewIDServiceClient(p.connection)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.NewUUID(ctx, &idserv.IdRequest{ClientId: clientID})
	if err != nil {
		log.Printf("could not generate id: %v", err)
		r.Uuid = ""
	}
	return r.Uuid
}
