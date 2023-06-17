package repositories

import (
	"ports_microservice/internal/domain/entities"

	"github.com/google/uuid"
)

type PortRepo interface {
	Create(port *entities.Port) error
	GetPort(portID uuid.UUID) (*entities.Port, error)
	GetPorts() ([]*entities.Port, error)
	UpdatePort(portID uuid.UUID) error
}
