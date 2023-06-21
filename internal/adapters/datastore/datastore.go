package datastore

import (
	models "ports_microservice/internal/adapters/datastore/model"
	"ports_microservice/internal/domain/entities"

	"github.com/google/uuid"
)

type PortDatastore struct {
	dbWrapper DBWrapper
}

func NewPortDataStore(dbW DBWrapper) *PortDatastore {
	return &PortDatastore{
		dbWrapper: dbW,
	}
}

func (s *PortDatastore) Insert(port *entities.Port) error {
	// TOOD: Move to a converters file like in grpc adapter.
	portShema := &models.PortSchema{
		Id:          uuid.New(),
		Name:        port.Name,
		City:        port.City,
		Country:     port.Country,
		Alias:       port.Alias,
		Regions:     port.Regions,
		Coordinates: port.Coordinates,
		Province:    port.Province,
		Timezone:    port.Timezone,
		Unlocs:      port.Unlocs,
		Code:        port.Code,
	}
	if err := s.dbWrapper.Create(portShema).Error(); err != nil {
		return err
	}

	return nil
}
