package grpcdriver

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Service interface {
	Register(s *grpc.Server)
}

type Server struct {
	server *grpc.Server
	port   int
}

func NewServer(service Service) *Server {
	s := grpc.NewServer()
	service.Register(s)

	return &Server{
		server: s,
		port:   5001,
	}
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", s.port))
	if err != nil {
		return err
	}

	return s.server.Serve(lis)
}
