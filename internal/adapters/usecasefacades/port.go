package usecasefacades

import (
	"context"
	"ports_microservice/internal/domain/entities"
)

type PortCreator interface {
	Execute(ctx context.Context, port *entities.Port) (*entities.Port, error)
}

type PortGetter interface {
	Execute() ([]*entities.Port, error)
}

type PortFacade struct {
	portCreator PortCreator
	portGetter  PortGetter
}

func NewPortFacade(pc PortCreator, pg PortGetter) *PortFacade {
	return &PortFacade{
		portCreator: pc,
		portGetter:  pg,
	}
}

func (s *PortFacade) CreatePort(ctx context.Context, port *entities.Port) (*entities.Port, error) {
	return s.portCreator.Execute(ctx, port)
}

func (s *PortFacade) GetPorts() ([]*entities.Port, error) {
	return s.portGetter.Execute()
}
