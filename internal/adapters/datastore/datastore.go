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

func (s *PortDatastore) GetMultiple() ([]*entities.Port, error) {
	// Get all ports.
	var ports []*models.PortSchema
	if err := s.dbWrapper.Find(&ports).Error(); err != nil {
		return nil, err
	}

	// Convert db to entity ports.
	var domainPorts []*entities.Port
	for _, port := range ports {
		domainPorts = append(domainPorts, models.ConvertDBToEntity(port))
	}

	return domainPorts, nil
}
