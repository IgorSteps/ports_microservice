package main

import (
	"ports_microservice/client"
	"ports_microservice/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode) // to switch off debug mode
	router := gin.Default()
	router.GET("/ports", client.GetMultiple)
	router.GET("/port/:port_code", client.GetSingle)

	domain.StartClient()
	router.Run("localhost:5080")

}
