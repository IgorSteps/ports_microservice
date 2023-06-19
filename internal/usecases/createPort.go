package usecases

import (
	"context"
	"log"
	"ports_microservice/internal/domain/entities"
	"ports_microservice/internal/domain/repositories"
)

type CreatePort struct {
	portStore repositories.PortStore
}

func NewCreatePort(ps repositories.PortStore) *CreatePort {
	return &CreatePort{
		portStore: ps,
	}
}

func (s *CreatePort) Execute(ctx context.Context, port *entities.Port) (*entities.Port, error) {
	if err := s.portStore.Insert(port); err != nil {
		log.Printf("inserting port failed: %s", err.Error())
		return nil, err
	}

	return port, nil
}
