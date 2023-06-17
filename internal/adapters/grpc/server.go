package grpc

import (
	"context"
	"log"
	port_pb "ports_microservice/external/proto"
)

type PortGrpc struct {
	facade PortFacade
}

func NewPortGrpc(f PortFacade) *PortGrpc {
	return &PortGrpc{
		facade: f,
	}
}

func (s *PortGrpc) CreatePort(ctx context.Context, req *port_pb.CreatePortRequest) *port_pb.CreatePortResponse {
	port := ConvertProtoToEntity(req)

	createdPort, err := s.facade.CreatePort(ctx, port)
	if err != nil {
		log.Fatalf("Error occurred creating port: %v", err)
	}

	resp := ConvertEntityToProto(createdPort)

	return resp
}
