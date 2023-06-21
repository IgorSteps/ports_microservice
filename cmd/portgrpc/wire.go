//go:build wireinject
// +build wireinject

package main

import (
	"ports_microservice/internal/adapters/grpc"
	portGrpc "ports_microservice/internal/adapters/grpc"
	"ports_microservice/internal/adapters/usecasefacades"
	"ports_microservice/internal/drivers/grpcdriver"
	"ports_microservice/internal/drivers/wiredriver"
	"ports_microservice/internal/usecases"

	"github.com/google/wire"
)

func BuildDIForApp() (*App, error) {
	wire.Build(
		// usecase:
		usecases.NewCreatePort,
		wire.Bind(new(usecasefacades.PortCreator), new(*usecases.CreatePort)),
		// facade for usecases:
		usecasefacades.NewPortFacade,
		grpc.NewPortGrpc,
		wire.Bind(new(portGrpc.PortFacade), new(*usecasefacades.PortFacade)),
		// grpc service:
		wiredriver.NewGRPCService,
		// grpc driver:
		grpcdriver.NewServer,
		// service:
		NewApp,
	)

	return &App{}, nil
}