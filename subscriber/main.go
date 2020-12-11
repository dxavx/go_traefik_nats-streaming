package main

import (
	"crypto/rand"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
)

const (
	ClusterName = "test-cluster"
	ClientID    = "test-1234"
)

var snc stan.Conn

func main() {

	var err error
	//snc, err = stan.Connect(ClusterName, ClientID, stan.NatsURL(nats.DefaultURL))
	snc, err = stan.Connect(ClusterName, GenUUIDv4())

	if err != nil {
		log.Fatalf("failed to create nates connection: %s", err.Error())
	}

	var Subject = "my-subject-stan"
	var Queue = "my-queue-stan"
	//var Chan =    make(chan string)

	sub, err := snc.QueueSubscribe(Subject, Queue, func(msg *stan.Msg) {
		//Chan <- fmt.Sprintf("receive message in stan subscriber %s: %s", "name", string(msg.Data))
		fmt.Println(string(msg.Data))

		//if err := msg.Ack(); err != nil {
		//	log.Fatalf("failed to send ack to stan in subscriber %s: %s", "name", err.Error())
		//}
		//}, stan.DurableName("q"))
	}, stan.DurableName("q"))
	if err != nil {
		log.Fatalf("failed to create stan subsciber %s: %s", "name", err.Error())
	}

	select {}

	defer func() {
		sub.Unsubscribe()
		snc.Close()
	}()

}

func GenUUIDv4() string {
	u := make([]byte, 16)
	rand.Read(u)
	//Set the version to 4
	u[6] = (u[6] | 0x40) & 0x4F
	u[8] = (u[8] | 0x80) & 0xBF
	ss := fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return ss
}
