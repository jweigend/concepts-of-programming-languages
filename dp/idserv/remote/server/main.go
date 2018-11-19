// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0
package main

import (
	"log"
	"net"

	"github.com/jweigend/concepts-of-programming-languages/dp/idserv/remote/idserv"
	"github.com/jweigend/concepts-of-programming-languages/dp/idserv/remote/stub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	idserv.RegisterIDServiceServer(s, &stub.Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
