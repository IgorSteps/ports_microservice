package db

import (
	"fmt"
	"log"
	models "ports_microservice/internal/adapters/datastore/model"
	"ports_microservice/internal/drivers"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dbURL := fmt.Sprintf(
		"%s://%s:%s@%s/%s",
		drivers.DBtype,
		drivers.User,
		drivers.Password,
		drivers.Port,
		drivers.DBname,
	)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open db connection: %v", err)
	}

	db.AutoMigrate(&models.PortSchema{})

	log.Printf("DB connection opened: %s", dbURL)

	return db
}
