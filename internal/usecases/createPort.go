package usecases

import (
	"context"
	"fmt"
	"ports_microservice/internal/domain/entities"
	"ports_microservice/internal/domain/repositories"
)

type CreatePort struct {
	portStore repositories.PortStore
}

func NewCreatePort() *CreatePort {
	return &CreatePort{}
}

func (s *CreatePort) Execute(ctx context.Context, port *entities.Port) (*entities.Port, error) {
	fmt.Print("hellow world")
	// if err := s.portStore.Insert(port); err != nil {
	// 	log.Printf("inserting port failed: %s", err.Error())
	// 	return nil, err
	// }

	return port, nil
}
