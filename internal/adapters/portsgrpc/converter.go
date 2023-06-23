package portsgrpc

import (
	port_pb "ports_microservice/external/proto/ports"
	"ports_microservice/internal/domain/entities"
)

func ConvertProtoToEntity(req *port_pb.CreatePortRequest) *entities.Port {
	return &entities.Port{
		Name:        req.Name,
		City:        req.City,
		Country:     req.Country,
		Alias:       req.Alias,
		Regions:     req.Regions,
		Coordinates: req.Coordinates,
		Province:    req.Province,
		Timezone:    req.Timezone,
		Unlocs:      req.Unlocs,
		Code:        req.Code,
	}
}

func ConvertEntityToProto(port *entities.Port) *port_pb.CreatePortResponse {
	return &port_pb.CreatePortResponse{
		Name:        port.Name,
		City:        port.City,
		Country:     port.Country,
		Alias:       port.Alias,
		Regions:     port.Regions,
		Coordinates: port.Coordinates,
		Province:    port.Province,
		Timezone:    port.Timezone,
		Unlocs:      port.Unlocs,
		Code:        port.Code,
	}
}
