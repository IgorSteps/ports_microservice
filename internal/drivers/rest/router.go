package rest

import (
	"fmt"
	"log"
	"ports_microservice/internal/adapters"
	"ports_microservice/internal/adapters/portsrest"
	"ports_microservice/internal/drivers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
	port   int
	api    *portsrest.PortsRestApi
}

func NewRouter(facade adapters.PortFacade) *Router {
	r := gin.Default()
	api := portsrest.NewPortRestApi(facade)

	r.GET("/ports", api.GetPorts)

	return &Router{
		router: r,
		port:   drivers.RESTRouterPort,
		api:    api,
	}
}

func (s *Router) Run() error {
	endpoint := fmt.Sprintf("localhost:%d", s.port)
	log.Printf("REST server listening at %s", endpoint)

	return s.router.Run(endpoint)
}
