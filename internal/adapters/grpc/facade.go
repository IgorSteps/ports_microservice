package grpc

import (
	"context"
	"ports_microservice/internal/domain/entities"
)

type PortFacade interface {
	CreatePort(ctx context.Context, port *entities.Port) (*entities.Port, error)
	GetPorts() ([]*entities.Port, error)
}
