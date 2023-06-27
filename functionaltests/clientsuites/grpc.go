package clientsuites

import (
	"context"
	"log"
	ports_v1 "ports_microservice/external/proto/ports"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClientSuite struct {
	suite.Suite

	Client ports_v1.PortsServiceClient
	Conn   *grpc.ClientConn
}

// Initilises all required objects in the stuct needed to create a port.
func (s *GrpcClientSuite) SetupSuite() {
	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial ports grpc endpoint: %v", err)
	}

	s.Conn = conn
	s.Client = ports_v1.NewPortsServiceClient(s.Conn)
}

// Called after every suite run to close connection.
func (s *GrpcClientSuite) TearDownSuite() {
	defer s.Conn.Close()
	s.T().Log("Closing connection...")
}

// Use this to call our gRPC client to create a port.
func (s *GrpcClientSuite) CreatePort(req *ports_v1.CreatePortRequest) (*ports_v1.CreatePortResponse, error) {
	ctx := context.Background()

	resp, err := s.Client.CreatePort(ctx, req)
	if err != nil {
		s.T().Logf("failed to create port: %v", err)
		return nil, err
	}

	return resp, nil
}
