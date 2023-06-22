package wiredriver

import (
	"ports_microservice/internal/adapters/grpc"
	"ports_microservice/internal/drivers/grpcdriver"
)

// Implements gprcdriver.Service interface.
func NewGRPCService(portService *grpc.PortGrpcApi) grpcdriver.Service {
	return portService
}
