package repositories

import (
	"ports_microservice/internal/domain/entities"
)

type PortStore interface {
	Insert(port *entities.Port) error
	GetSingle(id string) error
}
