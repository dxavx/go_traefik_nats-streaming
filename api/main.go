package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	"log"
	"net/http"
	"os"
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
	var natsUrl = os.Getenv("NATS_URL")
	var err error
	snc, err = stan.Connect(ClusterName, ClientID, stan.NatsURL(natsUrl))
	//snc, err = stan.Connect(ClusterName, ClientID)

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
	//router.Run(":8080")
	return router
}

func pubRandomNats(c *gin.Context) {

	var Subject = "my-subject-stan"
	var message = modules.RandomString(10)

	if err := snc.Publish(Subject, []byte(message)); err != nil {
		log.Fatalf("failed to publish to stan: %s", err.Error())
	}

	//time.Sleep(PublishDelay)
	//snc.Close()
	c.String(http.StatusOK, message)
}
