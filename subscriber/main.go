package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

const (
	ClusterName = "test-cluster"
)

var snc stan.Conn

func init() {
	var natsUrl = os.Getenv("NATS_URL")
	fmt.Println(nats.DefaultURL)
	fmt.Println(natsUrl)
	var err error
	snc, err = stan.Connect(ClusterName, GenUUIDv4(), stan.NatsURL(natsUrl))
	//snc, err = stan.Connect(ClusterName, GenUUIDv4())

	if err != nil {
		log.Fatalf("failed to create nates connection: %s", err.Error())
	}
}

func main() {

	var Subject = "my-subject-stan"
	var Queue = "my-queue-stan"
	//var Chan =    make(chan string)

	_, err := snc.QueueSubscribe(Subject, Queue, func(msg *stan.Msg) {
		//Chan <- fmt.Sprintf("receive message in stan subscriber %s: %s", "name", string(msg.Data))
		fmt.Println(string(msg.Data))

		//if err := msg.Ack(); err != nil {
		//	log.Fatalf("failed to send ack to stan in subscriber %s: %s", "name", err.Error())
		//}
	}, stan.DurableName("q"))
	if err != nil {
		log.Fatalf("failed to create stan subsciber %s: %s", "name", err.Error())
	}

	select {}

	//defer func() {
	//	sub.Unsubscribe()
	//	snc.Close()
	//}()

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
