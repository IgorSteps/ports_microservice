package main

import (
	"bytes"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	//pb.RegisterPortDomainServer(s, &server.Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} else {
		logger.Print("running on port 8080")
	}
}
