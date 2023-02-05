package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"ports_microservice/utils"

	"github.com/gin-gonic/gin"
)

var (
	ports = CreatePortsMap()
)

func CreatePortsMap() map[string]PortData {
	var ports map[string]PortData
	byteValue := utils.ReadFile()

	err := json.Unmarshal(byteValue, &ports)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return ports
}

func GetMultiple(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ports)
}

func GetSingle(c *gin.Context) {
	portCode := c.Param("port_code")

	port, exists := ports[portCode]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Port not found"})
		return
	}

	c.JSON(http.StatusOK, port)
}
