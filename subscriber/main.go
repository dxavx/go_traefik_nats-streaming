package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"os"
)

func main() {

	var Subject = "updates"
	var Queue = "workers"
	var Chan = make(chan string)
	//var Connection *nats.Conn

	var natsUrl = os.Getenv("NATS_URL")
	var err error
	fmt.Println("NATS Connect start : ", natsUrl)
	Connection, err := nats.Connect(natsUrl)

	if err != nil {
		log.Fatalf("failed to create nates connection: %s", err.Error())
	}

	for {
		_, err := Connection.QueueSubscribe(Subject, Queue, func(msg *nats.Msg) {
			Chan <- fmt.Sprintf("receive message in nats subscriber %s: %s", "name", string(msg.Data))
		})
		if err != nil {
			log.Fatalf("failed to create nats subsciber %s: %s", "name", err.Error())
		}
		//msg := <-nc.Chan
		fmt.Println(<-Chan)
	}

	//Asynchronous Subscriptions
	//Use a WaitGroup to wait for a message to arrive

	//Synchronous Subscriptions

	//for {
	//	// Subscribe
	//	sub, err := nc.SubscribeSync("updates")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	// Wait for a message
	//	msg, err := sub.NextMsg(120 * time.Second)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	// Use the response
	//	log.Printf("Reply: %s", msg.Data)
	//}
}

//func (n NatsConnection) SubscribeQuene(name string) {
//	_, err := n.Connection.QueueSubscribe(n.Subject, n.Queue, func(msg *nats.Msg) {
//		n.Chan <- fmt.Sprintf("receive message in nats subscriber %s: %s", name, string(msg.Data))
//	})
//	if err != nil {
//		log.Fatalf("failed to create nats subsciber %s: %s", name, err.Error())
//	}
//}
//
//func (n NatsConnection) Subscribe(name string) {
//	_, err := n.Connection.Subscribe(n.Subject, func(msg *nats.Msg) {
//		n.Chan <- fmt.Sprintf("receive message in nats subscriber %s: %s", name, string(msg.Data))
//	})
//	if err != nil {
//		log.Fatalf("failed to create nats subsciber %s: %s", name, err.Error())
//	}
//}
