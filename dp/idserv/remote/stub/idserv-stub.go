// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

package stub

import (
	"context"

	"github.com/jweigend/concepts-of-programming-languages/dp/idserv/impl"
	"github.com/jweigend/concepts-of-programming-languages/dp/idserv/remote/idserv"
)

// Server is used to implement idserv.IdServer
type Server struct{}

// NewUUID implements idserv.IdService interface
func (s *Server) NewUUID(c context.Context, r *idserv.IdRequest) (*idserv.IdReply, error) {
	service := impl.IDServiceImpl{}
	return &idserv.IdReply{Uuid: service.NewUUID(r.GetClientId())}, nil
}
