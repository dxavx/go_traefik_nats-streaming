package main

import (
	"log"
	"net/http"
	"os"
	"traefik_test/api/modules"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	//"crypto/rand"
)

const (
	ClusterName = "test-cluster"
	ClientID    = "test-123"
)

var snc stan.Conn

func init() {
	var natsUrl = os.Getenv("NATS_URL")
	var err error
	snc, err = stan.Connect(ClusterName, ClientID, stan.NatsURL(natsUrl))

	if err != nil {
		log.Fatalf("failed to create nates connection: %s", err.Error())
	}
}

func main() {
	setupServer().Run()
}

func setupServer() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	v1 := router.Group("/v1")
	{
		v1.GET("/pub", pubRandomNats)
	}
	return router
}

func pubRandomNats(c *gin.Context) {

	var Subject = "my-subject-stan"
	var message = modules.RandomString(10)

	if err := snc.Publish(Subject, []byte(message)); err != nil {
		log.Fatalf("failed to publish to stan: %s", err.Error())
	}
	c.String(http.StatusOK, message)
}
