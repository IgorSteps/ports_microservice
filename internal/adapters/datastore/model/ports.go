package models

import (
	"ports_microservice/internal/domain/entities"

	"gorm.io/gorm"
)

type PortSchema struct {
	gorm.Model
	entities.Port
}
