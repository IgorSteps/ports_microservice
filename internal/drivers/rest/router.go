package rest

import (
	"fmt"
	"log"
	"ports_microservice/internal/adapters/grpc"
	"ports_microservice/internal/adapters/rest"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
	port   int
	api    *rest.PortsRestApi
}

func NewRouter(facade grpc.PortFacade) *Router {
	r := gin.Default()
	api := rest.NewPortRestApi(facade)

	r.GET("/ports", api.GetPorts)

	return &Router{
		router: r,
		port:   3000,
		api:    api,
	}
}

func (s *Router) Run() error {
	endpoint := fmt.Sprintf("localhost:%d", s.port)
	log.Printf("REST server listening at %s", endpoint)

	return s.router.Run(endpoint)
}
