package main

import (
	"example/micro-service-project/client"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ports", client.GetMultiple)

	router.Run("localhost:5080")
}
