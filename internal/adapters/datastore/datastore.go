package datastore

import (
	models "ports_microservice/internal/adapters/datastore/model"
	"ports_microservice/internal/domain/entities"
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
	portShema := models.ConvertEntityToDB(port)
	if err := s.dbWrapper.Create(portShema).Error(); err != nil {
		return err
	}

	return nil
}
