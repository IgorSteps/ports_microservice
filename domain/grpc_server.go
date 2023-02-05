package domain

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	pb "ports_microservice/domain/proto"
	constants "ports_microservice/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", constants.GrpcServerPort, "The server port")
)

type Server struct {
	pb.UnimplementedPortDomainServer
}

func (s *Server) GetPortList(ctx context.Context, in *pb.GetPortListRequest) (*pb.GetPortListResponse, error) {
	return &pb.GetPortListResponse{
		Ports: getSamplePorts(),
	}, nil
}

func Start() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterPortDomainServer(s, &Server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getSamplePorts() []*pb.Port {
	samplePorts := []*pb.Port{
		{
			Id:      "AEAJM",
			Name:    "Ajman",
			City:    "Ajman",
			Country: "United Arab Emirates",
			Alias:   []string{},
			Regions: []string{},
			Coordinates: []float32{
				55.5136433,
				25.4052165,
			},
			Province: "Ajman",
			Timezoe:  "Asia/Dubai",
			Unlocs: []string{
				"AEAJM",
			},
			Code: "52000",
		},
	}
	return samplePorts
}
