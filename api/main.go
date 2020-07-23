package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"traefik_test/api/modules"
)

func main() {

	router := gin.Default()
	//gin.SetMode(gin.DebugMode)

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", ping)
		v1.GET("/connect", connectNats)
	}
	router.Run(":8080")
}

// Ping is function check health service
func ping(c *gin.Context) {
	c.String(http.StatusOK, "ОК")
}

func connectNats(c *gin.Context) {
	modules.ConnectNats()
}
