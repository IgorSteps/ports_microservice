package repositories

import (
	"ports_microservice/internal/domain/entities"
)

type PortStore interface {
	Insert(port *entities.Port) error
	GetMultiple() ([]*entities.Port, error)
}
