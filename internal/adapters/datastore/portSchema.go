package datastore

import "ports_microservice/internal/domain/entities"

type PortSchema struct {
	ID   int
	Port entities.Port
}
