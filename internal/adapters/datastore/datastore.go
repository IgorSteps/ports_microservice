package datastore

import (
	"ports_microservice/internal/domain/entities"
	"ports_microservice/internal/domain/repositories"
)

type PortDatastore struct {
	dbWrapper DBWrapper
	store     repositories.PortStore
}

func NewPortDataStore(dbW DBWrapper, s repositories.PortStore) *PortDatastore {
	return &PortDatastore{
		dbWrapper: dbW,
		store:     s,
	}
}

func (s *PortDatastore) Insert(port *entities.Port) error {
	if err := s.dbWrapper.Create(port).Error(); err != nil {
		return err
	}

	return nil
}
