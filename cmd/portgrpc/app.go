package portgrpc

import (
	"ports_microservice/internal/drivers/grpcdriver"
)

type App struct {
	server *grpcdriver.Server
}

func NewApp(s *grpcdriver.Server) *App {
	return &App{
		server: s,
	}
}
