package grpcdriver

import (
	"fmt"
	"log"
	"net"
	"ports_microservice/internal/drivers"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	reflection.Register(s)
	service.Register(s)

	return &Server{
		server: s,
		port:   drivers.GRPCServerPort,
	}
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.port))
	if err != nil {
		return err
	}

	log.Printf("GRPC server listening at %v", lis.Addr())

	return s.server.Serve(lis)
}
