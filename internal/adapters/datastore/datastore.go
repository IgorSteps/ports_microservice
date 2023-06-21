package datastore

import (
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
	if err := s.dbWrapper.Create(port).Error(); err != nil {
		return err
	}

	return nil
}
