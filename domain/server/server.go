package server

import (
	"context"
	"example/micro-service-project/domain/pb"
)

type Server struct {
	pb.UnimplementedPortDomainServer
}

func (s *Server) GetPortList(ctx context.Context, in *pb.GetPortListRequest) (*pb.GetPortListResponse, error) {
	return &pb.GetPortListResponse{
		Ports: getSamplePorts(),
	}, nil
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
