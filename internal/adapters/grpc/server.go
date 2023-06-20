package grpc

import (
	"context"
	"log"
	port_pb "ports_microservice/external/proto"

	"google.golang.org/grpc"
)

type PortGrpcApi struct {
	port_pb.UnimplementedPortServiceServer

	facade PortFacade
}

func NewPortGrpc(f PortFacade) *PortGrpcApi {
	return &PortGrpcApi{
		facade: f,
	}
}

func (s *PortGrpcApi) CreatePort(ctx context.Context, req *port_pb.CreatePortRequest) (*port_pb.CreatePortResponse, error) {
	port := ConvertProtoToEntity(req)

	createdPort, err := s.facade.CreatePort(ctx, port)
	if err != nil {
		log.Fatalf("Error occurred creating port: %v", err)
	}

	resp := ConvertEntityToProto(createdPort)

	return resp, nil
}

func (api *PortGrpcApi) Register(server *grpc.Server) {
	port_pb.RegisterPortServiceServer(server, api)
}
