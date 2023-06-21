package db

import (
	"log"
	models "ports_microservice/internal/adapters/datastore/model"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://pguser:password@localhost:5432/crud"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.PortSchema{})

	return db
}
