package models

import (
	"ports_microservice/internal/domain/entities"

	"github.com/google/uuid"
)

func ConvertEntityToDB(port *entities.Port) *PortSchema {
	return &PortSchema{
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
}

func ConvertDBToEntity(port *PortSchema) *entities.Port {
	return &entities.Port{
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
}
