//go:build wireinject
// +build wireinject

package main

import (
	"ports_microservice/internal/adapters/datastore"
	portGrpc "ports_microservice/internal/adapters/grpc"
	"ports_microservice/internal/adapters/usecasefacades"
	"ports_microservice/internal/domain/repositories"
	"ports_microservice/internal/drivers/db"
	"ports_microservice/internal/drivers/rest"
	"ports_microservice/internal/usecases"

	"github.com/google/wire"
)

func BuildDIForApp() (*App, error) {
	wire.Build(
		// datastore:
		datastore.NewPortDataStore,
		wire.Bind(new(repositories.PortStore), new(*datastore.PortDatastore)),
		// db:
		db.NewDB,
		// gorm db wrapper:
		db.NewGormDBWrapper,
		wire.Bind(new(datastore.DBWrapper), new(*db.GormDBWrapper)),
		// usecase:
		usecases.NewCreatePort,
		wire.Bind(new(usecasefacades.PortCreator), new(*usecases.CreatePort)),
		usecases.NewGetPorts,
		wire.Bind(new(usecasefacades.PortGetter), new(*usecases.GetPorts)),
		// facade for usecases:
		usecasefacades.NewPortFacade,
		wire.Bind(new(portGrpc.PortFacade), new(*usecasefacades.PortFacade)),
		// rest driver
		rest.NewRouter,
		// service:
		NewApp,
	)

	return &App{}, nil
}
