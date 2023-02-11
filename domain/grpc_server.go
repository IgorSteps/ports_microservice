package domain

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"ports_microservice/client"
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
	db map[string]client.PortData
}

// AddOrUpdate adds or updates a port
func (s *Server) AddOrUpdate(ctx context.Context, request *pb.Port) (*pb.Response, error) {
	s.db[request.Id] = client.PortData{
		Name:        request.Name,
		City:        request.City,
		Country:     request.Country,
		Alias:       request.Alias,
		Regions:     request.Regions,
		Coordinates: request.Coordinates,
		Province:    request.Province,
		Timezone:    request.Timezone,
		Unlocs:      request.Unlocs,
		Code:        request.Code,
	}
	return &pb.Response{
		Message: "Port added or updated successfully",
	}, nil
}

// Get retrieves a port from the database by its ID
func (s *Server) Get(ctx context.Context, id *pb.ID) (*pb.Port, error) {
	p, ok := s.db[id.Id]
	if !ok {
		return nil, fmt.Errorf("port not found")
	}
	return &pb.Port{
		Name:        p.Name,
		City:        p.City,
		Country:     p.Country,
		Alias:       p.Alias,
		Regions:     p.Regions,
		Coordinates: p.Coordinates,
		Province:    p.Province,
		Timezone:    p.Timezone,
		Unlocs:      p.Unlocs,
		Code:        p.Code,
	}, nil
}

// GetPortList retrieves all Ports from the DB
func (s *Server) GetPortList(ctx context.Context, in *pb.GetPortListRequest) (*pb.GetPortListResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())

	return &pb.GetPortListResponse{
		Ports: getSamplePorts(), // TODO: return from DB
	}, nil
}

// Start gRPC server
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

// Mock function to return Sample Port data for testing
func getSamplePorts() []*pb.Port {
	samplePorts := []*pb.Port{
		{
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
			Timezone: "Asia/Dubai",
			Unlocs: []string{
				"AEAJM",
			},
			Code: "52000",
		},
	}
	return samplePorts
}
