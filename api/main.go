package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()
	//gin.SetMode(gin.DebugMode)

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", Ping)
	}
	router.Run(":8080")
}

// Ping is function check health service
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "ОК")
}
