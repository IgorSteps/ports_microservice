package usecases

import (
	"log"
	"ports_microservice/internal/domain/entities"
	"ports_microservice/internal/domain/repositories"
)

type GetPorts struct {
	portStore repositories.PortStore
}

func NewGetPorts(ps repositories.PortStore) *GetPorts {
	return &GetPorts{
		portStore: ps,
	}
}

func (s *GetPorts) GetPorts() ([]*entities.Port, error) {
	ports, err := s.portStore.GetMultiple()
	if err != nil {
		log.Printf("Failed to get ports: %v", err)
		return nil, err
	}

	return ports, nil
}
