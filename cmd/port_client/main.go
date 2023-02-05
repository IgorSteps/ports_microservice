package main

import (
	"ports_microservice/client"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ports", client.GetMultiple)
	router.GET("/port/:port_code", client.GetSingle)

	router.Run("localhost:5080")
}
