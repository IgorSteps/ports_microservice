package wiredriver

import (
	"ports_microservice/internal/adapters/portsgrpc"
	"ports_microservice/internal/drivers/grpcdriver"
)

// Implements gprcdriver.Service interface.
func NewGRPCService(portService *portsgrpc.PortGrpcApi) grpcdriver.Service {
	return portService
}
