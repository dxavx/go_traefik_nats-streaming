package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	"log"
	"net/http"
	"time"
	"traefik_test/api/modules"
	//"crypto/rand"
)

const (
	PublishDelay = 100 * time.Millisecond
	ClusterName  = "test-cluster"
	ClientID     = "test-123"
)

var snc stan.Conn

func init() {
	var err error
	//snc, err = stan.Connect(ClusterName, ClientID, stan.NatsURL(nats.DefaultURL))
	snc, err = stan.Connect(ClusterName, ClientID)

	if err != nil {
		log.Fatalf("failed to create nates connection: %s", err.Error())
	}
}

func main() {

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/ping", ping)
		v1.GET("/connect", connectNats)
		v1.GET("/pub", pubRandomNats)
	}
	router.Run(":8080")
}

func ping(c *gin.Context) {
	if modules.PingNats() != nil {
		c.String(http.StatusInternalServerError, "NATS Ping Error")
	} else {
		c.String(http.StatusOK, "NATS Ping OK")
	}
}

func connectNats(c *gin.Context) {
	_, err := modules.ConnectNats()
	if err != nil {
		c.String(http.StatusInternalServerError, "NATS Connect Error")
	} else {
		c.String(http.StatusOK, "NATS Connect OK")
	}
}

func pubRandomNats(c *gin.Context) {

	var Subject = "my-subject-stan"
	var message = modules.RandomString(10)

	if err := snc.Publish(Subject, []byte(message)); err != nil {
		log.Fatalf("failed to publish to stan: %s", err.Error())
	}

	//time.Sleep(PublishDelay)
	//snc.Close()
	c.String(http.StatusOK, "Pub message : "+message)
}
