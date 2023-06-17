package repositories

import (
	"ports_microservice/internal/domain/entities"

	"github.com/google/uuid"
)

type PortStore interface {
	Insert(port *entities.Port) error
	GetSingle(portID uuid.UUID) error
	Update(portID uuid.UUID) error
}
