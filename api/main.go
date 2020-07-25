package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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
		v1.GET("/pub", pubRandomNats)
	}
	router.Run(":8080")
}

// Ping is function check health service
func ping(c *gin.Context) {
	if modules.PingNats() != nil {
		c.String(http.StatusInternalServerError, "NATS Ping Error")
	} else {
		c.String(http.StatusOK, "NATS Ping OK")
	}
}

func connectNats(c *gin.Context) {
	_, err := modules.ConnectNats()
	fmt.Println(modules.RandomeString(10))
	if err != nil {
		c.String(http.StatusInternalServerError, "NATS Connect Error")
	} else {
		c.String(http.StatusOK, "NATS Connect OK")
	}
}

func pubRandomNats(c *gin.Context) {

	var message = modules.RandomeString(10)
	nc, err := modules.ConnectNats()
	if err != nil {
		log.Fatal(err)
	} else {
		nc.Publish("update", []byte(message))
		c.String(http.StatusOK, "Pub message : "+message)
	}

	//defer nc.Close()
	//if err != nil {
	//	c.String(http.StatusInternalServerError, "NATS Connect Error")
	//} else {
	//	if err := nc.Publish("updates", []byte(message)); err != nil {
	//		log.Fatal(err)
	//	}
	//	c.String(http.StatusOK, "Pub message : " + message)
	//}

}
