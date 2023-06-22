package usecasefacades

import (
	"context"
	"ports_microservice/internal/domain/entities"
)

type PortCreator interface {
	Execute(ctx context.Context, port *entities.Port) (*entities.Port, error)
}

type PortFacade struct {
	portCreator PortCreator
}

func NewPortFacade(pc PortCreator) *PortFacade {
	return &PortFacade{
		portCreator: pc,
	}
}

func (s *PortFacade) CreatePort(ctx context.Context, port *entities.Port) (*entities.Port, error) {
	return s.portCreator.Execute(ctx, port)
}
