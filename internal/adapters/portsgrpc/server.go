package portsgrpc

import (
	"context"
	"log"
	port_pb "ports_microservice/external/proto/ports"
	"ports_microservice/internal/adapters"

	"google.golang.org/grpc"
)

type PortGrpcApi struct {
	port_pb.UnimplementedPortsServiceServer

	facade adapters.PortFacade
}

func NewPortGrpcApi(f adapters.PortFacade) *PortGrpcApi {
	return &PortGrpcApi{
		facade: f,
	}
}

func (s *PortGrpcApi) CreatePort(ctx context.Context, req *port_pb.CreatePortRequest) (*port_pb.CreatePortResponse, error) {
	port := ConvertProtoToEntity(req)

	createdPort, err := s.facade.CreatePort(ctx, port)
	if err != nil {
		log.Printf("Error occurred creating port: %v", err)
		return nil, err
	}

	resp := ConvertEntityToProto(createdPort)

	return resp, nil
}

// Register(server)
//
// Registers a gRPC server to Ports service server.
// Needs to implement gprcdriver.Service
func (api *PortGrpcApi) Register(server *grpc.Server) {
	port_pb.RegisterPortsServiceServer(server, api)
}
