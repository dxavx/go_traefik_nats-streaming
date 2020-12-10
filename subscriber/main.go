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

		// -----------------------------------------------------------------

		//_, err = Connection.Subscribe("test", func(msg *nats.Msg) {
		//	fmt.Println(string(msg.Data))
		//})
		//if err != nil {
		//	log.Fatalf("failed to create nats subsciber %s: %s", "name", err.Error())
		//}
	}
}
