package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"time"
)

func main() {

	var natsUrl = os.Getenv("NATS_URL")
	//natsUrl = "http://localhost:4222"

	nc, err := nats.Connect(natsUrl, nats.Timeout(time.Second*60))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	//Asynchronous Subscriptions
	//Use a WaitGroup to wait for a message to arrive

	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//
	//// Subscribe
	//if _, err := nc.Subscribe("updates", func(m *nats.Msg) {
	//	fmt.Println(string(m.Data))
	//	wg.Done()
	//}); err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Wait for a message to come in
	//wg.Wait()

	// Synchronous Subscriptions

	for {
		// Subscribe
		sub, err := nc.SubscribeSync("updates")
		if err != nil {
			log.Fatal(err)
		}

		// Wait for a message
		msg, err := sub.NextMsg(120 * time.Second)
		if err != nil {
			log.Fatal(err)
		}

		// Use the response
		log.Printf("Reply: %s", msg.Data)
	}
}
