package portsrest

import (
	"log"
	"net/http"
	"ports_microservice/internal/adapters"

	"github.com/gin-gonic/gin"
)

type PortsRestApi struct {
	facade adapters.PortFacade
}

func NewPortRestApi(f adapters.PortFacade) *PortsRestApi {
	return &PortsRestApi{
		facade: f,
	}
}

func (s *PortsRestApi) GetPorts(c *gin.Context) {
	ports, err := s.facade.GetPorts()
	if err != nil {
		log.Fatalf("Error occured getting ports %v", err)
	}

	c.IndentedJSON(http.StatusOK, ports)
}
