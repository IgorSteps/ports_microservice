package client

import (
	"encoding/json"
	"example/micro-service-project/utils"
	"fmt"
	"net/http"

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
	}
	return ports
}

func GetMultiple(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ports)
}
