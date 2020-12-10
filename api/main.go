package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"log"
	"net/http"
	"os"
	"time"
	"traefik_test/api/modules"
)

var nc *nats.Conn

func init() {

	var natsUrl = os.Getenv("NATS_URL")
	var err error

	nc, err = nats.Connect(natsUrl, nats.Timeout(time.Second*60))
	log.Println("Connected to " + natsUrl)
	if err != nil {
		log.Fatal(err)
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
	if err != nil {
		c.String(http.StatusInternalServerError, "NATS Connect Error")
	} else {
		c.String(http.StatusOK, "NATS Connect OK")
	}
}

func pubRandomNats(c *gin.Context) {

	var message = modules.RandomString(10)

	err := nc.Publish("updates", []byte(message))
	if err != nil {
		log.Fatal(err)
	} else {
		c.String(http.StatusOK, "Pub message : "+message)
	}
}
